package service

import (
	"fmt"
	"strconv"

	"github.com/lccmrx/rinha-bank/internal/domain"
	"github.com/lccmrx/rinha-bank/internal/infra/cache"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/pkg/errors"
)

type Statement struct {
	repo  database.DataManager
	cache cache.Cache
}

func NewStatement(data database.DataManager, cache cache.Cache) *Statement {
	return &Statement{
		repo:  data,
		cache: cache,
	}
}

func (s *Statement) GetStatement(clientID string) (client *domain.Client, transactions []*domain.Transaction, err error) {
	lockKey := fmt.Sprintf("lock:%s", clientID)
	s.cache.AcquireLock(lockKey)
	defer s.cache.ReleaseLock(lockKey)

	cID, _ := strconv.Atoi(clientID)
	client, err = s.repo.Client().FindByID(cID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "not found")
	}

	transactions, err = s.repo.Transaction().GetTransactionsByClientID(client.ID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "not found")
	}

	return client, transactions, nil
}
