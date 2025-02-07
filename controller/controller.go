package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"wms/domain"
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

// Standardize Error Response for API
func standardErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}

// Standardize Success Response for API
func standardSuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func (c *Controller) GetHubs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubs, err := c.service.FetchHubs(ctx)
		if err != nil {
			standardErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch hubs")
			return
		}
		standardSuccessResponse(ctx, http.StatusOK, "Hubs fetched successfully", hubs)
	}
}

func (c *Controller) GetSkus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skus, err := c.service.FetchSkus(ctx)
		if err != nil {
			standardErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch SKUs")
			return
		}
		standardSuccessResponse(ctx, http.StatusOK, "SKUs fetched successfully", skus)
	}
}

func (c *Controller) GetHubByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubID := ctx.Param("id")
		parsedHubID, err := uuid.Parse(hubID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid hub ID format")
			return
		}
		hub, err := c.service.FetchHubByID(ctx, parsedHubID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusNotFound, "Hub not found")
			return
		}
		standardSuccessResponse(ctx, http.StatusOK, "Hub fetched successfully", hub)
	}
}

func (c *Controller) GetSkuByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skuID := ctx.Param("id")
		parsedSkuID, err := uuid.Parse(skuID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid SKU ID format")
			return
		}
		sku, err := c.service.FetchSkuByID(ctx, parsedSkuID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusNotFound, "SKU not found")
			return
		}
		standardSuccessResponse(ctx, http.StatusOK, "SKU fetched successfully", sku)
	}
}

// POST API to create Hub
func (c *Controller) CreateHub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hub domain.Hub
		if err := ctx.ShouldBindJSON(&hub); err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}

		err := c.service.CreateHub(ctx, hub)
		if err != nil {
			standardErrorResponse(ctx, http.StatusInternalServerError, "Failed to create hub")
			return
		}
		standardSuccessResponse(ctx, http.StatusCreated, "Hub created successfully", nil)
	}
}

// POST API to create SKU
func (c *Controller) CreateSKU() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sku domain.SKU
		if err := ctx.ShouldBindJSON(&sku); err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}

		err := c.service.CreateSKU(ctx, sku)
		if err != nil {
			standardErrorResponse(ctx, http.StatusInternalServerError, "Failed to create SKU")
			return
		}
		standardSuccessResponse(ctx, http.StatusCreated, "SKU created successfully", nil)
	}
}

func (c *Controller) DecreaseInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			SkuID        uuid.UUID `json:"sku_id"`
			HubID        uuid.UUID `json:"hub_id"`
			AvailableQty int       `json:"available_qty"`
			AllocatedQty int       `json:"allocated_qty"`
			DamagedQty   int       `json:"damaged_qty"`
		}

		// Bind JSON body to the request struct
		if err := ctx.ShouldBindJSON(&request); err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}

		// Call service to decrease inventory quantities
		err := c.service.DecreaseInventoryQty(ctx, request.SkuID, request.HubID, request.AvailableQty, request.AllocatedQty, request.DamagedQty)
		if err != nil {
			standardErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		standardSuccessResponse(ctx, http.StatusOK, "Inventory updated successfully", nil)
	}
}
