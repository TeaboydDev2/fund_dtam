package model

type FileObjectDB struct {
	Name string `bson:"alt"`
	Ext  string `bson:"ext"`
	Path string `bson:"path"`
}
