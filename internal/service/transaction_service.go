package service

import (
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
)

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
