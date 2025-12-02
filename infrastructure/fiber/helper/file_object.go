package fiber_helper

import (
	"dtam-fund-cms-backend/domain/entities"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UploadFileHandler(c *fiber.Ctx, field string) (*entities.FileObject, error) {

	fileHeader, err := c.FormFile(field)
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer file.Close()

	fileExt := strings.ReplaceAll(strings.ToLower(filepath.Ext(fileHeader.Filename)), ".", "")
	fileName := strings.ReplaceAll(fileHeader.Filename, filepath.Ext(fileHeader.Filename), "")
	fileSize := fileHeader.Size
	contentType := fileHeader.Header.Get("Content-Type")

	return &entities.FileObject{
		Alt:         fileName,
		Ext:         fileExt,
		Size:        fileSize,
		ContentType: contentType,
		File:        file,
	}, nil
}
