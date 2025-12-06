package service

import (
	"context"
	"dtam-fund-cms-backend/config"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	"dtam-fund-cms-backend/service/helper"
	"fmt"

	"github.com/google/uuid"
)

type EBookService struct {
	ebookRepository       ports.EBookRepository
	fileStorageRepository ports.FileStorageRepository
	cfg                   *config.Minio
}

func NewEBookService(
	ebookRepository ports.EBookRepository,
	fileStorageRepository ports.FileStorageRepository,
	cfg *config.Minio,
) ports.EBookService {
	return &EBookService{
		ebookRepository:       ebookRepository,
		fileStorageRepository: fileStorageRepository,
		cfg:                   cfg,
	}
}

func (eb *EBookService) CreateEBook(ctx context.Context, ebook *entities.Ebook, thumbnail, file *entities.FileObject) (err error) {

	if thumbnail == nil {
		return fmt.Errorf("thumbnail is required")
	}

	if file == nil {
		return fmt.Errorf("file is required")
	}

	thumbnailPath := eb.ebookFilePath("thumbnail")
	filePath := eb.ebookFilePath("file")

	ebook.Thumbnail = cloneFile(thumbnail, thumbnailPath)
	ebook.EBookFile = cloneFile(file, filePath)

	ebook.ID = uuid.NewString()

	if err = eb.fileStorageRepository.Upload(ctx, thumbnailPath, thumbnail.ContentType, thumbnail.File, thumbnail.Size); err != nil {
		return
	}

	if err = eb.fileStorageRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
		return
	}

	if err = eb.ebookRepository.SaveEBook(ctx, ebook); err != nil {

		_ = eb.fileStorageRepository.DeleteObject(ctx, thumbnailPath)
		_ = eb.fileStorageRepository.DeleteObject(ctx, filePath)

		return
	}

	return
}

func (eb *EBookService) GetEBook(ctx context.Context, id string) (res *entities.Ebook, err error) {

	ebook, err := eb.ebookRepository.RetriveEBook(ctx, id)
	if err != nil {
		return
	}

	helper.AttachBaseURL(eb.cfg.BaseUrlFile, eb.cfg.BucketName, &ebook.Thumbnail.Path)

	res = ebook

	return
}

func (eb *EBookService) GetEBookList(ctx context.Context, page, limit string) (res []*entities.Ebook, err error) {

	parsePage := helper.StrToInt(page, int64(1))
	parseLimit := helper.StrToInt(limit, int64(10))

	ebookList, err := eb.ebookRepository.RetriveEBookList(ctx, parsePage, parseLimit)
	if err != nil {
		return
	}

	for i := range ebookList {
		helper.AttachBaseURL(eb.cfg.BaseUrlFile, eb.cfg.BucketName, &ebookList[i].Thumbnail.Path)
		helper.AttachBaseURL(eb.cfg.BaseUrlFile, eb.cfg.BucketName, &ebookList[i].EBookFile.Path)
	}

	res = ebookList

	return
}

func (eb *EBookService) EditEBook(ctx context.Context, id string, ebook *entities.Ebook, thumbnail, file *entities.FileObject) (err error) {
	oldEbook, err := eb.ebookRepository.RetriveEBook(ctx, id)
	if err != nil {
		return
	}

	thumbnailPath := eb.ebookFilePath("thumbnail")
	filePath := eb.ebookFilePath("file")

	if thumbnail != nil {
		if err = eb.fileStorageRepository.Upload(ctx, thumbnailPath, thumbnail.ContentType, thumbnail.File, thumbnail.Size); err != nil {
			return
		}
		oldEbook.Thumbnail = cloneFile(thumbnail, thumbnailPath)
	}

	if file != nil {
		if err = eb.fileStorageRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
			return
		}
		oldEbook.EBookFile = cloneFile(file, filePath)
	}

	if ebook.Title == "" {
		ebook.Title = oldEbook.Title
	}

	return
}

func (eb *EBookService) DeleteEBook(ctx context.Context, id string) (err error) {

	ebook, err := eb.ebookRepository.RetriveEBook(ctx, id)
	if err != nil {
		return
	}

	if err = eb.fileStorageRepository.DeleteObject(ctx, ebook.Thumbnail.Path); err != nil {
		return
	}

	if err = eb.fileStorageRepository.DeleteObject(ctx, ebook.EBookFile.Path); err != nil {
		return
	}

	if err = eb.ebookRepository.DeleteEBook(ctx, id); err != nil {
		return
	}

	return
}

func (eb *EBookService) ebookFilePath(fileType string) (res string) {

	id := uuid.NewString()

	switch fileType {
	case "thumbnail":
		return fmt.Sprintf("ebook/thumbnail/%s", id)
	case "file":
		return fmt.Sprintf("ebook/file/%s", id)
	}

	return
}
