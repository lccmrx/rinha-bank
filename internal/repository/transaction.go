package repository

import "github.com/wyreyx/rinha-bank/internal/domain"

type Transaction interface {
	Save(transaction *domain.Transaction) error
}
