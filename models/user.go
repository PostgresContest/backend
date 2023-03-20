package models

import "time"

type User struct {
	ID           int64
	Login        string
	FirstName    string
	LastName     string
	PasswordHash string

	RegisteredAt time.Time
	UpdatedAt    time.Time
}
