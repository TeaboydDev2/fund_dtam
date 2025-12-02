package minio_obj

import (
	"context"
	"dtam-fund-cms-backend/config"
	"dtam-fund-cms-backend/domain/ports"
	"log"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	minioClient *minio.Client
	bucketName  string
}

func EstablishConnection(ctx context.Context, cfg *config.Minio) (*MinioClient, error) {

	client, err := minio.New(cfg.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretAccessKey, ""),
		Secure: cfg.Secure,
	})

	if err != nil {
		return nil, err
	}

	pingCheck, err := client.BucketExists(ctx, cfg.BucketName)
	if err != nil {
		log.Fatalf("Failed to check bucket: %v", err)
		return nil, err
	}

	if pingCheck {
		log.Println("Connected! Bucket exists:", cfg.BucketName)
	} else {
		log.Println("Connected! Bucket does not exist:", cfg.BucketName)
	}

	return &MinioClient{
		minioClient: client,
		bucketName:  cfg.BucketName,
	}, nil
}

func NewMinioRepository(minioClient *MinioClient) ports.FileStorageRepository {
	return minioClient
}

func (mio *MinioClient) Upload(ctx context.Context, filePath, contentType string, reader multipart.File, size int64) error {

	_, err := mio.minioClient.PutObject(ctx, mio.bucketName, filePath, reader, size, minio.PutObjectOptions{
		PartSize:     64 * 1024 * 1024,
		CacheControl: "public, max-age=604800",
		ContentType:  contentType,
	})

	if err != nil {
		return err
	}

	return nil
}

func (mio *MinioClient) PresignObject(ctx context.Context, objectName string) (string, error) {

	reqParams := make(url.Values)

	presignedUrl, err := mio.minioClient.PresignedGetObject(ctx, mio.bucketName, objectName, 10*time.Minute, reqParams)
	if err != nil {
		return "", err
	}

	return presignedUrl.String(), nil
}
