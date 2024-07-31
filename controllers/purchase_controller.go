package controllers

import (
	"create-db-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	service *services.PurchaseService
}

func NewPurchaseController(service *services.PurchaseService) *PurchaseController {
	return &PurchaseController{service: service}
}

// @Summary Create a new purchase
// @Description Create a new purchase for a user
// @Tags purchases
// @Accept  json
// @Produce  json
// @Param purchase body map[string]interface{} true "Purchase details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchases [post]
func (c *PurchaseController) CreatePurchase(ctx *gin.Context) {
	var request struct {
		UserID   int     `json:"user_id"`
		ItemName string  `json:"item_name"`
		Amount   float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.CreatePurchase(request.UserID, request.ItemName, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Purchase successful"})
}

// @Summary Get all purchases
// @Description Get details of all purchases
// @Tags purchases
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Purchase
// @Failure 500 {object} map[string]string
// @Router /purchases [get]
func (c *PurchaseController) GetAllPurchases(ctx *gin.Context) {
	purchases, err := c.service.GetAllPurchases()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, purchases)
}
