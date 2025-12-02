package main

import (
	"context"
	cfg "dtam-fund-cms-backend/config"
	fiber "dtam-fund-cms-backend/infrastructure/fiber"
	minio_obj "dtam-fund-cms-backend/infrastructure/minio"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"log"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dotenv, err := cfg.SetUpEnviroment()
	if err != nil {
		log.Println("ENV NOT FOUND !")
		os.Exit(1)
	}

	log.Println("ENV has been loaded")

	mongo, err := mongodb.EstablishConnection(ctx, dotenv.Mongo)
	if err != nil {
		os.Exit(1)
	}
	defer mongo.Close(ctx)

	minio, err := minio_obj.EstablishConnection(ctx, dotenv.Minio)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	err = fiber.Start(ctx, *dotenv.HTTP, mongo, minio)
	if err != nil {
		os.Exit(1)
	}
}
