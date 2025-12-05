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

type OtherSeviceRepository struct {
	collection *mongo.Collection
}

func NewOtherServiceRepository(
	db *mongodb.MongoClient,
) ports.OtherSeviceRepository {
	return &OtherSeviceRepository{
		collection: db.Collection("other_service"),
	}
}

func (ots *OtherSeviceRepository) SaveService(ctx context.Context, service *entities.OtherSevice) error {

	if _, err := ots.collection.InsertOne(ctx, service); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSeviceRepository) RetriveService(ctx context.Context, id string) (res *entities.OtherSevice, err error) {

	filter := bson.M{"_id": id}

	if err = ots.collection.FindOne(ctx, filter).Decode(&res); err != nil {
		return
	}

	return
}

func (ots *OtherSeviceRepository) RetriveServiceList(ctx context.Context, page, limit int64) ([]*entities.OtherSevice, error) {

	res := make([]*entities.OtherSevice, 0, limit)

	skip := (page - 1) * limit

	filter := bson.M{"status": true}
	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.M{"number": 1})

	cursor, err := ots.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (ots *OtherSeviceRepository) EditService(ctx context.Context, id string, service *entities.OtherSevice) (err error) {

	update := bson.M{
		"$set": bson.M{
			"thumbnail": bson.M{
				"alt":  service.Thumbnail.Alt,
				"ext":  service.Thumbnail.Ext,
				"path": service.Thumbnail.Path,
			},
			"title":      service.Title,
			"url":        service.Url,
			"status":     service.Status,
			"updated_at": time.Now(),
		},
	}

	result, err := ots.collection.UpdateByID(ctx, id, update)
	if err != nil {
		return
	}

	if result.MatchedCount == 0 {
		return
	}

	return
}

func (ots *OtherSeviceRepository) EditSortNumber(ctx context.Context, service []*entities.OtherSevice) (err error) {
	return
}

func (ots *OtherSeviceRepository) EditStatus(ctx context.Context, id string, status bool) (err error) {

	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	result, err := ots.collection.UpdateByID(ctx, id, update)
	if err != nil {
		return
	}

	if result.MatchedCount == 0 {
		return
	}

	return
}

func (ots *OtherSeviceRepository) IncreaseViewStatic(ctx context.Context, id string) (err error) {

	filter := bson.M{"_id": id}
	inc := bson.M{"$inc": bson.M{"view_static": 1}}

	if _, err = ots.collection.UpdateOne(ctx, filter, inc); err != nil {
		return
	}

	return
}

func (ots *OtherSeviceRepository) DeleteService(ctx context.Context, id string) (err error) {

	filter := bson.M{"_id": id}

	if _, err = ots.collection.DeleteOne(ctx, filter); err != nil {
		return
	}

	return
}
