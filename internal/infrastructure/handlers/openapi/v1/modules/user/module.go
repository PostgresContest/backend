package user

import (
	"context"

	pctx "backend/internal/context"
	"backend/internal/errors"
	"backend/internal/infrastructure/repositories/user"
	"backend/internal/logger"

	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type ModuleUser struct {
	log            *logrus.Entry
	userRepository *user.Repository
}

func NewProvider(log *logger.Logger, userRepository *user.Repository) *ModuleUser {
	l := log.WithField("module", "openapi.user")

	return &ModuleUser{
		log:            l,
		userRepository: userRepository,
	}
}

func (m *ModuleUser) UserGet(ctx context.Context) (*oapi.User, error) {
	userID := pctx.UserID(ctx)
	if userID == 0 {
		return nil, errors.UnauthorizedHTTPError
	}

	usr, err := m.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &oapi.User{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Login:     usr.Login,
	}, nil
}
