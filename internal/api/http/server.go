package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wyreyx/rinha-bank/internal/api"
	"github.com/wyreyx/rinha-bank/internal/infra/config"
	"go.uber.org/fx"
)

type Server struct {
	*http.Server
}

func NewServer(lc fx.Lifecycle, cfg *config.Config, ctx context.Context) api.Api {
	fmt.Println(1)
	srv := &Server{Server: &http.Server{Addr: ":8080"}}
	fmt.Println(ctx.Value("verbose").(bool))
	return srv
}

func (s *Server) Run(ctx context.Context) error {
	return s.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down...")
	return s.Server.Shutdown(ctx)
}
