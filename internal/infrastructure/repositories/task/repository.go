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

func (r *Repository) GetAll(ctx context.Context) ([]models.Task, error) {
	q := "SELECT id, name, description, query_id, created_at FROM tasks ORDER BY created_at, id"
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Task

	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.QueryID,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (r *Repository) GetByID(ctx context.Context, ID int64) (*models.Task, error) {
	q := "SELECT id, name, description, query_id, created_at FROM tasks WHERE id = $1"
	row := r.pool.QueryRow(ctx, q, ID)
	var t models.Task
	err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Description,
		&t.QueryID,
		&t.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
