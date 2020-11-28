package main

import (
	"omdb/config"
	"omdb/movie"
	"omdb/server"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	router := gin.New()
	conf := config.New()
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieRPCServer := server.NewRPCServer(rpcServer, conf, movieService)
	movieRestServer := server.NewRestServer(router, conf, movieService)
	go movieRPCServer.Run()
	movieRestServer.Run()

}
