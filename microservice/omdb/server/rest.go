package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"omdb/config"
	"omdb/movie"
	"omdb/response"
	"strconv"

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
	s.Register()
	s.router.Run(fmt.Sprintf(":%d", s.config.RestPort))
}

func (s *restServer) Register() {
	v1 := s.router.Group("/v1")
	{
		v1.GET("/movie/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusBadRequest, response.Error(errors.New("Movie ID is required")))
				return
			}
			movie, err := s.service.Get(context.Background(), movie.IMDBID(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.Error(err))
				return
			}
			c.JSON(http.StatusOK, response.Success(movie))
		})

		v1.GET("/search", func(c *gin.Context) {
			pageQuery := c.DefaultQuery("p", "1")
			keyword := c.Query("k")
			size := c.
			if keyword == "" || pageQuery == "" {
				c.JSON(http.StatusBadRequest, response.Error(errors.New("Movie ID is required")))
				return
			}
			page, err := strconv.Atoi(pageQuery)
			param := movie.Param{
				Keyword: keyword,
				Page:    page,
			}
			if param.Valid() == false {
				c.JSON(http.StatusBadRequest, response.Error(errors.New("Invalid Query String")))
				return
			}
			searchResult, err := s.service.Search(context.Background(), param)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.Error(err))
				return
			}
			c.JSON(http.StatusOK, response.Success(searchResult))
		})
	}
}
