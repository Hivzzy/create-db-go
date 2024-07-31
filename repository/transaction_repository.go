package repository

import (
	"create-db-go/models"
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetTransactions(limit, offset int, sort string) ([]models.Transaction, error) {
	query := fmt.Sprintf("SELECT id, user_id, amount, transaction_date, is_chargeback, transaction_type, related_id FROM transactions ORDER BY transaction_date %s LIMIT ? OFFSET ?", sort)
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.TransactionDate, &transaction.IsChargeback, &transaction.TransactionType, &transaction.RelatedID); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTotalTransactions() (int, error) {
	var count int
	row := r.DB.QueryRow("SELECT COUNT(*) FROM transactions")
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TransactionRepository) GetTransactionByID(transactionID int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.DB.QueryRow(
		"SELECT id, user_id, amount, transaction_date, is_chargeback, transaction_type, related_id FROM transactions WHERE id = ?",
		transactionID,
	).Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.TransactionDate, &transaction.IsChargeback, &transaction.TransactionType, &transaction.RelatedID)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) RecordTransaction(transaction *models.Transaction) error {
	_, err := r.DB.Exec(
		"INSERT INTO transactions (user_id, amount, transaction_date, is_chargeback, transaction_type, related_id) VALUES (?, ?, ?, ?, ?, ?)",
		transaction.UserID, transaction.Amount, transaction.TransactionDate, transaction.IsChargeback, transaction.TransactionType, transaction.RelatedID,
	)
	return err
}

func (r *TransactionRepository) GetTransactionsByUserID(userID, limit, offset int, sort string) ([]models.Transaction, error) {
	query := fmt.Sprintf("SELECT id, user_id, amount, transaction_date, is_chargeback, transaction_type, related_id FROM transactions WHERE user_id = ? ORDER BY transaction_date %s LIMIT ? OFFSET ?", sort)
	rows, err := r.DB.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.TransactionDate, &transaction.IsChargeback, &transaction.TransactionType, &transaction.RelatedID); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTotalTransactionsByUserID(userID int) (int, error) {
	var count int
	row := r.DB.QueryRow("SELECT COUNT(*) FROM transactions WHERE user_id = ?", userID)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TransactionRepository) GetTransactionsByDateRange(startDate, endDate string, limit, offset int, sort string) ([]models.Transaction, error) {
	query := fmt.Sprintf("SELECT id, user_id, amount, transaction_date, is_chargeback, transaction_type, related_id FROM transactions WHERE transaction_date BETWEEN ? AND ? ORDER BY transaction_date %s LIMIT ? OFFSET ?", sort)
	rows, err := r.DB.Query(query, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.TransactionDate, &transaction.IsChargeback, &transaction.TransactionType, &transaction.RelatedID); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTotalTransactionsByDateRange(startDate, endDate string) (int, error) {
	var count int
	row := r.DB.QueryRow("SELECT COUNT(*) FROM transactions WHERE transaction_date BETWEEN ? AND ?", startDate, endDate)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TransactionRepository) MarkTransactionAsChargeback(transactionID int) error {
	_, err := r.DB.Exec("UPDATE transactions SET is_chargeback = TRUE WHERE id = ?", transactionID)
	return err
}

func (r *TransactionRepository) GetTransactionsByRelatedID(relatedID int) ([]models.Transaction, error) {
	rows, err := r.DB.Query("SELECT id, user_id, amount, transaction_date, is_chargeback, transaction_type, related_id FROM transactions WHERE related_id = ?", relatedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.TransactionDate, &transaction.IsChargeback, &transaction.TransactionType, &transaction.RelatedID); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
