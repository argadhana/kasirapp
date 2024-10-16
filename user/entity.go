package user

import "time"

type User struct {
	ID           int
	Email        string
	PasswordHash string
	Name         string
	Phone        int
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
