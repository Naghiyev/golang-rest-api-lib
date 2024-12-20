package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (ae AppError) AsMessage() *AppError {
	return &AppError{
		Message: ae.Message,
	}

}
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewAmountValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}

}

func NewAccountTypeException(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}

}

func NewTransactionTypeError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}

}
