package user

import "backend/internal/infrastructure/db/private"

type Repository struct {
	connection *private.Connection
}

func NewProvider(connection *private.Connection) *Repository {
	return &Repository{connection: connection}
}
