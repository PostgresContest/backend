package hydrators

import (
	"backend/models"
	oapi "github.com/PostgresContest/openapi/gen/v1"
)

type TaskOption func(*oapi.Task)

func WithQuery(query *models.Query) TaskOption {
	return func(task *oapi.Task) {
		task.Query = oapi.NewOptQuery(*HydrateQuery(query))
	}
}

func HydrateTask(task *models.Task, options ...TaskOption) *oapi.Task {
	res := &oapi.Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Query:       oapi.OptQuery{},
	}
	for _, option := range options {
		option(res)
	}
	return res
}