package minio_obj

import (
	"context"
	"fund_dtam/config"
	"fund_dtam/domain/ports"
	"io"
	"log"

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

func (mio *MinioClient) Upload(ctx context.Context, fileName string, reader io.Reader, size int64) error {

	_, err := mio.minioClient.PutObject(ctx, mio.bucketName, fileName, reader, size, minio.PutObjectOptions{
		PartSize: 64 * 1024 * 1024,
	})

	if err != nil {
		return err
	}

	return nil
}
