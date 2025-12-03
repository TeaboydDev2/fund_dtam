package model

import "dtam-fund-cms-backend/domain/entities"

type FileObjectDB struct {
	Alt  string `bson:"alt"`
	Ext  string `bson:"ext"`
	Path string `bson:"path"`
}

func FileToModel(file *entities.FileObject) *FileObjectDB {
	return &FileObjectDB{
		Alt:  file.Alt,
		Ext:  file.Ext,
		Path: file.Path,
	}
}

func FileToModelList(files []*entities.FileObject) []*FileObjectDB {

	file := make([]*FileObjectDB, len(files))

	for i, v := range files {
		file[i] = &FileObjectDB{
			Alt:  v.Alt,
			Ext:  v.Ext,
			Path: v.Path,
		}
	}

	return file
}

func FileToEntity(file *FileObjectDB) *entities.FileObject {
	return &entities.FileObject{
		Alt:  file.Alt,
		Ext:  file.Ext,
		Path: file.Path,
	}
}

func FileListToEntity(files []*FileObjectDB) []*entities.FileObject {

	file := make([]*entities.FileObject, len(files))

	for i, v := range files {
		file[i] = &entities.FileObject{
			Alt:  v.Alt,
			Ext:  v.Ext,
			Path: v.Path,
		}
	}

	return file
}
