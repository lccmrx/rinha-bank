package app

import (
	"github.com/wyreyx/rinha-bank/internal/infra/config"
)

func providers() []any {
	return []any{
		config.New,
	}
}
