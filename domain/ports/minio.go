package ports

import (
	"context"
	"mime/multipart"
)

type FileStorageRepository interface {
	Upload(ctx context.Context, filePath, contentType string, reader multipart.File, size int64) error
	PresignObject(ctx context.Context, objectName string) (string, error)
}
