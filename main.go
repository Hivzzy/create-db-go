package main

import (
	"create-db-go/config"
	"create-db-go/controllers"
	_ "create-db-go/docs" // Ganti dengan path ke package docs Anda
	"create-db-go/repository"
	"create-db-go/routes"
	"create-db-go/services"
)

func main() {
    db, err := config.ConnectDatabase()
    if err != nil {
        panic(err)
    }

    userRepo := repository.NewUserRepository(db)
    transactionRepo := repository.NewTransactionRepository(db)
    transferRepo := repository.NewTransferRepository(db)
    purchaseRepo := repository.NewPurchaseRepository(db)

    userService := services.NewUserService(userRepo)
    transactionService := services.NewTransactionService(transactionRepo, userRepo, transferRepo, purchaseRepo)
    transferService := services.NewTransferService(transferRepo, transactionRepo, userRepo)
    purchaseService := services.NewPurchaseService(purchaseRepo, transactionRepo, userRepo)

    userController := controllers.NewUserController(userService)
    transactionController := controllers.NewTransactionController(transactionService)
    transferController := controllers.NewTransferController(transferService)
    purchaseController := controllers.NewPurchaseController(purchaseService)

    router := routes.SetupRouter(userController, transactionController, transferController, purchaseController)

    if err := router.Run(":8080"); err != nil {
        panic(err)
    }
}
