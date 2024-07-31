package repository

import (
	"create-db-go/models"
	"database/sql"
)

type TransferRepository struct {
	DB *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{DB: db}
}

func (r *TransferRepository) CreateTransfer(transfer *models.Transfer) (int64, error) {
	result, err := r.DB.Exec(
		"INSERT INTO transfers (from_user_id, to_user_id, amount, transfer_date, is_chargeback) VALUES (?, ?, ?, ?, ?)",
		transfer.FromUserID, transfer.ToUserID, transfer.Amount, transfer.TransferDate, transfer.IsChargeback,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *TransferRepository) GetTransferByID(id int) (*models.Transfer, error) {
	var transfer models.Transfer
	err := r.DB.QueryRow(
		"SELECT id, from_user_id, to_user_id, amount, transfer_date, is_chargeback FROM transfers WHERE id = ?",
		id,
	).Scan(&transfer.ID, &transfer.FromUserID, &transfer.ToUserID, &transfer.Amount, &transfer.TransferDate, &transfer.IsChargeback)
	if err != nil {
		return nil, err
	}
	return &transfer, nil
}

func (r *TransferRepository) GetAllTransfers() ([]models.Transfer, error) {
	rows, err := r.DB.Query("SELECT id, from_user_id, to_user_id, amount, transfer_date, is_chargeback FROM transfers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []models.Transfer{}
	for rows.Next() {
		var transfer models.Transfer
		if err := rows.Scan(&transfer.ID, &transfer.FromUserID, &transfer.ToUserID, &transfer.Amount, &transfer.TransferDate, &transfer.IsChargeback); err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
