package errors

import (
	"fmt"
	"net/http"
)

type HttpError interface {
	Code() int
	Message() string
	Error() string
}

type httpError struct {
	code    int
	message string
}

func (e *httpError) Code() int {
	return e.code
}

func (e *httpError) Message() string {
	return e.message
}

func (e *httpError) Error() string {
	return fmt.Sprintf("code = %d; message = %s", e.code, e.message)
}

func NewHttpError(code int, message string) HttpError {
	return &httpError{code: code, message: message}
}

var (
	NotFoundHttpError     = NewHttpError(http.StatusNotFound, "Not found")
	UnauthorizedHttpError = NewHttpError(http.StatusUnauthorized, "Unauthorized")
)
