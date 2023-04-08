package models

import (
	"encoding/json"
	"time"

	"backend/internal/infrastructure/executor"
)

type Query struct {
	ID                int64
	QueryRaw          string
	QueryHash         string
	ResultRaw         string
	ResultHash        string
	FieldDescriptions string

	CreatedAt time.Time
}

func (q *Query) FromExecutorResult(res *executor.Result) *Query {
	q.QueryRaw = res.Query
	q.QueryHash = res.QueryHash
	q.ResultHash = res.ResultHash

	//nolint:all
	rowsJSON, _ := json.Marshal(res.Rows)
	q.ResultRaw = string(rowsJSON)

	//nolint:all
	fdJSON, _ := json.Marshal(res.FieldDescription)
	q.FieldDescriptions = string(fdJSON)

	return q
}

func (q *Query) SetCreatedNow() *Query {
	q.CreatedAt = time.Now()

	return q
}
