package handler

import (
	"context"
	"errors"
	"net/http"
	"omdb/config"
	"omdb/movie"
	"omdb/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type restHandler struct {
	config  *config.Config
	service movie.Service
}

func NewRestHandler(
	config *config.Config,
	service movie.Service) Handler {
	return &restHandler{
		config:  config,
		service: service,
	}
}

func (h *restHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.Error(errors.New("Movie ID is required")))
		return
	}
	movie, err := h.service.Get(context.Background(), movie.IMDBID(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(movie))
}

func (h *restHandler) Search(c *gin.Context) {
	pageQuery := c.DefaultQuery("p", "1")
	keyword := c.Query("k")
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
	searchResult, err := h.service.Search(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(searchResult))
}
