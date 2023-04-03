package query

import (
	"backend/internal/infrastructure/db/private"
	"backend/models"
	"context"
)

type Repository struct {
	connection *private.Connection
}

func NewProvider(connection *private.Connection) *Repository {
	return &Repository{connection: connection}
}

func (r *Repository) Create(ctx context.Context, query *models.Query) error {
	q := "INSERT INTO queries (query_raw, query_hash, result_raw, result_hash, field_descriptions, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	rows, err := r.connection.Pool.Query(
		ctx,
		q,
		query.QueryRaw,
		query.QueryHash,
		query.ResultRaw,
		query.ResultHash,
		query.FieldDescriptions,
		query.CreatedAt,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&query.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
