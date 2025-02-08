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

	// todo go wire
	newRepository := repo.NewRepository(pkg.GetCluster().DbCluster)
	newService := service.NewService(newRepository)
	newController := controller.NewController(newService)

	rtr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "mst"})
	})

	// Hub routes
	rtr.GET("/hub", newController.GetHubs())
	rtr.GET("/hub/:id", newController.GetHubByID())
	rtr.POST("/hub", newController.CreateHub())

	// SKU routes
	rtr.GET("/sku", newController.GetSkus())
	rtr.GET("/sku/:id", newController.GetSkuByID())
	rtr.POST("/sku", newController.CreateSKU())

	// Inventory routes
	rtr.POST("/inventory", newController.DecreaseInventory())
	rtr.GET("/inventory", newController.GetInventory())

	return
}
