package models

import "time"

type User struct {
	ID         int
	FirstName  string
	LastName   string
	UserActive int
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
