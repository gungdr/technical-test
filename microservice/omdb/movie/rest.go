package movie

import (
	"log"
	"omdb/config"

	"github.com/gin-gonic/gin"
)

type restServer struct {
	router  *gin.Engine
	config  *config.Config
	service Service
}

func NewRestServer(router *gin.Engine, conf *config.Config, service Service) Server {
	return &restServer{
		router:  router,
		config:  conf,
		service: service,
	}
}
func (s *restServer) Run() {
	log.Printf("REST Server is starting on Port :%d ...\n", s.config.RestPort)
	s.router.Run()
}
