package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
)

func InternalRoutes(ctx context.Context, s *http.Server) (err error) {
	rtr := s.Engine.Group("/api/v1")

	// make apis for it
	rtr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "mst"})
	})

	return
}
