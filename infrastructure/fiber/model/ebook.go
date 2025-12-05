package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"
)

type CreateEBook struct {
	Title string `form:"title"`
}

func NewEBook(eb CreateEBook) *entities.Ebook {
	return &entities.Ebook{
		Title:      eb.Title,
		Number:     0,
		Status:     true,
		ViewStatic: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
