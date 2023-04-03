package hydrators

import (
	"backend/internal/infrastructure/executor"
	oapi "github.com/PostgresContest/openapi/gen/v1"
)

func HydrateFieldDescription(description executor.FieldDescription) *oapi.FieldDescription {
	return &oapi.FieldDescription{
		Name:     oapi.NewOptString(description.Name),
		Datatype: oapi.NewOptString(description.Datatype),
	}
}

func HydrateFieldDescriptions(fds []executor.FieldDescription) []oapi.FieldDescription {
	res := make([]oapi.FieldDescription, len(fds))
	for i, fd := range fds {
		res[i] = *HydrateFieldDescription(fd)
	}
	return res
}
