package v1

import (
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/handlers/openapi/v1/modules/auth"
	"backend/internal/infrastructure/handlers/openapi/v1/modules/task"
	"backend/internal/infrastructure/handlers/openapi/v1/modules/user"
	"backend/internal/logger"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	devMode bool
	log     *logrus.Entry

	*auth.ModuleAuth
	*user.ModuleUser
	*task.ModuleTask
}

func (h *Handler) AttemptAttemptIDGet(ctx context.Context, params oapi.AttemptAttemptIDGetParams) (oapi.AttemptAttemptIDGetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) TaskTaskIDAttemptPost(ctx context.Context, req oapi.OptTaskTaskIDAttemptPostReq, params oapi.TaskTaskIDAttemptPostParams) (oapi.TaskTaskIDAttemptPostRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) TaskTaskIDAttemptsGet(ctx context.Context, params oapi.TaskTaskIDAttemptsGetParams) ([]oapi.Attempt, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) TasksGet(ctx context.Context) (oapi.TasksGetRes, error) {
	//TODO implement me
	panic("implement me")
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
