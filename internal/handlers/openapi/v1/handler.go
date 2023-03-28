package v1

import (
	"context"

	"backend/internal/config"
	"backend/internal/handlers/openapi/v1/modules/auth"
	"backend/internal/handlers/openapi/v1/modules/user"
	"backend/internal/logger"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	devMode bool
	log     *logrus.Entry

	*auth.ModuleAuth
	*user.ModuleUser
}

func (h *Handler) TaskPost(_ context.Context, _ oapi.OptTaskPostReq) (*oapi.Task, error) {
	panic("unimpl")
}

func NewProvider(
	log *logger.Logger,
	cfg *config.Config,

	authModule *auth.ModuleAuth,
	userModule *user.ModuleUser,
) *Handler {
	l := log.WithField("module", "openapi")

	return &Handler{
		ModuleAuth: authModule,
		ModuleUser: userModule,
		log:        l,
		devMode:    cfg.IsDevMode(),
	}
}
