package service

import (
	"errors"

	"github.com/lccmrx/rinha-bank/internal/domain"
)

type Transaction struct{}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (s *Transaction) Transact(tx *domain.Transaction, client *domain.Client) (err error) {

	// it's a credit transaction (aka add credit to the client's account, increasing balance)
	if tx.TransactionType == domain.Credit {
		client.Balance += tx.Value
		return
	}

	// it's a debit transaction (aka remove credit from the client's account, decreasing balance)
	if client.Balance-tx.Value < client.AccountLimit*-1 {
		return errors.New("invalid value")
	}

	client.Balance -= tx.Value
	return
}
