package errors

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	Code() int
	Message() string
	Error() string
	UserReadableMessage() string
}

type httpError struct {
	code                int
	message             string
	userReadableMessage string
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

func (e *httpError) UserReadableMessage() string {
	return e.userReadableMessage
}

func NewHTTPError(code int, message string) HTTPError {
	return &httpError{code: code, message: message}
}

func NewHTTPErrorWithUserReadableMessage(code int, message string, urMessage string) HTTPError {
	return &httpError{code: code, message: message, userReadableMessage: urMessage}
}

var (
	NotFoundHTTPError     = NewHTTPError(http.StatusNotFound, "Not found")
	UnauthorizedHTTPError = NewHTTPError(http.StatusUnauthorized, "Unauthorized")
)
