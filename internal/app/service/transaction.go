package service

import (
	"fmt"
	"strconv"

	"github.com/lccmrx/rinha-bank/internal/api/http/dto"
	"github.com/lccmrx/rinha-bank/internal/domain"
	"github.com/lccmrx/rinha-bank/internal/infra/cache"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/lccmrx/rinha-bank/internal/service"
	"github.com/pkg/errors"
)

type Transaction struct {
	service *service.Transaction
	repo    database.DataManager
	cache   cache.Cache
}

func NewTransaction(service *service.Transaction, data database.DataManager, cache cache.Cache) *Transaction {
	return &Transaction{
		service: service,
		repo:    data,
		cache:   cache,
	}
}

func (s *Transaction) Create(input dto.TransactionInput, clientID string) (client *domain.Client, err error) {
	lockKey := fmt.Sprintf("lock:%s", clientID)
	s.cache.AcquireLock(lockKey)
	defer s.cache.ReleaseLock(lockKey)

	cID, _ := strconv.Atoi(clientID)
	client, err = s.repo.Client().FindByID(cID)
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

	s.repo.Transaction().Save(transaction)
	s.repo.Client().Save(client)
	return client, nil
}
