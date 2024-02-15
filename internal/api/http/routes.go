package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) setupRoutes(cm *ControllerManager) {
	mainRouter := s.e.Group("")

	mainRouter.POST("/clientes/:id/transacoes/", cm.transaction.Create)
	mainRouter.GET("/clientes/:id/extrato/", cm.statement.Get)

	mainRouter.GET("/ping/", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
