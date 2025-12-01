package ports

import (
	"context"
	"fund_dtam/domain/entities"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *entities.User) error
	RetriveUser(ctx context.Context, name string) (*entities.User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, id string) (*entities.User, error)
}
