package repository

import (
	"create-db-go/models"
	"database/sql"
)

type PurchaseRepository struct {
	DB *sql.DB
}

func NewPurchaseRepository(db *sql.DB) *PurchaseRepository {
	return &PurchaseRepository{DB: db}
}

func (r *PurchaseRepository) CreatePurchase(purchase *models.Purchase) (int64, error) {
	result, err := r.DB.Exec(
		"INSERT INTO purchases (user_id, item_name, amount, purchase_date, is_chargeback) VALUES (?, ?, ?, ?, ?)",
		purchase.UserID, purchase.ItemName, purchase.Amount, purchase.PurchaseDate, purchase.IsChargeback,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *PurchaseRepository) GetPurchaseByID(id int) (*models.Purchase, error) {
	var purchase models.Purchase
	err := r.DB.QueryRow(
		"SELECT id, user_id, item_name, amount, purchase_date, is_chargeback FROM purchases WHERE id = ?",
		id,
	).Scan(&purchase.ID, &purchase.UserID, &purchase.ItemName, &purchase.Amount, &purchase.PurchaseDate, &purchase.IsChargeback)
	if err != nil {
		return nil, err
	}
	return &purchase, nil
}

func (r *PurchaseRepository) GetAllPurchases() ([]models.Purchase, error) {
	rows, err := r.DB.Query("SELECT id, user_id, item_name, amount, purchase_date, is_chargeback FROM purchases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	purchases := []models.Purchase{}
	for rows.Next() {
		var purchase models.Purchase
		if err := rows.Scan(&purchase.ID, &purchase.UserID, &purchase.ItemName, &purchase.Amount, &purchase.PurchaseDate, &purchase.IsChargeback); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}
