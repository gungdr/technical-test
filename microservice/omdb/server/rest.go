package server

import (
	"fmt"
	"log"
	"omdb/config"
	"omdb/movie"

	"github.com/gin-gonic/gin"
)

type restServer struct {
	router  *gin.Engine
	config  *config.Config
	service movie.Service
}

func NewRestServer(router *gin.Engine, conf *config.Config, service movie.Service) Server {
	return &restServer{
		router:  router,
		config:  conf,
		service: service,
	}
}
func (s *restServer) Run() {
	log.Printf("REST Server is starting on Port :%d ...\n", s.config.RestPort)
	s.router.Run(fmt.Sprintf(":%d", s.config.RestPort))
}
