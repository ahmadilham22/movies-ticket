package service

import (
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
	"testing"
)

type fakeTransactionRepository struct {
	cancelErr error
}

func (f *fakeTransactionRepository) GetTransactionByUserId(_ string) ([]model.Transaction, error) {
	return nil, nil
}

func (f *fakeTransactionRepository) CancelTransaction(_, _ string) error {
	return f.cancelErr
}

func TestTransactionServiceCancelTransaction(t *testing.T) {
	unexpectedErr := errors.New("unexpected error")
	tests := []struct {
		name    string
		repoErr error
		wantErr error
	}{
		{
			name:    "success",
			repoErr: nil,
			wantErr: nil,
		},
		{
			name:    "transaction not found",
			repoErr: repository.ErrTransactionNotFound,
			wantErr: ErrTransactionNotFound,
		},
		{
			name:    "transaction already canceled",
			repoErr: repository.ErrCancelTransaction,
			wantErr: ErrTransactionAlreadyCancelled,
		}, {
			name:    "unexpected error",
			repoErr: unexpectedErr,
			wantErr: unexpectedErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fakeRepo := fakeTransactionRepository{
				cancelErr: tc.repoErr,
			}
			transService := NewTransactionService(&fakeRepo)
			gotErr := transService.CancelTransaction("booking-test", "user-test")
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("CancelTransaction() error = %v, want %v", gotErr, tc.wantErr)
			}
		})
	}

}
