package task

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

func (r *Repository) Create(ctx context.Context, task *models.Task) error {
	q := "INSERT INTO tasks (name, description, query_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	rows, err := r.pool.Query(
		ctx,
		q,
		task.Name,
		task.Description,
		task.QueryID,
		task.CreatedAt,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&task.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
