package repository

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StatViewRepository struct {
	collection *mongo.Collection
}

func NewStatViewRepository(
	db *mongodb.MongoClient,
) ports.StatViewRepository {
	return &StatViewRepository{
		collection: db.Collection("web_view"),
	}
}

func (r *StatViewRepository) IncreaseWebView(ctx context.Context, date time.Time) (err error) {

	filter := bson.M{
		"date": date,
	}

	update := bson.M{
		"$setOnInsert": entities.WebView{
			Date:      date,
			CreatedAt: time.Now(),
		},
		"$inc": bson.M{
			"view_count": 1,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = r.collection.UpdateOne(ctx, filter, update, opts)
	return
}

func (r *StatViewRepository) QueryWebViewStat(ctx context.Context, date time.Time) (res entities.StatWebView, err error) {

	thisMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	thisYear := time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location())

	mapMatch := map[string]bson.M{}

	mapMatch["to_day"] = bson.M{
		"date": date,
	}
	mapMatch["yesterday"] = bson.M{
		"date": date.AddDate(0, 0, -1),
	}
	mapMatch["this_month"] = bson.M{
		"date": bson.M{
			"$gte": thisMonth,
		},
	}
	mapMatch["previous_month"] = bson.M{
		"date": bson.M{
			"$gte": thisMonth.AddDate(0, -1, 0),
			"$lt":  thisMonth,
		},
	}
	mapMatch["this_year"] = bson.M{
		"date": bson.M{
			"$gte": thisYear,
		},
	}
	mapMatch["previous_year"] = bson.M{
		"date": bson.M{
			"$gte": thisYear.AddDate(-1, 0, 0),
			"$lt":  thisYear,
		},
	}

	facetStage := bson.M{}
	project := bson.M{}

	for key, value := range mapMatch {
		facetStage[key] = []bson.M{
			{"$match": value},
			{"$group": bson.M{
				"_id": nil,
				"count": bson.M{
					"$sum": "$view_count",
				},
			}},
		}
		project[key] = bson.M{"$first": fmt.Sprintf("$%s.count", key)}

	}
	facetStage["all_time"] = []bson.M{
		{"$group": bson.M{
			"_id": nil,
			"count": bson.M{
				"$sum": "$view_count",
			},
		}},
	}
	project["all_time"] = bson.M{"$first": "$all_time.count"}

	pl := []bson.M{
		{"$facet": facetStage},
		{"$project": project},
	}

	curr, err := r.collection.Aggregate(ctx, pl)
	if err != nil {
		return
	}
	defer curr.Close(ctx)

	for curr.Next(ctx) {
		err = curr.Decode(&res)
		return
	}

	return
}
