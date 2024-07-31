package models

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
	IsChargeback    bool      `json:"is_chargeback"`
	TransactionType string    `json:"transaction_type"`
	RelatedID       int       `json:"related_id"`
}
