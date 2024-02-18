package app

import (
	"github.com/lccmrx/rinha-bank/internal/api/http"
	usecase "github.com/lccmrx/rinha-bank/internal/app/service"
	"github.com/lccmrx/rinha-bank/internal/service"

	"github.com/lccmrx/rinha-bank/internal/infra/config"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/lccmrx/rinha-bank/internal/infra/database/postgres"
	"go.uber.org/fx"
)

func providers() []any {
	globalDeps := []any{
		usecase.NewTransaction,
		config.New,
		http.NewControllerManager,
		fx.Annotate(postgres.Instance, fx.As(new(database.DataManager))),

		usecase.NewStatement,

		service.NewTransaction,
	}

	globalDeps = append(globalDeps, http.ControllersList...)
	return globalDeps
}
