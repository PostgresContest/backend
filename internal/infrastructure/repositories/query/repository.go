package query

import (
	"backend/internal/infrastructure/db/private"
	"backend/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewProvider(connection *private.Connection) *Repository {
	return &Repository{pool: connection.Pool}
}

func (r *Repository) Create(ctx context.Context, query *models.Query) error {
	q := "INSERT INTO queries (query_raw, query_hash, result_raw, result_hash, field_descriptions, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	rows, err := r.pool.Query(
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

func (r *Repository) GetByID(ctx context.Context, ID int64) (*models.Query, error) {
	q := "SELECT id, query_raw, query_hash, result_raw, result_hash, created_at, field_descriptions FROM queries WHERE id = $1"
	row := r.pool.QueryRow(ctx, q, ID)
	var query models.Query
	err := row.Scan(
		&query.ID,
		&query.QueryRaw,
		&query.QueryHash,
		&query.ResultRaw,
		&query.ResultHash,
		&query.CreatedAt,
		&query.FieldDescriptions,
	)
	if err != nil {
		return nil, err
	}
	return &query, nil
}
