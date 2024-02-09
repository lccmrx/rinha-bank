package app

import (
	"context"

	"github.com/wyreyx/rinha-bank/internal/api"
	"go.uber.org/fx"
)

var Version string = "v0.0.1"

func Start(ctx context.Context, starter any) {
	fx.New(
		// disables fx logger
		fx.NopLogger,

		fx.Provide(
			// provide app context
			func() context.Context {
				return ctx
			},

			// provide starter API handler
			starter,
		),

		fx.Provide(
			providers()...,
		),

		fx.Invoke(
			hook(),
		),
	).Run()
}

func hook() any {
	return func(ctx context.Context, lc fx.Lifecycle, api api.Api) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go api.Run(ctx)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return api.Shutdown(ctx)
			},
		})
	}
}
