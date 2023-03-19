package v1

import (
	"backend/internal/config"
	"backend/internal/handlers/openapi/v1/modules/auth"
	"backend/internal/logger"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	*auth.Handler
	devMode bool
	log     *logrus.Entry
}

func NewProvider(
	log *logger.Logger,
	cfg *config.Config,
	authModule *auth.Handler,
) *Handler {
	l := log.WithField("module", "openapi")

	return &Handler{
		Handler: authModule,
		log:     l,
		devMode: cfg.IsDevMode(),
	}
}

func (h *Handler) UserGet(ctx context.Context) (*oapi.User, error) {
	//TODO implement me
	panic("implement me")
}
