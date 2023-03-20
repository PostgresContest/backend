package v1

import (
	"backend/internal/errors"
	"context"
	errorsBuiltin "errors"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/jackc/pgx/v5"
)

func processError(err error) error {
	if errorsBuiltin.Is(err, pgx.ErrNoRows) {
		return errors.NotFoundHttpError
	}
	return err
}

var (
	internalStatus = &oapi.ErrorStatusCode{
		StatusCode: 500,
		Response: oapi.Error{
			Code:    500,
			Message: "Internal server error",
		},
	}
)

func (h *Handler) NewError(_ context.Context, err error) *oapi.ErrorStatusCode {
	err = processError(err)
	if err == nil {
		return internalStatus
	}

	if httpErr, ok := err.(errors.HttpError); ok {
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
		StatusCode: 500,
		Response: oapi.Error{
			Code:    500,
			Message: message,
		},
	}
}
