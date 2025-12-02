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

	return nil, nil
}

func (ots *OtherSeviceRepository) EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error {
	return nil
}

func (ots *OtherSeviceRepository) EditStatus(ctx context.Context, status bool) error {
	return nil
}

func (ots *OtherSeviceRepository) DeleteService(ctx context.Context) error {
	return nil
}
