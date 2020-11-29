package server

import (
	"fmt"
	"log"
	"omdb/config"
	"omdb/handler"

	"github.com/gin-gonic/gin"
)

type restServer struct {
	router  *gin.Engine
	config  *config.Config
	handler handler.Handler
}

func NewRestServer(
	router *gin.Engine,
	conf *config.Config,
	handler handler.Handler) Server {
	return &restServer{
		router:  router,
		config:  conf,
		handler: handler,
	}
}
func (s *restServer) Run() {
	log.Printf("REST Server is starting on Port :%d ...\n", s.config.RestPort)
	s.Register()
	s.router.Run(fmt.Sprintf(":%d", s.config.RestPort))
}

func (s *restServer) Register() {
	v1 := s.router.Group("/v1")
	{
		v1.GET("/movie/:id", s.handler.Get)
		v1.GET("/search", s.handler.Search)
	}
}
