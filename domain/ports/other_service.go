package ports

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
)

type OtherSeviceRepository interface {
	SaveService(ctx context.Context, service *entities.OtherSevice) error
	RetriveService(ctx context.Context, id string) (*entities.OtherSevice, error)
	RetriveServiceList(ctx context.Context, page, limit int64) ([]*entities.OtherSevice, error)
	EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error
	EditStatus(ctx context.Context, status bool) error
	DeleteService(ctx context.Context) error
}

type OtherSevice interface {
	CreateService(ctx context.Context, user *entities.User, file *entities.FileObject)
	GetService(ctx context.Context, id string) (*entities.OtherSevice, error)
	GetServiceList(ctx context.Context) ([]*entities.OtherSevice, error)
	EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error
	EditStatus(ctx context.Context, status string) error
	DeleteService(ctx context.Context) error
}
