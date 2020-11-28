package main

import (
	"fmt"
	"log"
	"net"
	"omdb/config"
	"omdb/movie"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	conf := config.New()
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieServer := movie.NewServer(rpcServer, conf, movieService)
	movieServer.Run()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
