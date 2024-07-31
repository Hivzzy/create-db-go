package services

import (
	"create-db-go/models"
	"create-db-go/repository"
	"fmt"
)

type TransactionService struct {
	repo         *repository.TransactionRepository
	userRepo     *repository.UserRepository
	transferRepo *repository.TransferRepository
	purchaseRepo *repository.PurchaseRepository
}

func NewTransactionService(repo *repository.TransactionRepository, userRepo *repository.UserRepository, transferRepo *repository.TransferRepository, purchaseRepo *repository.PurchaseRepository) *TransactionService {
	return &TransactionService{repo: repo, userRepo: userRepo, transferRepo: transferRepo, purchaseRepo: purchaseRepo}
}

func (s *TransactionService) GetTransactions(limit, offset int, sort string) (map[int]map[string][]models.Transaction, error) {
	transactions, err := s.repo.GetTransactions(limit, offset, sort)
	if err != nil {
		return nil, err
	}

	groupedTransactions := make(map[int]map[string][]models.Transaction)
	for _, transaction := range transactions {
		relatedID := transaction.RelatedID
		if _, exists := groupedTransactions[relatedID]; !exists {
			groupedTransactions[relatedID] = make(map[string][]models.Transaction)
		}
		transactionType := transaction.TransactionType
		groupedTransactions[relatedID][transactionType] = append(groupedTransactions[relatedID][transactionType], transaction)
	}

	return groupedTransactions, nil
}

func (s *TransactionService) GetTransactionByID(transactionID int) (*models.Transaction, error) {
	return s.repo.GetTransactionByID(transactionID)
}

func (s *TransactionService) GetTotalTransactions() (int, error) {
	return s.repo.GetTotalTransactions()
}

func (s *TransactionService) GetTransactionsByUserID(userID, limit, offset int, sort string) ([]models.Transaction, error) {
	return s.repo.GetTransactionsByUserID(userID, limit, offset, sort)
}

func (s *TransactionService) GetTotalTransactionsByUserID(userID int) (int, error) {
	return s.repo.GetTotalTransactionsByUserID(userID)
}

func (s *TransactionService) GetTransactionsByDateRange(startDate, endDate string, limit, offset int, sort string) ([]models.Transaction, error) {
	return s.repo.GetTransactionsByDateRange(startDate, endDate, limit, offset, sort)
}

func (s *TransactionService) GetTotalTransactionsByDateRange(startDate, endDate string) (int, error) {
	return s.repo.GetTotalTransactionsByDateRange(startDate, endDate)
}

func (s *TransactionService) ChargebackTransaction(transactionID int) error {
	transaction, err := s.repo.GetTransactionByID(transactionID)
	if err != nil {
		return err
	}

	if transaction.IsChargeback {
		return fmt.Errorf("transaction already marked as chargeback")
	}

	// Get all related transactions
	relatedTransactions, err := s.repo.GetTransactionsByRelatedID(transaction.RelatedID)
	if err != nil {
		return err
	}

	// Handle chargeback based on transaction type for all related transactions
	for _, relatedTransaction := range relatedTransactions {
		if relatedTransaction.TransactionType == "transfer" {
			// Get the related transfer
			transfer, err := s.transferRepo.GetTransferByID(relatedTransaction.RelatedID)
			if err != nil {
				return err
			}

			// Return money from to_user_id to from_user_id
			err = s.userRepo.UpdateUserBalance(transfer.ToUserID, -transfer.Amount)
			if err != nil {
				return err
			}

			err = s.userRepo.UpdateUserBalance(transfer.FromUserID, transfer.Amount)
			if err != nil {
				return err
			}
		} else if relatedTransaction.TransactionType == "purchase" {
			// Get the related purchase
			purchase, err := s.purchaseRepo.GetPurchaseByID(relatedTransaction.RelatedID)
			if err != nil {
				return err
			}

			// Return money to the user
			err = s.userRepo.UpdateUserBalance(purchase.UserID, purchase.Amount)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unsupported transaction type for chargeback")
		}

		// Mark each related transaction as chargeback
		err = s.repo.MarkTransactionAsChargeback(relatedTransaction.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
