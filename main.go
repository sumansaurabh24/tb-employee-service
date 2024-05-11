package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sumansaurabh24/tb-employee-service/pkg/constants"
	"github.com/sumansaurabh24/tb-employee-service/pkg/handlers"
	"github.com/sumansaurabh24/tb-employee-service/pkg/models"
	"github.com/sumansaurabh24/tb-employee-service/pkg/repositories"
	"github.com/sumansaurabh24/tb-employee-service/pkg/services"
	"go.uber.org/zap"
	"time"
)

func main() {
	// we can go with comprehensive configuration but as of now
	// going with the simple configuration of zap logger
	prodLog, _ := zap.NewProduction()
	logger := prodLog.Sugar()

	e := echo.New()

	// add logger middleware
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.With(
				zap.String("uri", values.URI),
				zap.Any("status", values.Status),
				zap.String("host", values.Host),
				zap.String("method", values.Method),
				zap.Any("duration", time.Since(values.StartTime).Milliseconds()),
				zap.Any("headers", values.Headers),
				zap.String("ip", values.RemoteIP),
				zap.String("trace_id", values.RequestID),
			).Info("request received")
			return nil
		},
	}))

	// register global error handler
	errH := handlers.ErrorHandler{Logger: logger}
	e.HTTPErrorHandler = errH.Global

	// create new Employee handler and attach it to the echo handler
	db := repositories.NewDatabase[models.Employee](logger)
	service := services.NewServiceImpl[models.Employee](logger, &db)
	eh := handlers.NewEmployeeHandler(logger, service)

	// employee api group and its handler
	employeeGroup := e.Group(constants.ApiV1EmployeeGroup)
	employeeGroup.POST(constants.PostEmployeesRoute, eh.CreateEmployee)
	employeeGroup.GET(constants.GetEmployeesRoute, eh.GetEmployeeById)
	employeeGroup.PUT(constants.PutEmployeesRoute, eh.UpdateEmployee)
	employeeGroup.DELETE(constants.DeleteEmployeesRoute, eh.DeleteEmployee)

	// start the server
	e.Logger.Fatal(e.Start(":8989"))
}
