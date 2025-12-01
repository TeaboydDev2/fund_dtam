package service

import (
	"context"
	"fund_dtam/domain/entities"
	"fund_dtam/domain/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(
	userRepository ports.UserRepository,
) ports.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *entities.User) error {
	return us.userRepository.SaveUser(ctx, user)
}

func (us *UserService) GetUser(ctx context.Context, id string) (*entities.User, error) {
	return us.userRepository.RetriveUser(ctx, id)
}
