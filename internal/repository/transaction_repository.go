package repository

import (
	"database/sql"
	"errors"
	"online-ticketing/internal/model"

	"github.com/jmoiron/sqlx"
)

var ErrTransactionNotFound = errors.New("transaction not found")
var ErrCancelTransaction = errors.New("can not cancel transaction")

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

func (tr *TransactionRepository) CancelTransaction(bookingCode string, userId string) error {
	transactionData := model.Transaction{}
	tx, err := tr.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.Get(&transactionData, "SELECT * FROM transactions WHERE user_id = $1 AND booking_code = $2 FOR UPDATE", userId, bookingCode)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrTransactionNotFound
		}
		return err
	}

	if transactionData.Status != "SUCCESS" {
		return ErrCancelTransaction
	}

	_, err = tx.Exec("UPDATE tickets SET quota = quota + $1 WHERE id = $2", transactionData.Quantity, transactionData.TicketID)

	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE transactions SET status = 'CANCELLED', updated_at = CURRENT_TIMESTAMP WHERE booking_code = $1 AND user_id = $2", bookingCode, userId)

	if err != nil {
		return err
	}

	return tx.Commit()

}
