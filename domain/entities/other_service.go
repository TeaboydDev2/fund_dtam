package entities

import "time"

type OtherSevice struct {
	ID         string
	Thumbnail  *FileObject
	Title      string
	Url        string
	Number     int
	Status     bool
	ViewStatic int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
