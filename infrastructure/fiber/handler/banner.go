package handler

import (
	"context"
	"dtam-fund-cms-backend/domain/ports"
	fiber_helper "dtam-fund-cms-backend/infrastructure/fiber/helper"
	"dtam-fund-cms-backend/infrastructure/fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BannerHandler struct {
	bannerService ports.BannerService
}

func NewBannerHandler(
	bannerService ports.BannerService,
) *BannerHandler {
	return &BannerHandler{
		bannerService: bannerService,
	}
}

func (bn *BannerHandler) CreateBanner(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	bannerModel := new(model.CreateBanner)

	if err := c.BodyParser(bannerModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	desktop, err := fiber_helper.UploadFileHandler(c, "banner_desktop")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	mobile, err := fiber_helper.UploadFileHandler(c, "banner_mobile")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	banner := model.BannerToEntity(bannerModel)

	if err := bn.bannerService.CreateBanner(ctx, banner, desktop, mobile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "create banner successfully",
	})
}

func (bn *BannerHandler) GetBanner(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 30*time.Second)
	defer cancel()

	id := c.Params("id")

	banner, err := bn.bannerService.GetBanner(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": banner,
	})
}

func (bn *BannerHandler) GetBannerList(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	query := c.Queries()
	page := query["page"]
	limit := query["limit"]

	bannerList, err := bn.bannerService.GetBannerList(ctx, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": bannerList,
	})
}

func (bn *BannerHandler) EditPosition(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	var banners []*model.EditPosition

	if err := c.BodyParser(&banners); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	bannerEntities := model.EditPositionBannerToEntity(banners)

	if err := bn.bannerService.EditPosition(ctx, bannerEntities); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "edit position succesfully",
	})
}

func (bn *BannerHandler) EditBanner(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	bannerModel := new(model.EditBanner)
	id := c.Params("id")

	if err := c.BodyParser(bannerModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	desktop, err := fiber_helper.UploadFileHandler(c, "banner_desktop")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	mobile, err := fiber_helper.UploadFileHandler(c, "banner_mobile")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	banner := model.EditBannerToEntity(bannerModel)

	if err := bn.bannerService.EditBanner(ctx, id, banner, desktop, mobile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "edit banner successfully",
	})
}

func (bn *BannerHandler) DeleteBanner(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	id := c.Params("id")

	if err := bn.bannerService.DeleteBanner(ctx, id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete banner succesfully",
	})
}
