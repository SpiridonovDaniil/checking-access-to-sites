package http

import (
	"context"

	_ "siteAccess/docs"
	"siteAccess/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/prometheus/client_golang/prometheus"
)

//go:generate mockgen -source=server.go -destination=mocks/mock.go

type service interface {
	GetTime(ctx context.Context, site string) (*domain.Answer, error)
	GetMinTime(ctx context.Context) (*domain.Site, error)
	GetMaxTime(ctx context.Context) (*domain.Site, error)
}

func NewServer(service service, counter *prometheus.CounterVec) *fiber.App {
	f := fiber.New()

	f.Use(HandleErrors)

	f.Get("/swagger/*", swagger.HandlerDefault)

	f.Get("/site", getTimeHandler(service, counter))
	f.Get("/min", getMinTimeHandler(service, counter))
	f.Get("/max", getMaxTimeHandler(service, counter))

	return f
}
