package hydrators

import (
	"backend/models"

	oapi "github.com/PostgresContest/openapi/gen/v1"
)

type AttemptOption func(attempt *oapi.Attempt)

func AttemptWithQuery(query *models.Query) AttemptOption {
	return func(attempt *oapi.Attempt) {
		attempt.Query = oapi.NewOptQuery(*HydrateQuery(query))
	}
}

func HydrateAttempt(attempt *models.Attempt, options ...AttemptOption) *oapi.Attempt {
	result := &oapi.Attempt{
		ID:       attempt.ID,
		Accepted: attempt.Accepted,
	}
	for _, option := range options {
		option(result)
	}
	return result
}
