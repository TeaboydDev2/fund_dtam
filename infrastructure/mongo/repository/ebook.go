package repository

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EBookRepository struct {
	collection *mongo.Collection
}

func NewEBookRepository(db *mongodb.MongoClient) ports.EBookRepository {
	return &EBookRepository{
		collection: db.Collection("ebook"),
	}
}

func (eb *EBookRepository) SaveEBook(ctx context.Context, ebook *entities.Ebook) (err error) {

	if _, err = eb.collection.InsertOne(ctx, ebook); err != nil {
		return
	}

	return
}

func (eb *EBookRepository) RetriveEBook(ctx context.Context, id string) (res *entities.Ebook, err error) {

	filter := bson.M{"_id": id}

	if err := eb.collection.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, err
	}

	return
}

func (eb *EBookRepository) RetriveEBookList(ctx context.Context, page, limit int64) ([]*entities.Ebook, error) {

	docs := make([]*entities.Ebook, 0, limit)

	skip := (page - 1) * limit

	filter := bson.M{"status": true}
	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.M{"number": 1})

	cursor, err := eb.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		doc := new(entities.Ebook)

		err := cursor.Decode(doc)
		if err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return docs, nil
}

func (eb *EBookRepository) EditEBook(ctx context.Context, id string, ebook *entities.Ebook) (err error) {

	edit := bson.M{
		"$set": bson.M{
			"thumbnail": bson.M{
				"alt":  ebook.Thumbnail.Alt,
				"ext":  ebook.Thumbnail.Ext,
				"path": ebook.Thumbnail.Path,
			},
			"ebook_file": bson.M{
				"alt":  ebook.EBookFile.Alt,
				"ext":  ebook.EBookFile.Ext,
				"path": ebook.EBookFile.Path,
			},
			"title":      ebook.Title,
			"status":     ebook.Status,
			"updated_at": time.Now,
		},
	}

	if _, err = eb.collection.UpdateByID(ctx, id, edit); err != nil {
		return
	}

	return
}

func (eb *EBookRepository) DeleteEBook(ctx context.Context, id string) (err error) {

	filter := bson.M{"_id": id}

	if _, err = eb.collection.DeleteOne(ctx, filter); err != nil {
		return
	}

	return
}
