package models

import (
	"time"
)

type Task struct {
	ID          int64
	Name        string
	Description string
	QueryID     int64

	CreatedAt time.Time
}
