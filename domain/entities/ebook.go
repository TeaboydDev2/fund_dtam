package entities

import "time"

type Ebook struct {
	ID         string
	Thumbnail  FileObject
	EBookFile  FileObject
	Title      string
	Url        string
	Number     int
	Status     bool
	ViewStatic int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
