package models

import (
	"backend/internal/infrastructure/executor"
	"encoding/json"
	"time"
)

type Query struct {
	ID         int64
	QueryRaw   string
	QueryHash  string
	ResultRaw  string
	ResultHash string

	CreatedAt time.Time
}

func (q *Query) FromExecutorResult(res *executor.Result) *Query {

	q.QueryRaw = res.Query
	q.QueryHash = res.QueryHash
	q.ResultHash = res.ResultHash

	rowsJson, _ := json.Marshal(res.Rows)
	q.ResultRaw = string(rowsJson)

	return q
}
