package main

import (
	"context"
	"fmt"
	cfg "fund_dtam/config"
	fiber "fund_dtam/infrastructure/fiber"
	minio_obj "fund_dtam/infrastructure/minio"
	mongodb "fund_dtam/infrastructure/mongo"
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
		os.Exit(1)
	}

	fmt.Printf("minio: %v\n", minio)

	err = fiber.Start(ctx, *dotenv.HTTP, mongo)
	if err != nil {
		os.Exit(1)
	}
}
