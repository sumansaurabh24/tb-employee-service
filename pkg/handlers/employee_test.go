package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sumansaurabh24/tb-employee-service/pkg/errors"
	"github.com/sumansaurabh24/tb-employee-service/pkg/models"
	"github.com/sumansaurabh24/tb-employee-service/pkg/repositories"
	"github.com/sumansaurabh24/tb-employee-service/pkg/services"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	prodLog, _   = zap.NewProduction()
	logger       = prodLog.Sugar()
	db           = repositories.NewDatabase[models.Employee](logger)
	service      = services.NewServiceImpl[models.Employee](logger, &db)
	eh           = NewEmployeeHandler(logger, service)
	employee     = models.Employee{}
	employeeJson = `{"name": "testN", "position": "testP", "salary": 12}`
)

func TestEmployeeHandler_CreateEmployee(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", strings.NewReader(employeeJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := eh.CreateEmployee(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	_ = json.Unmarshal(rec.Body.Bytes(), &employee)
	logger.With("employee", employee).Info("response")
	assert.Equal(t, "testN", employee.Name)
}

func TestEmployeeHandler_GetEmployeeById(t *testing.T) {
	db.Insert(models.Employee{
		Name:     "testN",
		Position: "testP",
		Salary:   0,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(employeeJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/employees/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := eh.GetEmployeeById(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	_ = json.Unmarshal(rec.Body.Bytes(), &employee)
	logger.With("employee", employee).Info("response from TestEmployeeHandler_GetEmployeeById")
	assert.Equal(t, "testN", employee.Name)
	assert.Equal(t, 1, employee.ID)
}

func TestEmployeeHandler_UpdateEmployee(t *testing.T) {
	db.Insert(models.Employee{
		Name:     "testN1",
		Position: "testP",
		Salary:   0,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(employeeJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/employees/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := eh.UpdateEmployee(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	_ = json.Unmarshal(rec.Body.Bytes(), &employee)
	logger.With("employee", employee).Info("response from TestEmployeeHandler_UpdateEmployee")
	assert.Equal(t, "testN", employee.Name)
	assert.Equal(t, 1, employee.ID)
}

func TestEmployeeHandler_DeleteEmployee(t *testing.T) {
	db.Insert(models.Employee{
		Name:     "testN1",
		Position: "testP",
		Salary:   0,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(employeeJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/employees/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := eh.DeleteEmployee(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	_, err = db.Find(1)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), errors.ErrEntityNotFound.Error())
}
