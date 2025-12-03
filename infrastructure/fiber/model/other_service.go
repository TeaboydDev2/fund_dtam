package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"strconv"
	"time"
)

type CreateOtherService struct {
	Title string `form:"title"`
	Url   string `form:"url"`
}

type OtherServiceJSON struct {
	ID         string     `json:"id"`
	Thumbnail  FileObject `json:"thumbnail"`
	Title      string     `json:"title"`
	Url        string     `json:"url"`
	Number     int        `json:"number"`
	Status     string     `json:"status"`
	ViewStatic int64      `json:"view_static"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func OtherServiceToEntity(otherService *CreateOtherService) *entities.OtherSevice {
	return &entities.OtherSevice{
		Title: otherService.Title,
		Url:   otherService.Url,
	}
}

func OtherServiceResponse(otherService *entities.OtherSevice) *OtherServiceJSON {
	return &OtherServiceJSON{
		ID: otherService.ID,
		Thumbnail: FileObject{
			Alt:  otherService.Thumbnail.Alt,
			Ext:  otherService.Thumbnail.Ext,
			Path: otherService.Thumbnail.Path,
		},
		Title:      otherService.Title,
		Url:        otherService.Url,
		Number:     otherService.Number,
		Status:     strconv.FormatBool(otherService.Status),
		ViewStatic: otherService.ViewStatic,
		CreatedAt:  otherService.CreatedAt,
		UpdatedAt:  otherService.UpdatedAt,
	}
}

func OtherServiceResponseList(ots []*entities.OtherSevice) []*OtherServiceJSON {

	serviceList := make([]*OtherServiceJSON, len(ots))

	for i, v := range ots {

		serviceList[i] = &OtherServiceJSON{
			ID: v.ID,
			Thumbnail: FileObject{
				Alt:  v.Thumbnail.Alt,
				Ext:  v.Thumbnail.Ext,
				Path: v.Thumbnail.Path,
			},
			Title:      v.Title,
			Url:        v.Url,
			Number:     v.Number,
			Status:     strconv.FormatBool(v.Status),
			ViewStatic: v.ViewStatic,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
	}

	return serviceList
}
