package handler

import (
	"context"
	"dtam-fund-cms-backend/domain/ports"
	fiberHelper "dtam-fund-cms-backend/infrastructure/fiber/helper"
	"dtam-fund-cms-backend/infrastructure/fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OtherServiceHandler struct {
	otherService ports.OtherSevice
	fileService  ports.FileStorageService
}

func NewOtherServiceHandler(
	otherService ports.OtherSevice,
	fileService ports.FileStorageService,
) *OtherServiceHandler {
	return &OtherServiceHandler{
		otherService: otherService,
		fileService:  fileService,
	}
}

func (ots *OtherServiceHandler) CreateOtherService(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	otherService := new(model.CreateOtherService)

	if err := c.BodyParser(otherService); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	file, err := fiberHelper.UploadFileHandler(c, "other_service_thumbnail")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	otherServiceEntities := model.OtherServiceToEntity(otherService)

	if err := ots.otherService.CreateService(ctx, otherServiceEntities, file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "service create successfully",
	})
}

func (ots *OtherServiceHandler) GetOtherService(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	id := c.Params("id")

	service, err := ots.otherService.GetService(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    service,
		"message": "service fetch successfully",
	})
}

func (ots *OtherServiceHandler) GetOtherServiceList(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	query := c.Queries()
	page := query["page"]
	limit := query["limit"]

	service, err := ots.otherService.GetServiceList(ctx, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    service,
		"message": "service fetch successfully",
	})
}

func (ots *OtherServiceHandler) EditService(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	id := c.Params("id")

	otherService := new(model.CreateOtherService)

	if err := c.BodyParser(otherService); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	file, err := fiberHelper.UploadFileHandler(c, "other_service_thumbnail")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	otherServiceEntities := model.OtherServiceToEntity(otherService)

	if err := ots.otherService.EditService(ctx, id, otherServiceEntities, file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "edit service successfully",
	})
}

func (ots *OtherServiceHandler) EditStatus(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	id := c.Query("id")
	status := c.QueryBool("status", false)

	if err := ots.otherService.EditStatus(ctx, id, status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "edit service successfully",
	})
}

func (ots *OtherServiceHandler) DeleteService(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	id := c.Params("id")

	if err := ots.otherService.DeleteService(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete service successfully",
	})
}
