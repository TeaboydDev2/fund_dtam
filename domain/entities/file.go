package entities

import "mime/multipart"

// file value object //
type FileObject struct {
	Alt         string         `json:"alt" bson:"alt"`
	Ext         string         `json:"ext" bson:"ext"`
	Path        string         `json:"path" bson:"path"`
	ContentType string         `json:"-" bson:"-"`
	Size        int64          `json:"-" bson:"-"`
	File        multipart.File `json:"-" bson:"-"`
}
