package main

import (
	"context"
	cfg "dtam-fund-cms-backend/config"
	fiber "dtam-fund-cms-backend/infrastructure/fiber"
	"dtam-fund-cms-backend/infrastructure/logger"
	minio_obj "dtam-fund-cms-backend/infrastructure/minio"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dotenv, err := cfg.SetUpEnviroment()
	if err != nil {
		log.Fatalf("Cannot Load Env.")
	}

	log.Println("ENV has been loaded")

	validator := validator.New()

	logger := logger.EstablishZeroLogger(*dotenv.App)

	mongo, err := mongodb.EstablishConnection(ctx, dotenv.Mongo)
	if err != nil {
		log.Fatalf("[Startup Error] Unable to connect to MongoDB: %v", err)
	}
	defer mongo.Close(ctx)

	minio, err := minio_obj.EstablishConnection(ctx, dotenv.Minio)
	if err != nil {
		log.Fatalf("[Startup Error] Unable to connect to Minio: %v", err)
	}

	err = fiber.Start(ctx, *dotenv, logger, mongo, minio, validator)
	if err != nil {
		log.Fatalf("Start failed:%v", err)
	}
}
