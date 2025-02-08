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

// Standardized error response
func standardErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}

// Standardized success response
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
		hubID, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid hub ID format")
			return
		}
		hub, err := c.service.FetchHubByID(ctx, hubID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusNotFound, "Hub not found")
			return
		}
		standardSuccessResponse(ctx, http.StatusOK, "Hub fetched successfully", hub)
	}
}

func (c *Controller) GetSkuByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skuID, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid SKU ID format")
			return
		}
		sku, err := c.service.FetchSkuByID(ctx, skuID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusNotFound, "SKU not found")
			return
		}
		standardSuccessResponse(ctx, http.StatusOK, "SKU fetched successfully", sku)
	}
}

// Fetch inventory details for a given SKU and Hub
func (c *Controller) GetInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skuIDParam := ctx.Query("sku_id")
		hubIDParam := ctx.Query("hub_id")

		// Validate UUID format
		skuID, err := uuid.Parse(skuIDParam)
		if err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid SKU ID format")
			return
		}

		hubID, err := uuid.Parse(hubIDParam)
		if err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid Hub ID format")
			return
		}

		// Fetch inventory from the service layer
		inventory, err := c.service.FetchInventory(ctx, skuID, hubID)
		if err != nil {
			standardErrorResponse(ctx, http.StatusNotFound, "Inventory not found")
			return
		}

		standardSuccessResponse(ctx, http.StatusOK, "Inventory fetched successfully", inventory)
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

// Update inventory quantities (decrease available, increase allocated/damaged)
func (c *Controller) DecreaseInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			SkuID uuid.UUID `json:"sku_id"`
			HubID uuid.UUID `json:"hub_id"`
			Qty   int       `json:"available_qty"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			standardErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}

		err := c.service.DecreaseInventoryQty(ctx, request.SkuID, request.HubID, request.Qty)
		if err != nil {
			standardErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		standardSuccessResponse(ctx, http.StatusOK, "Inventory updated successfully", nil)
	}
}
