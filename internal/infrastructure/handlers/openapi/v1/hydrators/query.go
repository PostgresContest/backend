package hydrators

import (
	"backend/models"
	oapi "github.com/PostgresContest/openapi/gen/v1"
)

func HydrateQuery(models *models.Query) *oapi.Query {
	return &oapi.Query{
		ID:           models.ID,
		QueryRow:     models.QueryRaw,
		QueryHash:    models.QueryHash,
		ResponseRaw:  models.ResultRaw,
		ResponseHash: models.ResultHash,
	}
}
