package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
	"wms/controller"
	"wms/pkg"
	"wms/repo"
	"wms/service"
)

func InternalRoutes(ctx context.Context, s *http.Server) (err error) {
	rtr := s.Engine.Group("/api/v1")

	// todo use go wire if needed
	newRepository := repo.NewRepository(pkg.GetCluster().DbCluster)
	newService := service.NewService(newRepository)
	controller := controller.NewController(newService)

	// make apis for it
	rtr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "mst"})
	})

	rtr.GET("/hub", controller.GetHubs())

	return
}
