package auth

import (
	"backend/internal/logger"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Entry
}

func NewProvider(log *logger.Logger) *Handler {
	l := log.WithField("module", "openapi.auth")

	return &Handler{log: l}
}

func (h *Handler) AuthLoginPost(ctx context.Context, req *oapi.LoginBody) (*oapi.Jwt, error) {
	//TODO implement me
	panic("implement me")
}
