package http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/lccmrx/rinha-bank/internal/api"
	"github.com/lccmrx/rinha-bank/internal/infra/config"
)

type Server struct {
	e *echo.Echo
}

var _ api.Api = (*Server)(nil)

func New(cfg *config.Config, controllerManager *ControllerManager) *Server {
	server := &Server{
		e: echo.New(),
	}

	server.setupMiddlewares()
	server.setupRoutes(controllerManager)

	server.e.Server.Addr = fmt.Sprintf(":%d", cfg.Server.Port)

	server.e.HidePort = true
	server.e.HideBanner = true
	return server
}

func (s *Server) Run(ctx context.Context) error {
	return s.e.Start(s.e.Server.Addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
