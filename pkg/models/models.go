package models

import "time"

// User describes the user object
type User struct {
	ID         int
	FirstName  string
	LastName   string
	UserActive int
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
