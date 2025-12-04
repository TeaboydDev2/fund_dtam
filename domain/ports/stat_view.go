package ports

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"time"
)

type StatViewService interface {
	GetCountWebView(ctx context.Context) (res entities.StatWebView, err error)
	IncreaseWebView(ctx context.Context) (err error)
}

type StatViewRepository interface {
	QueryWebViewStat(ctx context.Context, date time.Time) (res entities.StatWebView, err error)
	IncreaseWebView(ctx context.Context, date time.Time) (err error)
}
