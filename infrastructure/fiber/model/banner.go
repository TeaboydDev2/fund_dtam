package model

import (
	"dtam-fund-cms-backend/domain/entities"
	"time"
)

type CreateBanner struct {
	LinkUrl string `form:"link_url"`
}

type EditBanner struct {
	LinkUrl string `form:"link_url"`
	Status  bool   `form:"status"`
}

type EditPosition struct {
	ID       string `json:"id"`
	Position int    `json:"position"`
}

func BannerToEntity(banner *CreateBanner) *entities.Banner {
	timeSet := time.Now()
	return &entities.Banner{
		LinkUrl:    banner.LinkUrl,
		ViewStatic: 0,
		Status:     true,
		CreatedAt:  timeSet,
		UpdatedAt:  timeSet,
	}
}

func EditBannerToEntity(banner *EditBanner) *entities.Banner {
	return &entities.Banner{
		LinkUrl: banner.LinkUrl,
		Status:  banner.Status,
	}
}

func EditPositionBannerToEntity(banner []*EditPosition) []*entities.Banner {

	bannerList := make([]*entities.Banner, len(banner))

	for i, v := range banner {
		bannerList[i] = &entities.Banner{
			ID:       v.ID,
			Position: v.Position,
		}
	}

	return bannerList
}
