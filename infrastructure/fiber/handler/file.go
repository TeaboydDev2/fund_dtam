package handler

import (
	"context"
	"dtam-fund-cms-backend/domain/ports"
	fiber_helper "dtam-fund-cms-backend/infrastructure/fiber/helper"
	"io"

	"github.com/gofiber/fiber/v2"
)

type FileObjectHandler struct {
	fileService ports.FileStorageService
}

func NewFileObjectHandler(
	fileService ports.FileStorageService,
) *FileObjectHandler {
	return &FileObjectHandler{
		fileService: fileService,
	}
}

// test//
func (ots *FileObjectHandler) Dowload(c *fiber.Ctx) error {

	ctx := context.Background()

	file, err := ots.fileService.Dowload(ctx, "service_thumbnail/34a9f8fc-612c-471a-8afa-2552087d951d")
	if err != nil {
		return err
	}
	defer file.Close()

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", `attachment; filename="yourfile.pdf"`)

	_, err = io.Copy(c, file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

// for editor test //
func (ots *FileObjectHandler) UploadPlaygroud(c *fiber.Ctx) error {

	_, err := fiber_helper.UploadMultiFileEditor(c, "images", "blob_id")
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
