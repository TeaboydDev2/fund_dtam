package service

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	"io"
)

type FileObjectSevice struct {
	fileStorageRepository ports.FileStorageRepository
}

func NewFileObjectService(
	fileStorageRepository ports.FileStorageRepository,
) ports.FileStorageService {
	return &FileObjectSevice{
		fileStorageRepository: fileStorageRepository,
	}
}

func (mio *FileObjectSevice) Dowload(ctx context.Context, obJectName string) (io.ReadCloser, error) {
	return mio.fileStorageRepository.Dowload(ctx, obJectName)
}

func (mio *FileObjectSevice) PresignObjectServe(ctx context.Context, objectName string) (string, error) {
	return mio.fileStorageRepository.PresignObject(ctx, objectName)
}

func (mio *FileObjectSevice) PresignObjectServeList(ctx context.Context, objectName []string) (map[string]string, error) {

	urlList := make(map[string]string, len(objectName))

	for _, v := range objectName {
		url, err := mio.fileStorageRepository.PresignObject(ctx, v)
		if err != nil {
			return nil, err
		}

		urlList[v] = url
	}

	return urlList, nil
}

func cloneFile(src *entities.FileObject, path string) *entities.FileObject {

	n := *src
	n.Path = path

	return &n
}
