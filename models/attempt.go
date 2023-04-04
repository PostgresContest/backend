package models

import "time"

type Attempt struct {
	ID        int64
	UserID    int64
	QueryID   int64
	TaskID    int64
	Accepted  bool
	CreatedAt time.Time
}
