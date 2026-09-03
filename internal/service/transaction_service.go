package service

import (
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
)

var ErrTransactionNotFound = errors.New("transaction not found")
var ErrTransactionAlreadyCancelled = errors.New("transaction already cancelled")

type TransactionService struct {
	tr *repository.TransactionRepository
}

func NewTransactionService(tr *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		tr: tr,
	}
}

func (t *TransactionService) FetchTransactionByUserId(userId string) ([]model.Transaction, error) {
	result, err := t.tr.GetTransactionByUserId(userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

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
