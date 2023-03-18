package v1

import (
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
)

func (h *Handler) NewError(_ context.Context, err error) *oapi.ErrorStatusCode {
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
