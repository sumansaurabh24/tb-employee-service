package errors

import (
	"net/http"
)

const (
	// FailedParsingBody - Generic error codes
	FailedParsingBody = 101

	// FailedFetchingEmployee Resource related error codes
	FailedFetchingEmployee = 1001

	//EntityNotFound database related error
	EntityNotFound = 2001
)

// ErrorCodeMapping - error and its code mapping which will be used within application
var ErrorCodeMapping = map[int]Error{
	FailedParsingBody: {
		Code:       FailedParsingBody,
		HttpStatus: http.StatusUnprocessableEntity,
		Message:    http.StatusText(http.StatusUnprocessableEntity),
	},
	FailedFetchingEmployee: {
		Code:       FailedFetchingEmployee,
		HttpStatus: http.StatusInternalServerError,
		Message:    "failed fetching employee",
	},
	EntityNotFound: {
		Code:       EntityNotFound,
		HttpStatus: http.StatusNotFound,
		Message:    "entity not found",
	},
}

var (
	ErrFailedParsingBody      = New(FailedParsingBody)
	ErrFailedFetchingEmployee = New(FailedFetchingEmployee)
	ErrEntityNotFound         = New(EntityNotFound)
)
