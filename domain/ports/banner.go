package ports

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
)

type BannerRepository interface {
	SaveBanner(ctx context.Context, banner *entities.Banner) error
	RetriveBanner(ctx context.Context, id string) (*entities.Banner, error)
	RetriveBannerList(ctx context.Context, page, limit int64) ([]*entities.Banner, error)
	EditPosition(ctx context.Context, banners []*entities.EditBannerPosition) error
	EditBanner(ctx context.Context, id string, banner *entities.Banner) error
	DeleteBanner(ctx context.Context, id string) error
}

type BannerService interface {
	CreateBanner(ctx context.Context, banner *entities.Banner, desktop, mobile *entities.FileObject) error
	GetBanner(ctx context.Context, id string) (*entities.Banner, error)
	GetBannerList(ctx context.Context, page, limit string) ([]*entities.Banner, error)
	EditPosition(ctx context.Context, banners []*entities.EditBannerPosition) error
	EditBanner(ctx context.Context, id string, banner *entities.Banner, desktop, mobile *entities.FileObject) error
	DeleteBanner(ctx context.Context, id string) error
}
