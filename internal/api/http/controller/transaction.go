package controller

import "github.com/labstack/echo/v4"

type Transaction struct {
	// service service.Transaction
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) Create(ctx echo.Context) (err error) {
	// t.service.Create()
	return
}
