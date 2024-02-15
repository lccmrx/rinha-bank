package controller

import (
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lccmrx/rinha-bank/internal/api/http/dto"
	"github.com/lccmrx/rinha-bank/internal/app/service"
)

type Transaction struct {
	service *service.Transaction
}

func NewTransaction(service *service.Transaction) *Transaction {
	return &Transaction{
		service: service,
	}
}

func (c *Transaction) Create(ctx echo.Context) (err error) {
	var input = dto.TransactionInput{}
	err = ctx.Bind(&input)
	if err != nil {
		return ctx.JSON(422, err.Error())
	}

	var clientID = ctx.Param("id")

	client, err := c.service.Create(input, clientID)
	if err != nil {
		errSplitted := strings.Split(err.Error(), ":")
		if slices.Contains(errSplitted, "not found") {
			return ctx.String(404, err.Error())
		}

		if slices.Contains(errSplitted, "invalid") {
			return ctx.String(422, err.Error())
		}

		return ctx.String(500, err.Error())
	}

	return ctx.JSON(200, dto.TransactionOutput{
		Limit:   client.AccountLimit,
		Balance: client.Balance,
	})
}
