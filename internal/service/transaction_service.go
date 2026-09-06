package service

import (
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
)

var ErrTransactionNotFound = errors.New("transaction not found")
var ErrTransactionAlreadyCancelled = errors.New("transaction already cancelled")

// Interface for transaction repository
type transactionRepository interface {
	GetTransactionByUserId(userId string) ([]model.Transaction, error)
	CancelTransaction(bookingCode, userId string) error
}

// Dependency Injection for TransactionService
type TransactionService struct {
	tr transactionRepository
}

// NewTransactionService creates a new TransactionService
func NewTransactionService(tr transactionRepository) *TransactionService {
	return &TransactionService{
		tr: tr,
	}
}

// Method FetchTransactionByUserId fetches transactions by user ID
func (t *TransactionService) FetchTransactionByUserId(userId string) ([]model.Transaction, error) {
	result, err := t.tr.GetTransactionByUserId(userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Method CancelTransaction cancels a transaction
func (t *TransactionService) CancelTransaction(bookingCode, userId string) error {
	err := t.tr.CancelTransaction(bookingCode, userId)
	if err != nil {
		if errors.Is(err, repository.ErrTransactionNotFound) {
			return ErrTransactionNotFound
		}
		if errors.Is(err, repository.ErrCancelTransaction) {
			return ErrTransactionAlreadyCancelled
		}
		return err
	}

	return nil
}
