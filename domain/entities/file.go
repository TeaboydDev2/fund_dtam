package entities

import "mime/multipart"

// file value object //
type FileObject struct {
	Alt         string
	Ext         string
	Path        string
	ContentType string
	Size        int64
	File        multipart.File
}
