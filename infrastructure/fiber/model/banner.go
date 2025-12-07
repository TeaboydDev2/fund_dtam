package model

type CreateBanner struct {
	Title   string `form:"title"`
	LinkUrl string `form:"link_url"`
}
