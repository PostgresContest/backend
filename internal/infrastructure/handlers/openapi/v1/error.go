package v1

import (
	"context"
	errorsBuiltin "errors"
	"net/http"

	"backend/internal/errors"

	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/jackc/pgx/v5"
)

func processError(err error) error {
	if errorsBuiltin.Is(err, pgx.ErrNoRows) {
		return errors.NotFoundHTTPError
	}

	return err
}

var internalStatus = &oapi.ErrorStatusCode{
	StatusCode: http.StatusInternalServerError,
	Response: oapi.Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	},
}

func (h *Handler) NewError(_ context.Context, err error) *oapi.ErrorStatusCode {
	err = processError(err)
	if err == nil {
		return internalStatus
	}

	var httpErr errors.HTTPError
	if errorsBuiltin.As(err, &httpErr) && httpErr != nil {
		return &oapi.ErrorStatusCode{
			StatusCode: httpErr.Code(),
			Response: oapi.Error{
				Code:    httpErr.Code(),
				Message: httpErr.Message(),
			},
		}
	}

	message := "Internal server error"

	if h.devMode {
		message = err.Error()
	}

	return &oapi.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: oapi.Error{
			Code:    http.StatusInternalServerError,
			Message: message,
		},
	}
}
