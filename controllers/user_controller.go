package controllers

import (
	"create-db-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}
// @Summary Get user balance
// @Description Get balance of a specific user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user_id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{user_id}/balance [get]
func (c *UserController) GetUserBalance(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id parameter"})
		return
	}

	balance, err := c.service.GetUserBalance(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
}

// @Summary Transfer amount
// @Description Transfer amount between users
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   transfer body map[string]interface{} true "Transfer details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/transfer [post]
func (c *UserController) TransferAmount(ctx *gin.Context) {
	var request struct {
		FromUserID int     `json:"from_user_id"`
		ToUserID   int     `json:"to_user_id"`
		Amount     float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.TransferAmount(request.FromUserID, request.ToUserID, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}