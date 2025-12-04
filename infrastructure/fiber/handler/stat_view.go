package handler

import (
	"context"
	"dtam-fund-cms-backend/domain/ports"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StatViewHandler struct {
	statViewSrv ports.StatViewService
}

func NewStatViewHandler(
	statViewSrv ports.StatViewService,
) *StatViewHandler {
	return &StatViewHandler{
		statViewSrv: statViewSrv,
	}
}

func (h *StatViewHandler) IncreaseWebView(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	err := h.statViewSrv.IncreaseWebView(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Increase Stat WebView successfully",
	})
}

func (h *StatViewHandler) GetCountWebView(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	res, err := h.statViewSrv.GetCountWebView(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    res,
		"message": "Get Count Stat WebView successfully",
	})
}
