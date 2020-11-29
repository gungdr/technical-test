package main

import (
	"fmt"
	"log"
	"net"
	"omdb/config"
	"omdb/logger"
	"omdb/middleware"
	"omdb/movie"
	"omdb/server"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conf := config.New()
	dbLogger := logger.NewDBLogger("omdb")
	rpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			middleware.InterceptorWithDBLogger(dbLogger),
		),
	)
	router := gin.New()
	router.Use(middleware.WithDBLogger(dbLogger))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.RPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	movieClient := movie.NewClient(conf)
	movieService := movie.NewService(movieClient)
	movieRPCServer := server.NewRPCServer(rpcServer, conf, listener, movieService)
	movieRestServer := server.NewRestServer(router, conf, movieService)
	go movieRPCServer.Run()
	movieRestServer.Run()

}
