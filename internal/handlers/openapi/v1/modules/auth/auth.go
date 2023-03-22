package auth

import (
	"backend/internal/errors"
	"backend/internal/infrastructure/auth"
	"backend/internal/infrastructure/repositories/user"
	"backend/internal/logger"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type ModuleAuth struct {
	log            *logrus.Entry
	userRepository *user.Repository
	jwt            *auth.Jwt
}

func NewProvider(log *logger.Logger, jwt *auth.Jwt, userRepository *user.Repository) *ModuleAuth {
	l := log.WithField("module", "openapi.auth")

	return &ModuleAuth{
		log:            l,
		jwt:            jwt,
		userRepository: userRepository,
	}
}

func (m *ModuleAuth) AuthLoginPost(_ context.Context, req *oapi.LoginBody) (*oapi.Jwt, error) {
	u, err := m.userRepository.GetByLogin(req.Login)
	if err != nil {
		return nil, err
	}

	if !u.ComparePassword(req.Password) {
		return nil, errors.UnauthorizedHttpError
	}

	token, err := m.jwt.Generate(u.ID)

	return &oapi.Jwt{Token: token}, err
}
