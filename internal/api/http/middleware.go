package http

import "github.com/labstack/echo/v4/middleware"

func (s *Server) setupMiddlewares() {
	s.e.Pre(middleware.AddTrailingSlash())

	s.e.Use(middleware.Gzip())
	s.e.Use(middleware.Recover())
}
