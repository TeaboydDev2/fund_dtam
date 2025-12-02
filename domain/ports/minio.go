package ports

import (
	"context"
	"io"
	"mime/multipart"
)

type FileStorageRepository interface {
	Upload(ctx context.Context, filePath, contentType string, reader multipart.File, size int64) error
	Dowload(ctx context.Context, obJectName string) (io.ReadCloser, error)
	DeleteObject(ctx context.Context, obJectName string) error
	PresignObject(ctx context.Context, objectName string) (string, error)
}

type FileStorageService interface {
	PresignObjectServe(ctx context.Context, objectName string) (string, error)
	PresignObjectServeList(ctx context.Context, objectName []string) (map[string]string, error)
	Dowload(ctx context.Context, obJectName string) (io.ReadCloser, error)
}
