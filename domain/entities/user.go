package entities

import "time"

type User struct {
	ID             string
	FirstName      string
	LastName       string
	ProfilePicture *FileObject
	Illustration   []*FileObject
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
