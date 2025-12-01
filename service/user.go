package service

import (
	"context"
	"fmt"
	"fund_dtam/domain/entities"
	"fund_dtam/domain/ports"

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

func (us *UserService) CreateUser(ctx context.Context, user *entities.User, file *entities.FileObject) error {

	filePath := fmt.Sprintf("profile_picture/%s", uuid.New().String())

	if err := us.fileRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
		return err
	}

	user.ProfilePicture.Ext = file.Ext
	user.ProfilePicture.Name = file.Name
	user.ProfilePicture.Path = filePath

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

// func AttachBaseURL(url, BucketName, path string) string {

// 	baseUrl := url + "/" + BucketName + "/" + path

// 	return baseUrl
// }
