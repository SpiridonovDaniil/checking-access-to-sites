package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func getTimeHandler(service service) func(ctx *fiber.Ctx) error {
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

		return nil
	}
}

func getMinTimeHandler(service service) func(ctx *fiber.Ctx) error {
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

		return nil
	}
}

func getMaxTimeHandler(service service) func(ctx *fiber.Ctx) error {
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

		return nil
	}
}
