package service

import (
	"errors"
	"go-REST/internal/model"
	"go-REST/internal/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Deposit(userID int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	return s.repo.Deposit(userID, amount)
}

func (s *TransactionService) Transfer(senderID, recipientID int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	return s.repo.Transfer(senderID, recipientID, amount)
}

func (s *TransactionService) GetLastTransactions(userID int) ([]model.Transaction, error) {
	return s.repo.GetLastTransactions(userID)
}
