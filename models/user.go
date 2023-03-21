package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           int64
	Login        string
	FirstName    string
	LastName     string
	PasswordHash string

	RegisteredAt time.Time
	UpdatedAt    time.Time
}

func (u *User) SetPasswordHash(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.PasswordHash = string(passHash)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}
