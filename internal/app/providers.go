package app

import (
	"github.com/lccmrx/rinha-bank/internal/api/http"
	"github.com/lccmrx/rinha-bank/internal/infra/config"
)

func providers() []any {
	globalDeps := []any{
		config.New,
		http.NewControllerManager,
	}

	globalDeps = append(globalDeps, http.ControllersList...)
	return globalDeps
}
