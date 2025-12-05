package config

import (
	cfh "dtam-fund-cms-backend/config/helper"
	"fmt"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App   *App
		HTTP  *HTTP
		Mongo *Mongo
		Minio *Minio
	}

	App struct {
		Stage string
	}

	HTTP struct {
		Host          string
		Port          string
		HttpOrigin    string
		AllowedOrigin string
		Prefix        string
		BodyLimit     int
	}

	Mongo struct {
		Uri          string
		DatabaseName string
	}

	Minio struct {
		Host            string
		AccessKey       string
		SecretAccessKey string
		Secure          bool
		BucketName      string
		BaseUrlFile     string
	}
)

func SetUpEnviroment() (*Container, error) {

	err := godotenv.Load("./.env")
	if err != nil {
		return nil, err
	}

	app := &App{
		Stage: cfh.GetEnv("APP_STAGE", "development"),
	}

	http := &HTTP{
		Host:          cfh.GetEnv("SERVER_HOST", ""),
		Port:          cfh.GetEnv("SERVER_PORT", ""),
		HttpOrigin:    cfh.GetEnv("SERVER_ORIGIN", ""),
		AllowedOrigin: cfh.GetEnv("ALLOWED_ORIGIN", ""),
		Prefix:        cfh.GetEnv("PREFIX_API", "api"),
		BodyLimit:     cfh.ParseString(cfh.GetEnv("BODY_LIMIT", ""), 10),
	}

	mongo := &Mongo{
		Uri:          cfh.GetEnv("MONGO_URI", ""),
		DatabaseName: cfh.GetEnv("DB_NAME", ""),
	}

	ssl := ""

	switch app.Stage {
	case "development":
		ssl = "http"
	case "production":
		ssl = "https"
	}

	minio := &Minio{
		Host:            cfh.GetEnv("MINIO_HOST", ""),
		AccessKey:       cfh.GetEnv("MINIO_ACCESS_KEY", ""),
		SecretAccessKey: cfh.GetEnv("MINIO_SECRET_ACCESS_KEY", ""),
		Secure:          cfh.ParseString(cfh.GetEnv("MINIO_SECURE", ""), false),
		BucketName:      cfh.GetEnv("MINIO_BUCKET_NAME", ""),
		BaseUrlFile:     fmt.Sprintf("%s://%s:%s", ssl, http.Host, http.Port),
	}

	return &Container{
		App:   app,
		HTTP:  http,
		Mongo: mongo,
		Minio: minio,
	}, nil
}
