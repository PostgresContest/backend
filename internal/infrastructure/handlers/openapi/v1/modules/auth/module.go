package auth

import (
	"context"
	stdErrs "errors"

	"backend/internal/errors"
	"backend/internal/infrastructure/auth"
	"backend/internal/infrastructure/repositories/user"
	"backend/internal/logger"
	oapi "github.com/PostgresContest/openapi/go/gen/v1"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type ModuleAuth struct {
	log            *logrus.Entry
	userRepository *user.Repository
	atProvider     *auth.AccessTokenProvider
}

func NewProvider(
	log *logger.Logger,
	atProvider *auth.AccessTokenProvider,
	userRepository *user.Repository,
) *ModuleAuth {
	l := log.WithField("module", "openapi.auth")

	return &ModuleAuth{
		log:            l,
		atProvider:     atProvider,
		userRepository: userRepository,
	}
}

func (m *ModuleAuth) AuthLoginPost(ctx context.Context, req *oapi.AuthLoginPostReq) (*oapi.Jwt, error) {
	usr, err := m.userRepository.GetByLogin(ctx, req.Login)
	if err != nil {
		if stdErrs.Is(err, pgx.ErrNoRows) {
			return nil, errors.UnauthorizedHTTPError
		}

		return nil, err
	}

	if !usr.ComparePassword(req.Password) {
		return nil, errors.UnauthorizedHTTPError
	}

	token, claims, err := m.atProvider.Generate(usr.ID, usr.Role)

	return &oapi.Jwt{
		AccessToken: token,
		Exp:         claims.GetExpiration(),
		Role:        usr.Role,
	}, err
}

func (m *ModuleAuth) AuthVerifyGet(_ context.Context) (*oapi.OkResponse, error) {
	return &oapi.OkResponse{Status: "ok"}, nil
}
