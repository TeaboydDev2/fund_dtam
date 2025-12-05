package entities

import "time"

type Ebook struct {
	ID         string      `json:"id" bson:"_id"`
	Thumbnail  *FileObject `json:"thumbnail" bson:"thumbnail"`
	EBookFile  *FileObject `json:"ebook_file" bson:"ebook_file"`
	Title      string      `json:"title" bson:"title"`
	Number     int         `json:"number" bson:"number"`
	Status     bool        `json:"status" bson:"status"`
	ViewStatic int64       `json:"view_static" bson:"view_static"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at" bson:"updated_at"`
}
