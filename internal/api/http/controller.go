package http

import "github.com/lccmrx/rinha-bank/internal/api/http/controller"

type ControllerManager struct {
	transaction *controller.Transaction
	statement   *controller.Statement
}

var ControllersList = []any{
	controller.NewTransaction,
	controller.NewStatement,
}

func NewControllerManager(
	transaction *controller.Transaction,
	statement *controller.Statement,
) *ControllerManager {
	return &ControllerManager{
		transaction: transaction,
		statement:   statement,
	}
}
