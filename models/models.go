package models

import "time"

type User struct {
	ID           int64
	FirstName    string
	LastName     string
	PasswordHash string

	RegisteredAt time.Time
}
