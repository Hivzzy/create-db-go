package controllers

import (
	"create-db-go/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// @Summary Get transactions
// @Description Get a list of transactions with pagination
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   limit query int false "Limit"
// @Param   page query int false "Page"
// @Param   sort query string false "Sort" Enums(asc, desc)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
func (c *TransactionController) GetTransactions(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	sort := ctx.DefaultQuery("sort", "asc")
	if sort != "asc" && sort != "desc" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort parameter"})
		return
	}

	offset := (page - 1) * limit

	groupedTransactions, err := c.service.GetTransactions(limit, offset, sort)
	if err != nil {
		log.Printf("Error getting transactions: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := c.service.GetTotalTransactions()
	if err != nil {
		log.Printf("Error getting total transactions: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data":       groupedTransactions,
		"total":      total,
		"limit":      limit,
		"page":       page,
		"sort":       sort,
		"totalPages": (total + limit - 1) / limit,
	}

	log.Printf("Retrieved %d transactions", len(groupedTransactions))
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get transactions by user ID
// @Description Get transactions for a specific user with pagination
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   user_id path int true "User ID"
// @Param   limit query int false "Limit"
// @Param   page query int false "Page"
// @Param   sort query string false "Sort" Enums(asc, desc)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/user/{user_id} [get]
func (c *TransactionController) GetTransactionsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id parameter"})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	sort := ctx.DefaultQuery("sort", "asc")
	if sort != "asc" && sort != "desc" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort parameter"})
		return
	}

	offset := (page - 1) * limit

	transactions, err := c.service.GetTransactionsByUserID(userID, limit, offset, sort)
	if err != nil {
		log.Printf("Error getting transactions for user %d: %v", userID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := c.service.GetTotalTransactionsByUserID(userID)
	if err != nil {
		log.Printf("Error getting total transactions for user %d: %v", userID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data":       transactions,
		"total":      total,
		"limit":      limit,
		"page":       page,
		"sort":       sort,
		"totalPages": (total + limit - 1) / limit,
	}

	log.Printf("Retrieved %d transactions for user %d", len(transactions), userID)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get transactions by date range
// @Description Get transactions within a specific date range
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   start_date query string true "Start date"
// @Param   end_date query string true "End date"
// @Param   limit query int false "Limit"
// @Param   page query int false "Page"
// @Param   sort query string false "Sort" Enums(asc, desc)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/date [get]
func (c *TransactionController) GetTransactionsByDateRange(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	if startDate == "" || endDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	sort := ctx.DefaultQuery("sort", "asc")
	if sort != "asc" && sort != "desc" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort parameter"})
		return
	}

	offset := (page - 1) * limit

	transactions, err := c.service.GetTransactionsByDateRange(startDate, endDate, limit, offset, sort)
	if err != nil {
		log.Printf("Error getting transactions by date range: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := c.service.GetTotalTransactionsByDateRange(startDate, endDate)
	if err != nil {
		log.Printf("Error getting total transactions by date range: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data":       transactions,
		"total":      total,
		"limit":      limit,
		"page":       page,
		"sort":       sort,
		"totalPages": (total + limit - 1) / limit,
	}

	log.Printf("Retrieved %d transactions between %s and %s", len(transactions), startDate, endDate)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Chargeback a transaction
// @Description Perform a chargeback on a specific transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   transaction_id path int true "Transaction ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/chargeback/{transaction_id} [post]
func (c *TransactionController) ChargebackTransaction(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("transaction_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction_id parameter"})
		return
	}

	err = c.service.ChargebackTransaction(transactionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Chargeback successful"})
}

// @Summary Get transaction by ID
// @Description Get details of a specific transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   transaction_id path int true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{transaction_id} [get]
func (c *TransactionController) GetTransactionByID(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("transaction_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction_id parameter"})
		return
	}

	transaction, err := c.service.GetTransactionByID(transactionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
