package entities

import "time"

type Announcement struct {
	ID               string     `json:"id" bson:"_id"`
	Title            string     `json:"title" bson:"title"`
	FileAnnouncement FileObject `json:"file_announcement" bson:"file_announcement"`
	ViewStatic       int64      `json:"view_static" bson:"view_static"`
	Status           bool       `json:"status" bson:"status"`
	AnnouncementType string     `json:"announcement_type" bson:"announcement_type"`
	CreatedAt        time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" bson:"updated_at"`
}
