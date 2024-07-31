package routes

import (
	"create-db-go/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, transactionController *controllers.TransactionController, transferController *controllers.TransferController, purchaseController *controllers.PurchaseController) *gin.Engine {
	router := gin.Default()

	// Transactions Route
	router.GET("/transactions", transactionController.GetTransactions)
	router.GET("/transactions/user/:user_id", transactionController.GetTransactionsByUserID)
	router.GET("/transactions/date-range", transactionController.GetTransactionsByDateRange)
	router.POST("/transactions/:transaction_id/chargeback", transactionController.ChargebackTransaction)
	router.GET("/transactions/:transaction_id", transactionController.GetTransactionByID) // Route baru

	// User Route
	router.GET("/users/:user_id/balance", userController.GetUserBalance)

	// Transfer Route
	router.POST("/transfers", transferController.CreateTransfer)
	router.GET("/transfers", transferController.GetAllTransfers) // Route baru

	// Purchase Route
	router.POST("/purchases", purchaseController.CreatePurchase)
	router.GET("/purchases", purchaseController.GetAllPurchases) // Route baru

	return router
}
