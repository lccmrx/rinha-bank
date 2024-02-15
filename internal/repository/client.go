package repository

import "github.com/lccmrx/rinha-bank/internal/domain"

type Client interface {
	FindByID(id int) (*domain.Client, error)
}
