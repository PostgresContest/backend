package v1

import (
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/handlers/openapi/v1/modules/auth"
	"backend/internal/infrastructure/handlers/openapi/v1/modules/task"
	"backend/internal/infrastructure/handlers/openapi/v1/modules/user"
	"backend/internal/logger"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	devMode bool
	log     *logrus.Entry

	*auth.ModuleAuth
	*user.ModuleUser
	*task.ModuleTask
}

func NewProvider(
	log *logger.Logger,
	cfg *config.Config,

	authModule *auth.ModuleAuth,
	userModule *user.ModuleUser,
	taskModule *task.ModuleTask,
) *Handler {
	l := log.WithField("module", "openapi")

	return &Handler{
		ModuleAuth: authModule,
		ModuleUser: userModule,
		ModuleTask: taskModule,
		log:        l,
		devMode:    cfg.IsDevMode(),
	}
}
