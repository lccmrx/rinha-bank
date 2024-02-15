package model

import (
	"time"

	"github.com/lccmrx/rinha-bank/internal/domain"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
)

type Transaction struct {
	Conn database.Executor
}

type TransactionModel struct {
	ClientID        int    `db:"client_id"`
	Value           int    `db:"value"`
	TransactionType string `db:"type"`
	Description     string `db:"description"`
	Timestamp       string `db:"timestamp"`
}

func (t *Transaction) Save(tx *domain.Transaction) error {
	_, err := t.Conn.Exec("INSERT INTO transaction (client_id, value, type, description) VALUES ($1, $2, $3, $4)", tx.ClientID, tx.Value, tx.TransactionType, tx.Description)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) GetTransactionsByClientID(clientID int) (transactions []*domain.Transaction, err error) {
	var models []TransactionModel
	err = t.Conn.Select(&models, "SELECT * FROM transaction WHERE client_id = $1 order by timestamp desc limit 10", clientID)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		parsedTime, _ := time.Parse(time.RFC3339Nano, model.Timestamp)

		transaction := &domain.Transaction{
			ClientID:        model.ClientID,
			Value:           model.Value,
			TransactionType: domain.TransactionType(model.TransactionType),
			Description:     model.Description,
			Timestamp:       parsedTime,
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
