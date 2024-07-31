package controllers

import (
	"create-db-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	service *services.TransferService
}

func NewTransferController(service *services.TransferService) *TransferController {
	return &TransferController{service: service}
}

// @Summary Create a new transfer
// @Description Create a transfer between users
// @Tags transfers
// @Accept  json
// @Produce  json
// @Param transfer body map[string]interface{} true "Transfer details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transfers [post]
func (c *TransferController) CreateTransfer(ctx *gin.Context) {
	var request struct {
		FromUserID int     `json:"from_user_id"`
		ToUserID   int     `json:"to_user_id"`
		Amount     float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.CreateTransfer(request.FromUserID, request.ToUserID, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}

// @Summary Get all transfers
// @Description Get details of all transfers
// @Tags transfers
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Transfer
// @Failure 500 {object} map[string]string
// @Router /transfers [get]
func (c *TransferController) GetAllTransfers(ctx *gin.Context) {
	transfers, err := c.service.GetAllTransfers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transfers)
}
