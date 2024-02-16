package repository

import "github.com/lccmrx/rinha-bank/internal/domain"

type Client interface {
	FindByID(id string) (*domain.Client, error)
	Save(*domain.Client) error
}
