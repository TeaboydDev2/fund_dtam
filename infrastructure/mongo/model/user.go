package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FirstName      string             `bson:"first_name"`
	LastName       string             `bson:"last_name"`
	ProfilePicture FileObjectDB       `bson:"file_object"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

func ToEntity(user *UserDB) *entities.User {
	return &entities.User{
		ID:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ProfilePicture: entities.FileObject{
			Alt:  user.ProfilePicture.Alt,
			Ext:  user.ProfilePicture.Ext,
			Path: user.ProfilePicture.Path,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToModel(user *entities.User) (*UserDB, error) {
	return &UserDB{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ProfilePicture: FileObjectDB{
			Alt:  user.ProfilePicture.Alt,
			Ext:  user.ProfilePicture.Ext,
			Path: user.ProfilePicture.Path,
		},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}
