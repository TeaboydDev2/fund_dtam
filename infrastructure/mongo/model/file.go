package model

type FileObjectDB struct {
	Alt  string `bson:"alt"`
	Ext  string `bson:"ext"`
	Path string `bson:"path"`
}
