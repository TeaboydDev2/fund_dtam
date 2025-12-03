package service

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository ports.UserRepository
	fileRepository ports.FileStorageRepository
}

func NewUserService(
	userRepository ports.UserRepository,
	fileRepository ports.FileStorageRepository,
) ports.UserService {
	return &UserService{
		userRepository: userRepository,
		fileRepository: fileRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *entities.User, file *entities.FileObject, illustration []*entities.FileObject) error {

	filePath := fmt.Sprintf("profile_picture/%s", uuid.New().String())

	if err := us.fileRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
		return err
	}

	user.ProfilePicture.Ext = file.Ext
	user.ProfilePicture.Alt = file.Alt
	user.ProfilePicture.Path = filePath

	user.Illustration = make([]*entities.FileObject, len(illustration))

	for i, v := range illustration {

		multiFilePath := fmt.Sprintf("illustration_picture/%s", uuid.New().String())

		if err := us.fileRepository.Upload(ctx, multiFilePath, v.ContentType, v.File, v.Size); err != nil {
			return err
		}

		user.Illustration[i] = &entities.FileObject{
			Alt:  v.Alt,
			Ext:  v.Ext,
			Path: multiFilePath,
		}
	}

	return us.userRepository.SaveUser(ctx, user)
}

func (us *UserService) GetUser(ctx context.Context, id string) (*entities.User, string, error) {

	user, err := us.userRepository.RetriveUser(ctx, id)
	if err != nil {
		return nil, "", err
	}

	pictureProfile, err := us.fileRepository.PresignObject(ctx, user.ProfilePicture.Path)
	if err != nil {
		return nil, "", err
	}

	return user, pictureProfile, nil
}
