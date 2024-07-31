package services

import (
	"create-db-go/models"
	"create-db-go/repository"
	"fmt"
	"time"
)

type TransferService struct {
	transferRepo    *repository.TransferRepository
	transactionRepo *repository.TransactionRepository
	userRepo        *repository.UserRepository
}

func NewTransferService(transferRepo *repository.TransferRepository, transactionRepo *repository.TransactionRepository, userRepo *repository.UserRepository) *TransferService {
	return &TransferService{transferRepo: transferRepo, transactionRepo: transactionRepo, userRepo: userRepo}
}

func (s *TransferService) CreateTransfer(fromUserID, toUserID int, amount float64) error {
	fromUser, err := s.userRepo.GetUserByID(fromUserID)
	if err != nil {
		return err
	}

	if fromUser.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	// Transfer saldo
	err = s.userRepo.TransferAmount(fromUserID, toUserID, amount)
	if err != nil {
		return err
	}

	// Catat transfer
	transfer := &models.Transfer{
		FromUserID:   fromUserID,
		ToUserID:     toUserID,
		Amount:       amount,
		TransferDate: time.Now(),
		IsChargeback: false,
	}
	transferID, err := s.transferRepo.CreateTransfer(transfer)
	if err != nil {
		return err
	}

	// Catat transaksi untuk pengirim (User A)
	transactionFrom := &models.Transaction{
		UserID:          fromUserID,
		Amount:          -amount,
		TransactionDate: time.Now(),
		IsChargeback:    false,
		TransactionType: "transfer",
		RelatedID:       int(transferID),
	}
	err = s.transactionRepo.RecordTransaction(transactionFrom)
	if err != nil {
		return err
	}

	// Catat transaksi untuk penerima (User B)
	transactionTo := &models.Transaction{
		UserID:          toUserID,
		Amount:          amount,
		TransactionDate: time.Now(),
		IsChargeback:    false,
		TransactionType: "transfer",
		RelatedID:       int(transferID),
	}
	return s.transactionRepo.RecordTransaction(transactionTo)
}

func (s *TransferService) GetAllTransfers() ([]models.Transfer, error) {
	return s.transferRepo.GetAllTransfers()
}
