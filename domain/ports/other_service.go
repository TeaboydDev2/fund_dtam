package ports

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
)

type OtherSeviceRepository interface {
	SaveService(ctx context.Context, service *entities.OtherSevice) error
	RetriveService(ctx context.Context, id string) (*entities.OtherSevice, error)
	RetriveServiceList(ctx context.Context, page, limit int64) ([]*entities.OtherSevice, error)
	EditService(ctx context.Context, id string, service *entities.OtherSevice) error
	EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error
	EditStatus(ctx context.Context, id string, status bool) error
	IncreaseViewStatic(ctx context.Context, id string) error
	DeleteService(ctx context.Context, id string) error
}

type OtherSevice interface {
	CreateService(ctx context.Context, service *entities.OtherSevice, file *entities.FileObject) error
	GetService(ctx context.Context, id string) (*entities.OtherSevice, error)
	GetServiceList(ctx context.Context, page, limit string) ([]*entities.OtherSevice, []string, error)
	EditService(ctx context.Context, id string, service *entities.OtherSevice, file *entities.FileObject) error
	EditSortNumber(ctx context.Context, service []*entities.OtherSevice) error
	EditStatus(ctx context.Context, id string, status bool) error
	IncreaseViewStatic(ctx context.Context, id string) error
	DeleteService(ctx context.Context, id string) error
}
