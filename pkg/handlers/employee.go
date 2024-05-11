package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sumansaurabh24/tb-employee-service/pkg/constants"
	"github.com/sumansaurabh24/tb-employee-service/pkg/errors"
	"github.com/sumansaurabh24/tb-employee-service/pkg/models"
	"github.com/sumansaurabh24/tb-employee-service/pkg/services"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// EmployeeHandler - handler for the employees api
type EmployeeHandler struct {
	RequestHandler
	EmployeeService services.Service[models.Employee]
}

// NewEmployeeHandler - EmployeeHandler constructor
func NewEmployeeHandler(logger *zap.SugaredLogger,
	employeeService services.ServiceImpl[models.Employee]) *EmployeeHandler {
	return &EmployeeHandler{
		RequestHandler:  RequestHandler{logger: logger},
		EmployeeService: &employeeService,
	}
}

// CreateEmployee - save/create new employee object into the database
func (e *EmployeeHandler) CreateEmployee(ctx echo.Context) error {
	e.logger.Info("invoked employee handler")
	var employee models.Employee
	err := e.Body(ctx, &employee)
	if err != nil {
		return err
	}
	e.logger.With(zap.Any("employee", employee)).Info("serialized employee value")
	created := e.EmployeeService.Create(employee)
	return ctx.JSON(http.StatusCreated, created)
}

// GetEmployeeById - read employee from the database
func (e *EmployeeHandler) GetEmployeeById(ctx echo.Context) error {
	strId := ctx.Param(constants.QueryParamId)
	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.ErrFailedFetchingEmployee.WithInfo("error", err)
	}
	existing, err := e.EmployeeService.FindById(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, existing)
}

// UpdateEmployee - update employee in database
func (e *EmployeeHandler) UpdateEmployee(ctx echo.Context) error {
	strId := ctx.Param(constants.QueryParamId)
	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.ErrFailedFetchingEmployee.WithInfo("error", err)
	}

	var employee models.Employee
	err = e.Body(ctx, &employee)
	if err != nil {
		return err
	}

	updated, err := e.EmployeeService.Update(id, employee)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updated)
}

// DeleteEmployee - delete employee in database
func (e *EmployeeHandler) DeleteEmployee(ctx echo.Context) error {
	strId := ctx.Param(constants.QueryParamId)
	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.ErrFailedFetchingEmployee.WithInfo("error", err)
	}
	err = e.EmployeeService.Delete(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusNoContent, "")
}
