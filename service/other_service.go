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

type OtherSevice struct {
	otherServiceRepository ports.OtherSeviceRepository
	fileStorageRepository  ports.FileStorageRepository
	cfg                    *config.Minio
}

func NewOtherService(
	otherServiceRepository ports.OtherSeviceRepository,
	fileStorageRepository ports.FileStorageRepository,
	cfg *config.Minio,
) ports.OtherSevice {
	return &OtherSevice{
		otherServiceRepository: otherServiceRepository,
		fileStorageRepository:  fileStorageRepository,
		cfg:                    cfg,
	}
}

func (ots *OtherSevice) CreateService(ctx context.Context, service *entities.OtherSevice, file *entities.FileObject) error {

	filePath := fmt.Sprintf("service_thumbnail/%s", uuid.New().String())

	if err := ots.fileStorageRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
		return err
	}

	thumbnail := entities.FileObject{
		Alt:  file.Alt,
		Ext:  file.Ext,
		Path: filePath,
	}

	service.Thumbnail = &thumbnail
	service.ID = uuid.NewString()

	if err := ots.otherServiceRepository.SaveService(ctx, service); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSevice) GetService(ctx context.Context, id string) (*entities.OtherSevice, error) {

	service, err := ots.otherServiceRepository.RetriveService(ctx, id)
	if err != nil {
		return nil, err
	}

	service.Thumbnail.Path = helper.AttachBaseURL(ots.cfg.BaseUrlFile, ots.cfg.BucketName, service.Thumbnail.Path)

	return service, nil
}

func (ots *OtherSevice) GetServiceList(ctx context.Context, page, limit string) ([]*entities.OtherSevice, error) {

	parsePage := helper.StrToInt(page, int64(1))
	parseLimit := helper.StrToInt(limit, int64(6))

	serviceList, err := ots.otherServiceRepository.RetriveServiceList(ctx, parsePage, parseLimit)
	if err != nil {
		return nil, err
	}

	for _, v := range serviceList {
		v.Thumbnail.Path = helper.AttachBaseURL(ots.cfg.BaseUrlFile, ots.cfg.BucketName, v.Thumbnail.Path)
	}

	return serviceList, nil
}

func (ots *OtherSevice) EditService(ctx context.Context, id string, service *entities.OtherSevice, file *entities.FileObject) error {

	oldService, err := ots.otherServiceRepository.RetriveService(ctx, id)
	if err != nil {
		return err
	}

	if file != nil {

		filePath := fmt.Sprintf("service_thumbnail/%s", uuid.New().String())

		if err := ots.fileStorageRepository.DeleteObject(ctx, oldService.Thumbnail.Path); err != nil {
			return err
		}

		if err := ots.fileStorageRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
			return err
		}

		thumbnail := entities.FileObject{
			Alt:  file.Alt,
			Ext:  file.Ext,
			Path: filePath,
		}

		service.Thumbnail = &thumbnail
	} else {
		service.Thumbnail = oldService.Thumbnail
	}

	if service.Title == "" {
		service.Title = oldService.Title
	}

	if service.Url == "" {
		service.Url = oldService.Url
	}

	oldService.Status = service.Status

	if err := ots.otherServiceRepository.EditService(ctx, oldService.ID, service); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSevice) EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error {
	return nil
}

func (ots *OtherSevice) EditStatus(ctx context.Context, id string, status bool) error {

	if err := ots.otherServiceRepository.EditStatus(ctx, id, status); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSevice) IncreaseViewStatic(ctx context.Context, id string) error {

	if err := ots.otherServiceRepository.IncreaseViewStatic(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ots *OtherSevice) DeleteService(ctx context.Context, id string) error {

	service, err := ots.otherServiceRepository.RetriveService(ctx, id)
	if err != nil {
		return err
	}

	if err := ots.fileStorageRepository.DeleteObject(ctx, service.Thumbnail.Path); err != nil {
		return err
	}

	if err := ots.otherServiceRepository.DeleteService(ctx, service.ID); err != nil {
		return err
	}

	return nil
}
