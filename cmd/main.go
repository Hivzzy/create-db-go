package main

import (
	"create-db-go/config"
	"create-db-go/controllers"
	"create-db-go/repository"
	"create-db-go/routes"
	"create-db-go/services"
)

func main() {
	// Menghubungkan ke database
	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	// Inisialisasi repository
	userRepo := repository.NewUserRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	transferRepo := repository.NewTransferRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db)

	// Inisialisasi service
	userService := services.NewUserService(userRepo)
	transactionService := services.NewTransactionService(transactionRepo, userRepo, transferRepo, purchaseRepo)
	transferService := services.NewTransferService(transferRepo, transactionRepo, userRepo)
	purchaseService := services.NewPurchaseService(purchaseRepo, transactionRepo, userRepo)

	// Inisialisasi controller
	userController := controllers.NewUserController(userService)
	transactionController := controllers.NewTransactionController(transactionService)
	transferController := controllers.NewTransferController(transferService)
	purchaseController := controllers.NewPurchaseController(purchaseService)

	// Mengatur router
	router := routes.SetupRouter(userController, transactionController, transferController, purchaseController)

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
