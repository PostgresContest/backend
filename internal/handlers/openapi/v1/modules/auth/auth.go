package auth

import (
	"context"

	"backend/internal/errors"
	"backend/internal/infrastructure/auth"
	"backend/internal/infrastructure/repositories/user"
	"backend/internal/logger"

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

func (m *ModuleAuth) AuthLoginPost(ctx context.Context, req *oapi.LoginBody) (*oapi.Jwt, error) {
	usr, err := m.userRepository.GetByLogin(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if !usr.ComparePassword(req.Password) {
		return nil, errors.UnauthorizedHTTPError
	}

	token, err := m.jwt.Generate(usr.ID)

	return &oapi.Jwt{Token: token}, err
}
