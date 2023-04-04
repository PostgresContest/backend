package task

import (
	"backend/internal/context"
	"backend/internal/errors"
	"backend/internal/infrastructure/executor"
	"backend/internal/infrastructure/handlers/openapi/v1/hydrators"
	attemptRepository "backend/internal/infrastructure/repositories/attempt"
	queryRepository "backend/internal/infrastructure/repositories/query"
	taskRepository "backend/internal/infrastructure/repositories/task"
	"backend/internal/logger"
	"backend/models"
	builtinContext "context"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"github.com/sirupsen/logrus"
	"time"
)

type ModuleTask struct {
	log               *logrus.Entry
	executor          *executor.Executor
	queryRepository   *queryRepository.Repository
	taskRepository    *taskRepository.Repository
	attemptRepository *attemptRepository.Repository
}

func NewProvider(
	log *logger.Logger,
	executor *executor.Executor,
	queryRepository *queryRepository.Repository,
	taskRepository *taskRepository.Repository,
	attemptRepository *attemptRepository.Repository,
) *ModuleTask {
	l := log.WithField("module", "openapi.task")

	return &ModuleTask{
		log:               l,
		executor:          executor,
		queryRepository:   queryRepository,
		taskRepository:    taskRepository,
		attemptRepository: attemptRepository,
	}
}

func (m *ModuleTask) TaskPost(ctx builtinContext.Context, req *oapi.TaskPostReq) (oapi.TaskPostRes, error) {
	result, err := m.executor.Execute(ctx, req.QueryRaw)
	if err != nil {
		return nil, err
	}

	query := new(models.Query).FromExecutorResult(result).SetCreatedNow()

	if err := m.queryRepository.Create(ctx, query); err != nil {
		return nil, err
	}

	task := &models.Task{
		Name:        req.Name,
		Description: req.Description,
		QueryID:     query.ID,
		CreatedAt:   time.Now(),
	}

	if err = m.taskRepository.Create(ctx, task); err != nil {
		return nil, err
	}

	taskRes := hydrators.HydrateTask(task, hydrators.TaskWithQuery(query))

	return taskRes, nil
}

func (m *ModuleTask) TasksGet(ctx builtinContext.Context) (oapi.TasksGetRes, error) {
	tasks, err := m.taskRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	tasksIDs := make([]int64, len(tasks))
	for i, task := range tasks {
		tasksIDs[i] = task.ID
	}

	lastAttemptsMap, err := m.attemptRepository.GetLastAttemptsToTasks(ctx, context.UserID(ctx), tasksIDs)
	if err != nil {
		return nil, err
	}
	result := make(oapi.TasksGetOKApplicationJSON, len(tasks))
	for i, task := range tasks {
		var options []hydrators.TaskOption
		if attempt, ok := lastAttemptsMap[task.ID]; ok {
			options = append(options, hydrators.TaskWithLastAttempt(&attempt))
		}
		result[i] = *hydrators.HydrateTask(&task, options...)
	}

	return &result, nil
}

func checkIsResponseCorrect(referenceQuery *models.Query, checkingQueryID *models.Query) bool {
	if referenceQuery.QueryHash == checkingQueryID.QueryHash {
		return true
	}
	if referenceQuery.ResultHash == checkingQueryID.ResultHash {
		return true
	}
	return false
}

func (m *ModuleTask) TaskTaskIDAttemptPost(ctx builtinContext.Context, req *oapi.TaskTaskIDAttemptPostReq, params oapi.TaskTaskIDAttemptPostParams) (oapi.TaskTaskIDAttemptPostRes, error) {
	q := req.QueryRaw
	result, err := m.executor.Execute(ctx, q)
	if err != nil {
		return nil, err
	}

	query := new(models.Query).FromExecutorResult(result).SetCreatedNow()

	if err = m.queryRepository.Create(ctx, query); err != nil {
		return nil, err
	}

	task, err := m.taskRepository.GetByID(ctx, params.TaskID)
	if err != nil {
		return nil, err
	}

	referenceQuery, err := m.queryRepository.GetByID(ctx, task.QueryID)
	if err != nil {
		return nil, err
	}

	userID := context.UserID(ctx)

	attempt := &models.Attempt{
		UserID:    userID,
		QueryID:   query.ID,
		TaskID:    task.ID,
		Accepted:  checkIsResponseCorrect(referenceQuery, query),
		CreatedAt: time.Now(),
	}

	err = m.attemptRepository.Create(ctx, attempt)
	if err != nil {
		return nil, err
	}

	attemptResult := hydrators.HydrateAttempt(attempt, hydrators.AttemptWithQuery(query))

	return attemptResult, nil
}

func (m *ModuleTask) TaskTaskIDAttemptsGet(ctx builtinContext.Context, params oapi.TaskTaskIDAttemptsGetParams) ([]oapi.Attempt, error) {
	attempts, err := m.attemptRepository.GetByTaskID(ctx, params.TaskID)
	if err != nil {
		return nil, err
	}

	result := make([]oapi.Attempt, len(attempts))
	for i, attempt := range attempts {
		result[i] = *hydrators.HydrateAttempt(&attempt)
	}

	return result, err
}

func (m *ModuleTask) AttemptAttemptIDGet(ctx builtinContext.Context, params oapi.AttemptAttemptIDGetParams) (oapi.AttemptAttemptIDGetRes, error) {
	attempt, err := m.attemptRepository.GetByID(ctx, params.AttemptID)
	if err != nil {
		return nil, err
	}

	if attempt.UserID != context.UserID(ctx) && context.UserRole(ctx) != models.UserRoleAdmin {
		return nil, errors.UnauthorizedHTTPError
	}

	query, err := m.queryRepository.GetByID(ctx, attempt.QueryID)
	if err != nil {
		return nil, err
	}

	return hydrators.HydrateAttempt(attempt, hydrators.AttemptWithQuery(query)), err
}
