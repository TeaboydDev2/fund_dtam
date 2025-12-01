package entities

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
