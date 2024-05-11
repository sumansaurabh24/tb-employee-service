package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	errors2 "github.com/sumansaurabh24/tb-employee-service/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

// ErrorHandler - struct to handle the global error
type ErrorHandler struct {
	Logger *zap.SugaredLogger
}

// Global - global error handler for the service
func (eh *ErrorHandler) Global(err error, ctx echo.Context) {
	eh.Logger.Info("intercepted in global error handling")
	var appError *errors2.Error
	var echoErr *echo.HTTPError
	if errors.As(err, &appError) {
		eh.Logger.Error(appError)
		err = ctx.JSON(appError.HttpStatus, appError)
		return
	} else if errors.As(err, &echoErr) {
		eh.Logger.Error(echoErr)
		err = ctx.JSON(echoErr.Code, echoErr)
		return
	}
	eh.Logger.Warn("incoming error cannot be parsed into app error")
	err = ctx.String(http.StatusInternalServerError, err.Error())
}
