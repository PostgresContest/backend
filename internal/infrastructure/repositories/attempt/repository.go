package attempt

import (
	"context"

	"backend/internal/infrastructure/db/private"
	"backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewProvider(connection *private.Connection) *Repository {
	return &Repository{pool: connection.Pool}
}

func (r *Repository) Create(ctx context.Context, attempt *models.Attempt) error {
	q := `INSERT INTO attempts (user_id, query_id, task_id, accepted, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`

	rows, err := r.pool.Query(ctx, q, attempt.UserID, attempt.QueryID, attempt.TaskID, attempt.Accepted, attempt.CreatedAt)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&attempt.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) GetByTaskID(ctx context.Context, taskID int64) ([]models.Attempt, error) {
	q := `SELECT id, user_id, query_id, task_id, accepted, created_at
FROM attempts
WHERE task_id = $1
ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, q, taskID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []models.Attempt

	for rows.Next() {
		var a models.Attempt
		err = rows.Scan(
			&a.ID,
			&a.UserID,
			&a.QueryID,
			&a.TaskID,
			&a.Accepted,
			&a.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, a)
	}

	return result, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.Attempt, error) {
	q := `SELECT id, user_id, query_id, task_id, accepted, created_at
FROM attempts
WHERE id = $1`

	rows := r.pool.QueryRow(ctx, q, id)

	var a models.Attempt

	err := rows.Scan(
		&a.ID,
		&a.UserID,
		&a.QueryID,
		&a.TaskID,
		&a.Accepted,
		&a.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *Repository) GetLastAttemptsToTasks(
	ctx context.Context,
	userID int64,
	taskIDs []int64,
) (map[int64]models.Attempt, error) {
	q := `SELECT DISTINCT ON (task_id) task_id, id, user_id, query_id, accepted, created_at
FROM attempts
WHERE user_id = $1
  AND task_id = ANY ($2)
ORDER BY task_id, created_at DESC`

	rows, err := r.pool.Query(ctx, q, userID, taskIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[int64]models.Attempt)

	for rows.Next() {
		var a models.Attempt

		err = rows.Scan(
			&a.TaskID,
			&a.ID,
			&a.UserID,
			&a.QueryID,
			&a.Accepted,
			&a.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		m[a.TaskID] = a
	}

	return m, nil
}
