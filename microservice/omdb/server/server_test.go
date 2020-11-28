package server

import (
	"context"
	"fmt"
	"log"
	"omdb/config"
	"omdb/movie"
	"omdb/proto"
	"testing"
	"time"

	"google.golang.org/grpc"
)

var omdbClient proto.OMDBServiceClient

// const bufSize = 1024 * 1024

// var lis *bufconn.Listener
var conf *config.Config

func init() {
	rpcServer := grpc.NewServer()
	conf = config.New()
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieServer := NewRPCServer(rpcServer, conf, movieService)
	movieServer.Run()

	// conn, err := grpc.Dial(fmt.Sprintf("http://localhost:%d", conf.RPCPort), grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// // lis = bufconn.Listen(bufSize)
	// go func() {
	// 	if err := rpcServer.Serve(lis); err != nil {
	// 		log.Fatalf("Server exited with error: %v", err)
	// 	}
	// }()
}

// func bufDialer(context.Context, string) (net.Conn, error) {
// 	return lis.Dial()
// }

func Test_server_Search(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.Dial(fmt.Sprintf("http://localhost:%d", conf.RPCPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
