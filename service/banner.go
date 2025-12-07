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

type BannerService struct {
	bannerRepository      ports.BannerRepository
	fileStorageRepository ports.FileStorageRepository
	logger                ports.Logger
	cfg                   *config.Minio
}

func NewBannerService(
	bannerRepository ports.BannerRepository,
	fileStorageRepository ports.FileStorageRepository,
	logger ports.Logger,
	cfg *config.Minio,
) ports.BannerService {
	return &BannerService{
		bannerRepository:      bannerRepository,
		fileStorageRepository: fileStorageRepository,
		logger:                logger,
		cfg:                   cfg,
	}
}

func (bn *BannerService) CreateBanner(ctx context.Context, banner *entities.Banner, desktop, mobile *entities.FileObject) error {

	if desktop == nil {
		return fmt.Errorf("desktop banner is required")
	}

	if mobile == nil {
		return fmt.Errorf("mobile banner is required")
	}

	desktopPath := bn.bannerFilePath("desktop")
	mobilePath := bn.bannerFilePath("mobile")

	banner.BannerDesktop = *cloneFile(desktop, desktopPath)
	banner.BannerMobile = *cloneFile(mobile, mobilePath)

	banner.ID = uuid.NewString()

	if err := bn.fileStorageRepository.Upload(ctx, desktopPath, desktop.ContentType, desktop.File, desktop.Size); err != nil {

		bn.logger.ErrorF("upload desktop banner failed", err, map[string]interface{}{
			"path": desktopPath,
		})
		return fmt.Errorf("upload desktop banner failed: %w", err)
	}

	bn.logger.Info("desktop_banner upload successfully", map[string]interface{}{
		"banner_id": banner.ID,
		"path":      desktopPath,
	})

	if err := bn.fileStorageRepository.Upload(ctx, mobilePath, mobile.ContentType, mobile.File, mobile.Size); err != nil {

		bn.logger.ErrorF("upload mobile banner failed", err, map[string]interface{}{
			"path": mobilePath,
		})

		_ = bn.fileStorageRepository.DeleteObject(ctx, desktopPath) // rollback
		return fmt.Errorf("upload mobile banner failed: %w", err)
	}

	bn.logger.Info("mobile banner uploaded", map[string]interface{}{
		"banner_id": banner.ID,
		"path":      mobilePath,
	})

	if err := bn.bannerRepository.SaveBanner(ctx, banner); err != nil {

		bn.logger.ErrorF("cannot save banner entity", err, map[string]interface{}{
			"banner_id": banner.ID,
		})

		_ = bn.fileStorageRepository.DeleteObject(ctx, desktopPath)
		_ = bn.fileStorageRepository.DeleteObject(ctx, mobilePath)

		return fmt.Errorf("cannot save banner: %w", err)
	}

	bn.logger.Info("banner created successfully", map[string]interface{}{
		"banner_id": banner.ID,
	})

	return nil
}

func (bn *BannerService) GetBanner(ctx context.Context, id string) (res *entities.Banner, err error) {

	banner, err := bn.bannerRepository.RetriveBanner(ctx, id)
	if err != nil {
		return
	}

	helper.AttachBaseURL(bn.cfg.BaseUrlFile, bn.cfg.BucketName, &banner.BannerDesktop.Path)
	helper.AttachBaseURL(bn.cfg.BaseUrlFile, bn.cfg.BucketName, &banner.BannerMobile.Path)

	return
}

func (bn *BannerService) GetBannerList(ctx context.Context, page, limit string) ([]*entities.Banner, error) {
	return nil, nil
}

func (bn *BannerService) EditPosition(ctx context.Context, banners []*entities.EditBannerPosition) error {
	return nil
}

func (bn *BannerService) EditBanner(ctx context.Context, id string, banner *entities.Banner, desktop, mobile *entities.FileObject) error {
	return nil
}

func (bn *BannerService) DeleteBanner(ctx context.Context, id string) error {
	return nil
}

func (eb *BannerService) bannerFilePath(fileType string) (res string) {

	id := uuid.NewString()

	switch fileType {
	case "desktop":
		return fmt.Sprintf("banner/desktop/%s", id)
	case "mobile":
		return fmt.Sprintf("banner/mobile/%s", id)
	}

	return
}
