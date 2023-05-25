package http

import (
	"fmt"
	"net/http"

	_ "siteAccess/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

// GetTime godoc
// @Summary      get the time on the site
// @Description  gets the access time to the transferred site
// @Accept       json
// @Produce      json
// @Param site query string true "Example: yandex.ru"
// @Success      200  {object}  domain.Answer
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /site [get]
func getTimeHandler(service service, counter *prometheus.CounterVec) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		site := ctx.Query("site")
		if site == "" {
			ctx.Status(http.StatusBadRequest)
			return fmt.Errorf("[getTimeHandler] search parameters are not specified")
		}
		resp, err := service.GetTime(ctx.Context(), site)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getTimeHandler] %w", err)
		}
		err = ctx.JSON(resp)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getTimeHandler] failed to return JSON answer, error: %w", err)
		}
		counter.With(prometheus.Labels{"endpoints": "getTimeHandler"}).Inc()

		return nil
	}
}

// GetMin godoc
// @Summary      get name site with minimal time
// @Description  gets the name of the site with the minimum access time
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Site
// @Failure      500  {object}  error
// @Router       /min [get]
func getMinTimeHandler(service service, counter *prometheus.CounterVec) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		resp, err := service.GetMinTime(ctx.Context())
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getMinTimeHandler] %w", err)
		}
		err = ctx.JSON(resp)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getMinTimeHandler] failed to return JSON answer, error: %w", err)
		}
		counter.With(prometheus.Labels{"endpoints": "getMinTimeHandler"}).Inc()

		return nil
	}
}

// GetMax godoc
// @Summary      get name site with maximum time
// @Description  gets the name of the site with the maximum access time
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Site
// @Failure      500  {object}  error
// @Router       /max [get]
func getMaxTimeHandler(service service, counter *prometheus.CounterVec) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		resp, err := service.GetMaxTime(ctx.Context())
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getMaxTimeHandler] %w", err)
		}
		err = ctx.JSON(resp)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return fmt.Errorf("[getMaxTimeHandler] failed to return JSON answer, error: %w", err)
		}
		counter.With(prometheus.Labels{"endpoints": "getMaxTimeHandler"}).Inc()

		return nil
	}
}
