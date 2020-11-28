package main

import (
	"omdb/config"
	"omdb/movie"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	router := gin.New()
	conf := config.New()
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieRPCServer := movie.NewRPCServer(rpcServer, conf, movieService)
	movieRestServer := movie.NewRestServer(router, conf, movieService)
	movieRPCServer.Run()
	movieRestServer.Run()

}
