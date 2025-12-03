package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EBookDB struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Thumbnail  *FileObjectDB      `bson:"thumbnail"`
	EBookFile  *FileObjectDB      `bson:"ebook_file"`
	Title      string             `bson:"title"`
	Number     int                `bson:"number"`
	Status     bool               `bson:"status"`
	ViewStatic int64              `bson:"view_static"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

func EbookToModel(ebook *entities.Ebook) *EBookDB {
	return &EBookDB{
		Thumbnail:  FileToModel(ebook.Thumbnail),
		EBookFile:  FileToModel(ebook.EBookFile),
		Title:      ebook.Title,
		Number:     0,
		Status:     true,
		ViewStatic: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func EBookToEntity(ebook *EBookDB) *entities.Ebook {
	return &entities.Ebook{
		ID:         ebook.ID.Hex(),
		Thumbnail:  FileToEntity(ebook.Thumbnail),
		EBookFile:  FileToEntity(ebook.EBookFile),
		Title:      ebook.Title,
		Number:     ebook.Number,
		Status:     ebook.Status,
		ViewStatic: ebook.ViewStatic,
		CreatedAt:  ebook.CreatedAt,
		UpdatedAt:  ebook.UpdatedAt,
	}
}
