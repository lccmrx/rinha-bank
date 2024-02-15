package app

import (
	"github.com/lccmrx/rinha-bank/internal/api/http"
	usecase "github.com/lccmrx/rinha-bank/internal/app/service"
	"github.com/lccmrx/rinha-bank/internal/service"

	"github.com/lccmrx/rinha-bank/internal/infra/cache"
	"github.com/lccmrx/rinha-bank/internal/infra/cache/redis"
	"github.com/lccmrx/rinha-bank/internal/infra/config"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/lccmrx/rinha-bank/internal/infra/database/postgres"
	"go.uber.org/fx"
)

func providers() []any {
	globalDeps := []any{
		config.New,
		http.NewControllerManager,
		fx.Annotate(redis.New, fx.As(new(cache.Cache))),
		fx.Annotate(postgres.Instance, fx.As(new(database.DataManager))),

		usecase.NewTransaction,
		usecase.NewStatement,

		service.NewTransaction,
	}

	globalDeps = append(globalDeps, http.ControllersList...)
	return globalDeps
}
