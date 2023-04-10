package hydrators

import (
	"encoding/json"

	"backend/internal/infrastructure/executor"
	"backend/models"
	oapi "github.com/PostgresContest/openapi/go/gen/v1"
)

type QueryOption func(query *oapi.Query)

func QueryHideSolution() QueryOption {
	return func(query *oapi.Query) {
		query.QueryRow = ""
		query.QueryHash = ""
	}
}

func HydrateQuery(q *models.Query, options ...QueryOption) *oapi.Query {
	var fds []executor.FieldDescription
	_ = json.Unmarshal([]byte(q.FieldDescriptions), &fds)

	query := &oapi.Query{
		ID:               q.ID,
		QueryRow:         q.QueryRaw,
		QueryHash:        q.QueryHash,
		ResultRaw:        q.ResultRaw,
		ResultHash:       q.ResultHash,
		FieldDescription: HydrateFieldDescriptions(fds),
	}

	for _, option := range options {
		option(query)
	}
	return query
}
