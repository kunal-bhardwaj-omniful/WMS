package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wms/service"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) GetHubs() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		hubs := c.service.FetchHubs(ctx)
		//if err != nil {
		//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hubs"})
		//	return
		//}

		ctx.JSON(http.StatusOK, hubs)

	}
}
