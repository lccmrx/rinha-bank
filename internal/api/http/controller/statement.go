package controller

import "github.com/labstack/echo/v4"

type Statement struct {
	// service service.Statment
}

func NewStatement() *Statement {
	return &Statement{}
}

func (s *Statement) Get(ctx echo.Context) (err error) {
	// s.service.Get()
	return
}
