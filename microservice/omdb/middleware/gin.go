package middleware

import (
	"fmt"
	"omdb/logger"

	"github.com/gin-gonic/gin"
)

func WithDBLogger(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		go log.Log("[status]:", fmt.Sprintf("[%d]", c.Writer.Status()), "[request]:", c.Request, "[params]", c.Params)
	}
}
