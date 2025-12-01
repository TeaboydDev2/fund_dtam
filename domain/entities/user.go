package entities

import "time"

type User struct {
	ID             string
	FirstName      string
	LastName       string
	ProfilePicture FileObject
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
