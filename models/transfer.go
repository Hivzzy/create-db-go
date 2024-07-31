package models

import "time"

type Transfer struct {
	ID           int       `json:"id"`
	FromUserID   int       `json:"from_user_id"`
	ToUserID     int       `json:"to_user_id"`
	Amount       float64   `json:"amount"`
	TransferDate time.Time `json:"transfer_date"`
	IsChargeback bool      `json:"is_chargeback"`
}
