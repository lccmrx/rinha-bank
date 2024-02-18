package model

import (
	"errors"

	"github.com/lccmrx/rinha-bank/internal/domain"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
)

type Client struct {
	Conn database.Executor
}

type ClientModel struct {
	ID      int `db:"id"`
	Limit   int `db:"limit"`
	Balance int `db:"balance"`
}

func (c *Client) FindByID(id string) (client *domain.Client, err error) {
	var models []ClientModel
	err = c.Conn.Select(&models, "SELECT * FROM client WHERE id = $1 for update", id)
	if err != nil {
		return nil, err
	}

	if len(models) == 0 {
		return nil, errors.New("not found")
	}

	model := models[0]

	client = &domain.Client{
		ID:      model.ID,
		Limit:   model.Limit,
		Balance: model.Balance,
	}

	return client, nil
}

func (c *Client) Save(client *domain.Client) error {
	_, err := c.Conn.Exec("UPDATE client SET balance = $1 WHERE id = $2", client.Balance, client.ID)
	if err != nil {
		return err
	}

	return nil
}
