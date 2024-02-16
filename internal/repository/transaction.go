package repository

import "github.com/lccmrx/rinha-bank/internal/domain"

type Transaction interface {
	Save(transaction *domain.Transaction) error
	GetTransactionsByClientID(clientID string) ([]*domain.Transaction, error)
}
