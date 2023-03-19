package auth

import (
	"backend/internal/logger"
	"backend/internal/repositories/user"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log            *logrus.Entry
	userRepository *user.Repository
}

func NewProvider(
	log *logger.Logger,
	userRepository *user.Repository,
) *Handler {
	l := log.WithField("module", "openapi.auth")

	return &Handler{
		log:            l,
		userRepository: userRepository,
	}
}

func (h *Handler) AuthLoginPost(ctx context.Context, req *oapi.LoginBody) (*oapi.Jwt, error) {
	//TODO implement me
	panic("implement me")
}