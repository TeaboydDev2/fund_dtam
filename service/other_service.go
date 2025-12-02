package service

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	"dtam-fund-cms-backend/service/helper"
	"fmt"

	"github.com/google/uuid"
)

type OtherSevice struct {
	otherServiceRepository ports.OtherSeviceRepository
	fileStorageRepository  ports.FileStorageRepository
}

func NewOtherService(
	otherServiceRepository ports.OtherSeviceRepository,
	fileStorageRepository ports.FileStorageRepository,
) ports.OtherSevice {
	return &OtherSevice{
		otherServiceRepository: otherServiceRepository,
		fileStorageRepository:  fileStorageRepository,
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

	service.Thumbnail = thumbnail

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

	return service, nil
}

func (ots *OtherSevice) GetServiceList(ctx context.Context, page, limit string) ([]*entities.OtherSevice, []string, error) {

	parsePage := helper.StrToInt(page, int64(1))
	parseLimit := helper.StrToInt(limit, int64(6))

	serviceList, err := ots.otherServiceRepository.RetriveServiceList(ctx, parsePage, parseLimit)
	if err != nil {
		return nil, nil, err
	}

	thumbnail := make([]string, 0, len(serviceList))

	for _, s := range serviceList {
		thumbnail = append(thumbnail, s.Thumbnail.Path)
	}

	return serviceList, thumbnail, nil
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

		service.Thumbnail = thumbnail
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
