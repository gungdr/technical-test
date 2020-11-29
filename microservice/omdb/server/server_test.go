package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"omdb/config"
	"omdb/handler"
	"omdb/movie"
	"omdb/proto"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	pbwrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc"
)

var omdbClient proto.OMDBServiceClient

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var router *gin.Engine

func init() {
	rpcServer := grpc.NewServer()
	conf := config.New()
	lis = bufconn.Listen(bufSize)
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieRPCHandler := handler.NewRPCHandler(conf, movieService)
	movieServer := NewRPCServer(rpcServer, conf, lis, movieRPCHandler)
	router = gin.New()
	movieRestHandler := handler.NewRestHandler(conf, movieService)
	movieRestServer := NewRestServer(router, conf, movieRestHandler)
	go movieRestServer.Run()
	go movieServer.Run()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_server_Search(t *testing.T) {
	requestID, _ := uuid.NewUUID()
	ctx := context.WithValue(context.Background(), "request-id", requestID)
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	omdbClient := proto.NewOMDBServiceClient(conn)

	param := &proto.Param{
		Keyword: "kill",
		Page:    1,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = omdbClient.Search(ctx, param)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func Test_server_Get(t *testing.T) {
	requestID, _ := uuid.NewUUID()
	ctx := context.WithValue(context.Background(), "request-id", requestID)
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	omdbClient := proto.NewOMDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	_, err = omdbClient.Get(ctx, &pbwrapper.StringValue{Value: "tt0048119"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func Test_restHandler_Get(t *testing.T) {
	ts := httptest.NewServer(router)
	defer ts.Close()
	resp, err := http.Get(fmt.Sprintf("%s/v1/movie/%s", ts.URL, "tt0120912"))
	if err != nil {
		t.Errorf("Should not error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func Test_restHandler_Search(t *testing.T) {
	ts := httptest.NewServer(router)
	defer ts.Close()
	resp, err := http.Get(fmt.Sprintf("%s/v1/search?k=Batman&p=1", ts.URL))
	if err != nil {
		t.Errorf("Should not error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
