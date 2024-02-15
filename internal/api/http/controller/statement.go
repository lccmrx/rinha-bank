package controller

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lccmrx/rinha-bank/internal/api/http/dto"
	"github.com/lccmrx/rinha-bank/internal/app/service"
)

type Statement struct {
	service *service.Statement
}

func NewStatement(service *service.Statement) *Statement {
	return &Statement{
		service: service,
	}
}

func (s *Statement) Get(ctx echo.Context) (err error) {
	var clientID = ctx.Param("id")

	client, transactions, err := s.service.GetStatement(clientID)

	if err != nil {
		return ctx.String(404, err.Error())
	}

	var transactionsOutput []dto.StatementTransactionOutput
	for _, t := range transactions {
		transaction := dto.StatementTransactionOutput{
			Value:       t.Value,
			Type:        string(t.TransactionType),
			Description: t.Description,
			Timestamp:   t.Timestamp.Format(time.RFC3339Nano),
		}
		transactionsOutput = append(transactionsOutput, transaction)
	}

	return ctx.JSON(200, dto.StatementOutput{
		Balance: dto.StatementBalanceOutput{
			Total:         client.Balance,
			StatementDate: time.Now().Format(time.RFC3339Nano),
			Limit:         client.AccountLimit,
		},
		Transactions: transactionsOutput,
	})
}
