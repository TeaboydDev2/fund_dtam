package handler

import (
	"context"
	"dtam-fund-cms-backend/domain/ports"
	fiber_helper "dtam-fund-cms-backend/infrastructure/fiber/helper"
	"dtam-fund-cms-backend/infrastructure/fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type EBookHandler struct {
	ebookService ports.EBookService
}

func NewEBookHandler(
	ebookService ports.EBookService,
) *EBookHandler {
	return &EBookHandler{
		ebookService: ebookService,
	}
}

func (eb *EBookHandler) CreateEBook(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	title := new(model.CreateEBook)

	if err := c.BodyParser(title); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	thumbnail, err := fiber_helper.UploadFileHandler(c, "thumbnail")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ebookFile, err := fiber_helper.UploadFileHandler(c, "ebook_file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ebook := model.NewEBook(*title)

	if err := eb.ebookService.CreateEBook(ctx, ebook, thumbnail, ebookFile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "create ebook successfully",
	})
}

func (eb *EBookHandler) GetEBook(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	id := c.Params("id")

	ebook, err := eb.ebookService.GetEBook(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    ebook,
		"message": "ebook fetch successfully",
	})
}

func (eb *EBookHandler) GetEBookList(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	q := c.Queries()
	page := q["page"]
	limit := q["limit"]

	ebookList, err := eb.ebookService.GetEBookList(ctx, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    ebookList,
		"message": "ebook list fetch successfully",
	})
}

func (eb *EBookHandler) EditEBook(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	id := c.Params("id")
	title := new(model.CreateEBook)

	if err := c.BodyParser(title); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	thumbnail, err := fiber_helper.UploadFileHandler(c, "thumbnail")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ebookFile, err := fiber_helper.UploadFileHandler(c, "ebook_file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ebook := model.NewEBook(*title)

	if err := eb.ebookService.EditEBook(ctx, id, ebook, thumbnail, ebookFile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "edit ebook successfully",
	})
}

func (eb *EBookHandler) DeleteEBook(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Minute)
	defer cancel()

	id := c.Params("id")

	if err := eb.ebookService.DeleteEBook(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete ebook successfully",
	})
}
