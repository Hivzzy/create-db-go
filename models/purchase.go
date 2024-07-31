package models

import "time"

type Purchase struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	ItemName     string    `json:"item_name"`
	Amount       float64   `json:"amount"`
	PurchaseDate time.Time `json:"purchase_date"`
	IsChargeback bool      `json:"is_chargeback"`
}
