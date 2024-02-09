package repository

import "github.com/wyreyx/rinha-bank/internal/domain"

type Client interface {
	FindByID(id int) (*domain.Client, error)
}
