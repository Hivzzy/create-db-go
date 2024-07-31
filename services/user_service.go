package services

import (
	"create-db-go/models"
	"create-db-go/repository"
	"fmt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *UserService) UpdateUserBalance(userID int, amount float64) error {
	return s.userRepo.UpdateUserBalance(userID, amount)
}

func (s *UserService) GetUserBalance(userID int) (float64, error) {
	return s.userRepo.GetUserBalance(userID)
}

func (s *UserService) TransferAmount(fromUserID, toUserID int, amount float64) error {
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

	return nil
}
