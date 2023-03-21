package auth

import (
	"backend/internal/errors"
	"backend/internal/infrastructure/auth"
	"backend/internal/logger"
	"backend/internal/repositories/user"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log            *logrus.Entry
	userRepository *user.Repository
	jwt            *auth.Jwt
}

func NewProvider(
	log *logger.Logger,
	jwt *auth.Jwt,
	userRepository *user.Repository,
) *Handler {
	l := log.WithField("module", "openapi.auth")

	return &Handler{
		log:            l,
		jwt:            jwt,
		userRepository: userRepository,
	}
}

func (h *Handler) AuthLoginPost(_ context.Context, req *oapi.LoginBody) (*oapi.Jwt, error) {
	u, err := h.userRepository.GetByLogin(req.Login)
	if err != nil {
		return nil, err
	}

	if !u.ComparePassword(req.Password) {
		return nil, errors.UnauthorizedHttpError
	}

	token, err := h.jwt.Generate(u.ID)

	return &oapi.Jwt{Token: token}, err
}
