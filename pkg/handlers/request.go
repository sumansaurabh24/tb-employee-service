package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sumansaurabh24/tb-employee-service/pkg/errors"
)

// Body - handles the request body and bind it to the struct
func (r *RequestHandler) Body(ctx echo.Context, T interface{}) error {
	r.logger.Info("binding request body to object")
	err := ctx.Bind(T)
	if err != nil {
		return &errors.ErrFailedParsingBody
	}
	return nil
}
