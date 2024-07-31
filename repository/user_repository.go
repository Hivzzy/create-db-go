package repository

import (
	"create-db-go/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email, balance, created_at FROM users WHERE id = ?", userID)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Balance, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserBalance(userID int, amount float64) error {
	_, err := r.DB.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, userID)
	return err
}

func (r *UserRepository) GetUserBalance(userID int) (float64, error) {
	var balance float64
	err := r.DB.QueryRow("SELECT balance FROM users WHERE id = ?", userID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r *UserRepository) TransferAmount(fromUserID, toUserID int, amount float64) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	// Kurangi saldo pengirim
	_, err = tx.Exec("UPDATE users SET balance = balance - ? WHERE id = ?", amount, fromUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Tambah saldo penerima
	_, err = tx.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, toUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
