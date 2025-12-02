package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OtherSeviceDB struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Thumbnail  FileObjectDB       `bson:"thumbnail"`
	Title      string             `bson:"title"`
	Url        string             `bson:"url"`
	Number     int                `bson:"number"`
	Status     bool               `bson:"status"`
	ViewStatic int64              `bson:"view_static"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

func ToModelService(ots *entities.OtherSevice) *OtherSeviceDB {
	return &OtherSeviceDB{
		Thumbnail: FileObjectDB{
			Alt:  ots.Thumbnail.Alt,
			Ext:  ots.Thumbnail.Ext,
			Path: ots.Thumbnail.Path,
		},
		Title:      ots.Title,
		Url:        ots.Url,
		Number:     0,
		Status:     true,
		ViewStatic: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func ToEntityService(ots *OtherSeviceDB) *entities.OtherSevice {
	return &entities.OtherSevice{
		ID: ots.ID.Hex(),
		Thumbnail: entities.FileObject{
			Alt:  ots.Thumbnail.Alt,
			Ext:  ots.Thumbnail.Ext,
			Path: ots.Thumbnail.Path,
		},
		Title:      ots.Title,
		Url:        ots.Url,
		Number:     ots.Number,
		Status:     ots.Status,
		ViewStatic: ots.ViewStatic,
		CreatedAt:  ots.CreatedAt,
		UpdatedAt:  ots.UpdatedAt,
	}
}

func ToEntityServiceList(ots []*OtherSeviceDB) []*entities.OtherSevice {

	serviceList := make([]*entities.OtherSevice, len(ots))

	for i, v := range ots {
		serviceList[i] = &entities.OtherSevice{
			ID: v.ID.Hex(),
			Thumbnail: entities.FileObject{
				Alt:  v.Thumbnail.Alt,
				Ext:  v.Thumbnail.Ext,
				Path: v.Thumbnail.Path,
			},
			Title:      v.Title,
			Url:        v.Url,
			Number:     v.Number,
			Status:     v.Status,
			ViewStatic: v.ViewStatic,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

	}

	return serviceList
}
