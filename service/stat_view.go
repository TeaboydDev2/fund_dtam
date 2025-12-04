package service

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	"time"
)

type statViewService struct {
	statViewRepo ports.StatViewRepository
}

func NewStatViewServiceService(
	statViewRepo ports.StatViewRepository,
) ports.StatViewService {
	return &statViewService{
		statViewRepo: statViewRepo,
	}
}

func (s *statViewService) IncreaseWebView(ctx context.Context) (err error) {
	now := time.Now().Local()
	err = s.statViewRepo.IncreaseWebView(ctx, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	return
}

func (s *statViewService) GetCountWebView(ctx context.Context) (res entities.StatWebView, err error) {
	now := time.Now().Local()
	res, err = s.statViewRepo.QueryWebViewStat(ctx, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	return
}
