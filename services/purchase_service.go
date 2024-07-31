package services

import (
	"create-db-go/models"
	"create-db-go/repository"
	"fmt"
	"time"
)

type PurchaseService struct {
	purchaseRepo    *repository.PurchaseRepository
	transactionRepo *repository.TransactionRepository
	userRepo        *repository.UserRepository
}

func NewPurchaseService(purchaseRepo *repository.PurchaseRepository, transactionRepo *repository.TransactionRepository, userRepo *repository.UserRepository) *PurchaseService {
	return &PurchaseService{purchaseRepo: purchaseRepo, transactionRepo: transactionRepo, userRepo: userRepo}
}

func (s *PurchaseService) CreatePurchase(userID int, itemName string, amount float64) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	// Kurangi saldo pengguna
	err = s.userRepo.UpdateUserBalance(userID, -amount)
	if err != nil {
		return err
	}

	// Catat pembelian
	purchase := &models.Purchase{
		UserID:       userID,
		ItemName:     itemName,
		Amount:       amount,
		PurchaseDate: time.Now(),
		IsChargeback: false,
	}
	purchaseID, err := s.purchaseRepo.CreatePurchase(purchase)
	if err != nil {
		return err
	}

	// Catat transaksi
	transaction := &models.Transaction{
		UserID:          userID,
		Amount:          amount,
		TransactionDate: time.Now(),
		IsChargeback:    false,
		TransactionType: "purchase",
		RelatedID:       int(purchaseID),
	}
	return s.transactionRepo.RecordTransaction(transaction)
}

func (s *PurchaseService) GetAllPurchases() ([]models.Purchase, error) {
	return s.purchaseRepo.GetAllPurchases()
}
