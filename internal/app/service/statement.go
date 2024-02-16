package service

import (
	"fmt"

	"github.com/lccmrx/rinha-bank/internal/domain"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/pkg/errors"
)

type Statement struct {
	repo database.DataManager
}

func NewStatement(data database.DataManager) *Statement {
	return &Statement{
		repo: data,
	}
}

func (s *Statement) GetStatement(clientID string) (client *domain.Client, transactions []*domain.Transaction, err error) {
	client, err = s.repo.Client().FindByID(clientID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "not found")
	}

	transactions, err = s.repo.Transaction().GetTransactionsByClientID(fmt.Sprintf("%d", client.ID))
	if err != nil {
		return nil, nil, errors.Wrap(err, "not found")
	}

	return client, transactions, nil
}
