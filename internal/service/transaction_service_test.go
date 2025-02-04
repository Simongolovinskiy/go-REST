package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-REST/internal/model"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Deposit(userID int, amount float64) error {
	args := m.Called(userID, amount)
	return args.Error(0)
}

func (m *MockRepository) Transfer(senderID, recipientID int, amount float64) error {
	args := m.Called(senderID, recipientID, amount)
	return args.Error(0)
}

func (m *MockRepository) GetLastTransactions(userID int) ([]model.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func TestDeposit(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewTransactionService(mockRepo)
	mockRepo.On("Deposit", 1, 100.0).Return(nil)
	err := service.Deposit(1, 100.0)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTransfer(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewTransactionService(mockRepo)
	mockRepo.On("Transfer", 1, 2, 50.0).Return(nil)
	err := service.Transfer(1, 2, 50.0)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetLastTransactions(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewTransactionService(mockRepo)
	mockRepo.On("GetLastTransactions", 1).Return([]model.Transaction{
		{ID: 1, UserID: 1, Amount: 100.0, Type: "deposit", CreatedAt: "2023-01-01"},
	}, nil)
	transactions, err := service.GetLastTransactions(1)
	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
	mockRepo.AssertExpectations(t)
}
