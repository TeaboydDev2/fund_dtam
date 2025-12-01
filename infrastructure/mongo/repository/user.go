package repository

import (
	"context"
	"fund_dtam/domain/entities"
	"fund_dtam/domain/ports"
	mongodb "fund_dtam/infrastructure/mongo"
	"fund_dtam/infrastructure/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongodb.MongoClient) ports.UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (us *UserRepository) SaveUser(ctx context.Context, user *entities.User) error {

	userDoc, err := model.ToModel(user)
	if err != nil {
		return err
	}

	if _, err := us.collection.InsertOne(ctx, userDoc); err != nil {
		return err
	}

	return nil
}

func (us *UserRepository) RetriveUser(ctx context.Context, id string) (*entities.User, error) {

	var doc model.UserDB

	obj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": obj}

	if err := us.collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		return nil, err
	}

	return model.ToEntity(&doc), nil
}
