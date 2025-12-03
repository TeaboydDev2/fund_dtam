package model

import "dtam-fund-cms-backend/domain/entities"

type FileObject struct {
	Alt  string `json:"alt"`
	Ext  string `json:"ext"`
	Path string `json:"url"`
}

func FileToResponse(files []*entities.FileObject) []*FileObject {

	file := make([]*FileObject, len(files))

	for i, v := range files {
		file[i] = &FileObject{
			Alt:  v.Alt,
			Ext:  v.Ext,
			Path: v.Path,
		}
	}
	return file
}
