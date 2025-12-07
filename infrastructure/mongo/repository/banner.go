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

type BannerRepository struct {
	collection *mongo.Collection
}

func NewBannerRepository(db *mongodb.MongoClient) ports.BannerRepository {
	return &BannerRepository{
		collection: db.Collection("banner"),
	}
}

func (bn *BannerRepository) SaveBanner(ctx context.Context, banner *entities.Banner) (err error) {

	if _, err = bn.collection.InsertOne(ctx, banner); err != nil {
		return
	}

	return
}

func (bn *BannerRepository) RetriveBanner(ctx context.Context, id string) (res *entities.Banner, err error) {

	filter := bson.M{"_id": id}

	if err = bn.collection.FindOne(ctx, filter).Decode(&res); err != nil {
		return
	}

	return
}

func (bn *BannerRepository) RetriveBannerList(ctx context.Context, page, limit int64) (res []*entities.Banner, err error) {

	banners := make([]*entities.Banner, 0, limit)

	skip := (page - 1) * limit

	filter := bson.M{"status": true}
	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.M{"position": 1})

	cursor, err := bn.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		banner := new(entities.Banner)

		err := cursor.Decode(banner)
		if err != nil {
			return nil, err
		}

		banners = append(banners, banner)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	res = banners

	return
}

func (bn *BannerRepository) EditPosition(ctx context.Context, banners []*entities.EditBannerPosition) (err error) {

	models := make([]mongo.WriteModel, 0, len(banners))

	for _, b := range banners {

		filter := bson.M{"_id": b.ID}
		update := bson.M{"$set": bson.M{"position": b.Position}}

		models = append(models, mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update))
	}

	if _, err = bn.collection.BulkWrite(ctx, models); err != nil {
		return
	}

	return
}

func (bn *BannerRepository) EditBanner(ctx context.Context, id string, banner *entities.Banner) (err error) {

	edit := bson.M{
		"$set": bson.M{
			"banner_desktop": bson.M{
				"alt":  banner.BannerDesktop.Alt,
				"ext":  banner.BannerDesktop.Ext,
				"path": banner.BannerDesktop.Path,
			},
			"banner_mobile": bson.M{
				"alt":  banner.BannerMobile.Alt,
				"ext":  banner.BannerMobile.Ext,
				"path": banner.BannerMobile.Path,
			},
			"link_url":   banner.LinkUrl,
			"status":     banner.Status,
			"updated_at": time.Now(),
		},
	}

	if _, err = bn.collection.UpdateByID(ctx, id, edit); err != nil {
		return
	}

	return
}

func (bn *BannerRepository) DeleteBanner(ctx context.Context, id string) (err error) {

	filter := bson.M{"_id": id}

	if _, err = bn.collection.DeleteOne(ctx, filter); err != nil {
		return
	}

	return
}
