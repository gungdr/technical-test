package server

import (
	"context"
	"log"
	"net"
	"omdb/config"
	"omdb/movie"
	"omdb/proto"
	"testing"
	"time"

	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc"
)

var omdbClient proto.OMDBServiceClient

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	rpcServer := grpc.NewServer()
	conf := config.New()
	lis = bufconn.Listen(bufSize)
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieServer := NewRPCServer(rpcServer, conf, lis, movieService)
	go movieServer.Run()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_server_Search(t *testing.T) {
	ctx := context.Background()
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
