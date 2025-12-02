package repository

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"dtam-fund-cms-backend/infrastructure/mongo/helper"
	"dtam-fund-cms-backend/infrastructure/mongo/model"
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

	if _, err := ots.collection.InsertOne(ctx, model.ToModelService(service)); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSeviceRepository) RetriveService(ctx context.Context, id string) (*entities.OtherSevice, error) {

	var doc model.OtherSeviceDB

	obj, err := helper.ToPrimitiveObj(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": obj}

	if err := ots.collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		return nil, err
	}

	return model.ToEntityService(&doc), nil
}

func (ots *OtherSeviceRepository) RetriveServiceList(ctx context.Context, page, limit int64) ([]*entities.OtherSevice, error) {

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

	var docs []*model.OtherSeviceDB
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := model.ToEntityServiceList(docs)

	return result, nil
}

func (ots *OtherSeviceRepository) EditService(ctx context.Context, id string, service *entities.OtherSevice) error {

	obj, err := helper.ToPrimitiveObj(id)
	if err != nil {
		return err
	}

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

	result, err := ots.collection.UpdateByID(ctx, obj, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return err
	}

	return nil
}

func (ots *OtherSeviceRepository) EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error {
	return nil
}

func (ots *OtherSeviceRepository) EditStatus(ctx context.Context, id string, status bool) error {

	obj, err := helper.ToPrimitiveObj(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	result, err := ots.collection.UpdateByID(ctx, obj, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return err
	}

	return nil
}

func (ots *OtherSeviceRepository) IncreaseViewStatic(ctx context.Context, id string) error {

	obj, err := helper.ToPrimitiveObj(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": obj}
	inc := bson.M{"$inc": bson.M{"view_static": 1}}

	if _, err := ots.collection.UpdateOne(ctx, filter, inc); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSeviceRepository) DeleteService(ctx context.Context, id string) error {

	obj, err := helper.ToPrimitiveObj(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": obj}

	if _, err := ots.collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
