package repository

import (
	"context"
	"siteAccess/internal/domain"
	"time"
)

type Repository interface {
	GetTime(ctx context.Context, site string) (*domain.Time, error)
	GetMinTime(ctx context.Context) (*domain.Site, error)
	GetMaxTime(ctx context.Context) (*domain.Site, error)
	Update(ctx context.Context, newData map[string]time.Duration) error //TODO
}
