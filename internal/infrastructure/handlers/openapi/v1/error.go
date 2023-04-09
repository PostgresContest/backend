package v1

import (
	"context"
	errorsBuiltin "errors"
	"net/http"

	"backend/internal/errors"
	oapi "github.com/PostgresContest/openapi/go/gen/v1"
	"github.com/jackc/pgx/v5"
)

func processError(err error) error {
	if errorsBuiltin.Is(err, pgx.ErrNoRows) {
		return errors.NotFoundHTTPError
	}

	return err
}

var internalErr = &oapi.ErrStatusCode{
	StatusCode: http.StatusInternalServerError,
	Response: oapi.Err{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	},
}

func (h *Handler) NewError(_ context.Context, err error) *oapi.ErrStatusCode {
	err = processError(err)
	if err == nil {
		return internalErr
	}

	var httpErr errors.HTTPError
	if errorsBuiltin.As(err, &httpErr) && httpErr != nil {
		return &oapi.ErrStatusCode{
			StatusCode: httpErr.Code(),
			Response: oapi.Err{
				Code:    httpErr.Code(),
				Message: httpErr.Message(),
			},
		}
	}

	message := "Internal server error"

	if h.devMode {
		message = err.Error()
	}

	return &oapi.ErrStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: oapi.Err{
			Code:    http.StatusInternalServerError,
			Message: message,
		},
	}
}
