package http

func (s *Server) setupRoutes(cm *ControllerManager) {
	mainRouter := s.e.Group("")

	mainRouter.POST("/clientes/:id/transacoes/", cm.transaction.Create)
	mainRouter.GET("/clientes/:id/extrato/", cm.statement.Get)
}
