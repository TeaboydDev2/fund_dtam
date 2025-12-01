package ports

import (
	"context"
	"io"
)

type FileStorageRepository interface {
	Upload(ctx context.Context, fileName string, reader io.Reader, size int64) error
}
