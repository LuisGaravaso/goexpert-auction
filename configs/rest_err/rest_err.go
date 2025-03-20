package rest_err

import (
	"net/http"

	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *RestErr) Error() string {
	return e.Message
}

func NewBadRequestError(message string, causes ...Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
	return result
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}

func ConvertError(internalError *internal_errors.InternalError) *RestErr {
	switch internalError.Err {
	case "bad_request":
		return NewBadRequestError(internalError.Message)
	case "internal_server_error":
		return NewInternalServerError(internalError.Message, nil)
	case "not_found":
		return NewNotFoundError(internalError.Message)
	default:
		return NewInternalServerError(internalError.Message, nil)
	}
}
