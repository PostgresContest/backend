package hydrators

import (
	"encoding/json"

	"backend/internal/infrastructure/executor"
	"backend/models"
	oapi "github.com/PostgresContest/openapi/gen/v1"
)

func HydrateQuery(q *models.Query) *oapi.Query {
	var fds []executor.FieldDescription
	_ = json.Unmarshal([]byte(q.FieldDescriptions), &fds)

	return &oapi.Query{
		ID:               q.ID,
		QueryRow:         q.QueryRaw,
		QueryHash:        q.QueryHash,
		ResultRaw:        q.ResultRaw,
		ResultHash:       q.ResultHash,
		FieldDescription: HydrateFieldDescriptions(fds),
	}
}
