package ports

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *entities.User) error
	RetriveUser(ctx context.Context, name string) (*entities.User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User, file *entities.FileObject, illustration []*entities.FileObject) error
	GetUser(ctx context.Context, id string) (*entities.User, string, error)
}
