package error

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

// ResponseError represent the responseError error struct
type ResponseError struct {
	Message string `json:"message"`
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Resource Not Found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Confict")

	ErrBadRequest = errors.New("Request body not valid")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Invalid Parameter")

	TotalStockLessThanQty   = errors.New("Available stock is less than quantiry")
	CheckSumPaymentNotMatch = errors.New("check sum amount failed")
	OrderHasBeenExist       = errors.New("Order has been process")
	TokenExpired            = errors.New("token expired")
)

// error handler
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case TotalStockLessThanQty:
		return http.StatusBadRequest
	case CheckSumPaymentNotMatch:
		return http.StatusBadRequest
	case OrderHasBeenExist:
		return http.StatusBadRequest
	case TokenExpired:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
