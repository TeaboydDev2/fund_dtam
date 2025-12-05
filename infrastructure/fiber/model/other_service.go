package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"
)

type CreateOtherService struct {
	Title string `form:"title"`
	Url   string `form:"url"`
}

func OtherServiceToEntity(otherService *CreateOtherService) *entities.OtherSevice {
	return &entities.OtherSevice{
		Title:     otherService.Title,
		Url:       otherService.Url,
		Status:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
