package service

import (
	"github.com/lccmrx/rinha-bank/internal/api/http/dto"
	"github.com/lccmrx/rinha-bank/internal/domain"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/lccmrx/rinha-bank/internal/service"
	"github.com/pkg/errors"
)

type Transaction struct {
	service *service.Transaction
	repo    database.DataManager
}

func NewTransaction(service *service.Transaction, data database.DataManager) *Transaction {
	return &Transaction{
		service: service,
		repo:    data,
	}
}

func (s *Transaction) Create(input dto.TransactionInput, clientID string) (client *domain.Client, err error) {
	tx, err := s.repo.Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}

	client, err = tx.Client().FindByID(clientID)
	if err != nil {
		return nil, errors.Wrap(err, "not found")
	}

	transaction, err := domain.NewTransaction(domain.TransactionParams{
		ClientID:        client.ID,
		Value:           input.Value,
		TransactionType: input.Type,
		Description:     input.Description,
	})
	if err != nil {
		return nil, errors.Wrap(err, "invalid")
	}

	err = s.service.Transact(transaction, client)
	if err != nil {
		return nil, errors.Wrap(err, "invalid")
	}

	tx.Transaction().Save(transaction)
	tx.Client().Save(client)

	tx.Commit()
	return client, nil
}
