package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"siteAccess/internal/domain"
)

type service interface {
	GetTime(ctx context.Context, site string) (*domain.Time, error)
	GetMinTime(ctx context.Context) (*domain.Site, error)
	GetMaxTime(ctx context.Context) (*domain.Site, error)
}

func NewServer(service service) *fiber.App {
	f := fiber.New()

	f.Use(HandleErrors)

	//f.Get("/swagger/*", swagger.HandlerDefault)
	//
	//f.Use(auth())
	f.Get("/site", getTimeHandler(service))
	f.Get("/min", getMinTimeHandler(service))
	f.Get("/max", getMaxTimeHandler(service))

	return f
}
