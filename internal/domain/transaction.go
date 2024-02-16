package domain

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Debit  TransactionType = "d"
	Credit TransactionType = "c"
)

var validTransactionTypes = map[string]TransactionType{
	"d": Debit,
	"c": Credit,
}

type Transaction struct {
	ClientID        int             `json:"client-id"`
	TransactionType TransactionType `json:"transaction-type"`
	Value           int             `json:"value"`
	Description     string          `json:"description"`
	Timestamp       time.Time       `json:"timestamp"`
}

type TransactionParams struct {
	ClientID        int    `json:"client-id"`
	TransactionType string `json:"transaction-type"`
	Value           int    `json:"value"`
	Description     string `json:"description"`
}

func (t *TransactionParams) Validate() error {
	if _, ok := validTransactionTypes[t.TransactionType]; !ok {
		return errors.New("invalid transaction type")
	}

	if t.Value <= 0 {
		return errors.New("value must be greater than 0")
	}

	if t.Description == "" || len(t.Description) > 10 {
		return errors.New("description must not be empty or bigger than 10 characters")
	}

	return nil
}

func NewTransaction(params TransactionParams) (*Transaction, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	return &Transaction{
		ClientID:        params.ClientID,
		TransactionType: TransactionType(params.TransactionType),
		Value:           params.Value,
		Description:     params.Description,
		Timestamp:       time.Now(),
	}, nil
}
