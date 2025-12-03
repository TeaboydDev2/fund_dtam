package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"
)

type CreateUser struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

type UserJson struct {
	ID             string        `json:"id"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	PictureProfile FileObject    `json:"picture_profile"`
	Illustration   []*FileObject `json:"illustration"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

func ToEntity(user *CreateUser) *entities.User {
	return &entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToResponse(user *entities.User, picProfile string) *UserJson {
	return &UserJson{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PictureProfile: FileObject{
			Alt:  user.ProfilePicture.Alt,
			Ext:  user.ProfilePicture.Ext,
			Path: picProfile,
		},
		Illustration: FileToResponse(user.Illustration),
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
