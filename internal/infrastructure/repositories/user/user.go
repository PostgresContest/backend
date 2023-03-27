package user

import (
	"context"

	"backend/internal/infrastructure/db/private"
	"backend/models"
)

type Repository struct {
	connection *private.Connection
}

func NewProvider(connection *private.Connection) *Repository {
	return &Repository{connection: connection}
}

func (r *Repository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	row := r.connection.Pool.QueryRow(
		ctx,
		"SELECT id, login, password_hash, first_name, last_name, role, registered_at, updated_at FROM users WHERE login = $1",
		login,
	)
	result := models.User{}

	err := row.Scan(
		&result.ID,
		&result.Login,
		&result.PasswordHash,
		&result.FirstName,
		&result.LastName,
		&result.Role,
		&result.RegisteredAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	row := r.connection.Pool.QueryRow(
		ctx,
		"SELECT id, login, password_hash, first_name, last_name, role, registered_at, updated_at FROM users WHERE id = $1",
		id,
	)
	result := models.User{}

	err := row.Scan(
		&result.ID,
		&result.Login,
		&result.PasswordHash,
		&result.FirstName,
		&result.LastName,
		&result.Role,
		&result.RegisteredAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
