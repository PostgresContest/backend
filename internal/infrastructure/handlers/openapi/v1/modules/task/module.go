package task

import (
	"backend/internal/infrastructure/executor"
	"backend/internal/infrastructure/handlers/openapi/v1/hydrators"
	queryRepository "backend/internal/infrastructure/repositories/query"
	taskRepository "backend/internal/infrastructure/repositories/task"
	"backend/internal/logger"
	"backend/models"
	"context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
	"time"
)

type ModuleTask struct {
	log             *logrus.Entry
	executor        *executor.Executor
	queryRepository *queryRepository.Repository
	taskRepository  *taskRepository.Repository
}

func NewProvider(
	log *logger.Logger,
	executor *executor.Executor,
	queryRepository *queryRepository.Repository,
	taskRepository *taskRepository.Repository,
) *ModuleTask {
	l := log.WithField("module", "openapi.task")

	return &ModuleTask{
		log:             l,
		executor:        executor,
		queryRepository: queryRepository,
		taskRepository:  taskRepository,
	}
}

func (m *ModuleTask) TaskPost(ctx context.Context, req oapi.OptTaskPostReq) (*oapi.Task, error) {
	result, err := m.executor.Execute(ctx, req.Value.QueryRaw)
	if err != nil {
		return nil, err
	}

	query := new(models.Query).FromExecutorResult(result)

	if err := m.queryRepository.Create(ctx, query); err != nil {
		return nil, err
	}

	task := &models.Task{
		Name:        req.Value.Name,
		Description: req.Value.Description,
		QueryID:     query.ID,
		CreatedAt:   time.Now(),
	}

	if err = m.taskRepository.Create(ctx, task); err != nil {
		return nil, err
	}

	taskRes := hydrators.HydrateTask(task, hydrators.WithQuery(query))

	return taskRes, nil
}
