package service

import (
	"context"
	"fmt"
	"siteAccess/internal/domain"
	"siteAccess/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTime(ctx context.Context, site string) (*domain.Time, error) {
	a, err := s.repo.GetTime(ctx, site)
	if err != nil {
		return nil, fmt.Errorf("[getTime] error in obtaining data about access to the site %s, error: %w", site, err)
	}
	return a, err
}

func (s *Service) GetMinTime(ctx context.Context) (*domain.Site, error) {
	a, err := s.repo.GetMinTime(ctx)
	if err != nil {
		return nil, fmt.Errorf("[getMinTime] error getting the site name with minimal access time, error: %w", err)
	}
	return a, err
}

func (s *Service) GetMaxTime(ctx context.Context) (*domain.Site, error) {
	a, err := s.repo.GetMaxTime(ctx)
	if err != nil {
		return nil, fmt.Errorf("[getMaxTime] error getting the site name with maximum access time, error: %w", err)
	}
	return a, err
}
