package models

import (
	"backend/internal/infrastructure/executor"
	"encoding/json"
	"time"
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

	rowsJson, _ := json.Marshal(res.Rows)
	q.ResultRaw = string(rowsJson)

	fdJson, _ := json.Marshal(res.FieldDescription)
	q.FieldDescriptions = string(fdJson)

	return q
}

func (q *Query) SetCreatedNow() *Query {
	q.CreatedAt = time.Now()
	return q
}
