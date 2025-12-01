package mongodb

import (
	"context"
	"fund_dtam/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	mongo  *mongo.Client
	dbName string
}

func EstablishConnection(ctx context.Context, cfg *config.Mongo) (*MongoClient, error) {

	mongoUri := cfg.Uri

	clientOption := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Println("Failed to connect to MongoDB")
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Ping failed: Cannot connect to MongoDB")
		return nil, err
	} else {
		log.Println("Connected to MongoDB successfully")
	}

	return &MongoClient{mongo: client, dbName: cfg.DatabaseName}, nil
}

func (mgdb *MongoClient) Collection(name string) *mongo.Collection {
	return mgdb.mongo.Database(mgdb.dbName).Collection(name)
}

func (mgdb *MongoClient) Close(ctx context.Context) error {
	if err := mgdb.mongo.Disconnect(ctx); err != nil {
		log.Println("Database close error :", err)
		return err
	}
	log.Println("Database close!")
	return nil
}
