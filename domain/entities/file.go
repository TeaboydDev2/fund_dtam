package entities

import "mime/multipart"

// file value object //
type FileObject struct {
	Name        string
	Ext         string
	Path        string
	ContentType string
	Size        int64
	File        multipart.File
}
