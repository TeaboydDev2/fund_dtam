package repository

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"dtam-fund-cms-backend/infrastructure/mongo/helper"
	"dtam-fund-cms-backend/infrastructure/mongo/model"

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

func (eb *EBookRepository) SaveEBook(ctx context.Context, ebook *entities.Ebook) error {

	if _, err := eb.collection.InsertOne(ctx, model.EbookToModel(ebook)); err != nil {
		return err
	}

	return nil
}
func (eb *EBookRepository) RetriveEBook(ctx context.Context, id string) (*entities.Ebook, error) {

	doc := new(model.EBookDB)

	obj, err := helper.ToPrimitiveObj(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": obj}

	if err := eb.collection.FindOne(ctx, filter).Decode(doc); err != nil {
		return nil, err
	}

	return model.EBookToEntity(doc), nil
}

func (eb *EBookRepository) RetriveEBookList(ctx context.Context, page, limit int64) ([]*entities.Ebook, error) {

	docs := make([]*model.EBookDB, 0, 50)

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

	for cursor.Next(ctx) {

		var doc model.EBookDB

		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}

		docs = append(docs, &doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return model.EBookToEntityList(docs), nil
}

func (eb *EBookRepository) EditEBook(ctx context.Context, id string, ebook *entities.Ebook) error {
	return nil
}
func (eb *EBookRepository) DeleteEBook(ctx context.Context, id string) error {
	return nil
}
