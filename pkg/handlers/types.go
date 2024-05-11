package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Handler - Handler interface with all crud methods
type Handler interface {
	CreateEmployee(ctx echo.Context) error
	GetEmployeeById(ctx echo.Context) error
	UpdateEmployee(ctx echo.Context) error
	DeleteEmployee(ctx echo.Context) error
}

// RequestHandler - request handler to have all basic methods for the apis
type RequestHandler struct {
	logger *zap.SugaredLogger
}
