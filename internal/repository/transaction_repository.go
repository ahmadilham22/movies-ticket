package repository

import (
	"online-ticketing/internal/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (tr *TransactionRepository) GetTransactionByUserId(userId string) ([]model.Transaction, error) {
	transaction := []model.Transaction{}
	err := tr.db.Select(&transaction, "SELECT * FROM transactions WHERE user_id = $1 ORDER BY created_at asc", userId)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
