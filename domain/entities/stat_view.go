package entities

import "time"

type StatWebView struct {
	ToDay         int64 `json:"to_day" bson:"to_day"`
	Yesterday     int64 `json:"yesterday" bson:"yesterday"`
	ThisMonth     int64 `json:"this_month" bson:"this_month"`
	PreviousMonth int64 `json:"previous_month" bson:"previous_month"`
	ThisYear      int64 `json:"this_year" bson:"this_year"`
	PreviousYear  int64 `json:"previous_year" bson:"previous_year"`
	AllTime       int64 `json:"all_time" bson:"all_time"`
}

type WebView struct {
	Date      time.Time `json:"date" bson:"date"`
	ViewCount int64     `json:"view_count" bson:"-"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"-"`
}
