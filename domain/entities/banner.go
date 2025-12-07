package entities

import "time"

type Banner struct {
	ID            string     `bson:"_id" json:"id"`
	BannerDesktop FileObject `bson:"banner_desktop" json:"banner_desktop"`
	BannerMobile  FileObject `bson:"banner_mobile" json:"banner_mobile"`
	LinkUrl       string     `bson:"link_url" json:"link_url"`
	ViewStatic    int64      `bson:"view_static" json:"view_static"`
	Status        bool       `bson:"status" json:"status"`
	Position      int        `bson:"position" json:"position"`
	CreatedAt     time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `bson:"updated_at" json:"updated_at"`
}

type EditBannerPosition struct {
	ID       string `bson:"_id" json:"id"`
	Position int    `bson:"position" json:"position"`
}
