package model

import (
	"fund_dtam/domain/entities"
	"time"
)

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserJson struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToEntity(user *CreateUser) *entities.User {
	return &entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToResponse(user *entities.User) *UserJson {
	return &UserJson{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
