package fiber_helper

import (
	"dtam-fund-cms-backend/domain/entities"
	"fmt"
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

func UploadMultiFileHandler(c *fiber.Ctx, field string) ([]*entities.FileObject, error) {

	fileHeader, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	multiFile := fileHeader.File[field]

	multiFileObject := make([]*entities.FileObject, len(multiFile))

	for k, v := range multiFile {

		file, err := v.Open()
		if err != nil {
			return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		fileExt := strings.ReplaceAll(strings.ToLower(filepath.Ext(v.Filename)), ".", "")
		fileName := strings.ReplaceAll(v.Filename, filepath.Ext(v.Filename), "")
		fileSize := v.Size
		contentType := v.Header.Get("Content-Type")

		multiFileObject[k] = &entities.FileObject{
			Alt:         fileName,
			Ext:         fileExt,
			Size:        fileSize,
			ContentType: contentType,
			File:        file,
		}
	}

	return multiFileObject, nil
}

// for editor //
func UploadMultiFileEditor(c *fiber.Ctx, field, blobField string) ([]*entities.FileObjectWithBlob, error) {

	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	multiFile := form.File[field]
	blobIDs := form.Value[blobField]

	if len(multiFile) != len(blobIDs) {
		return nil, fmt.Errorf("number of files and blob_ids mismatch")
	}

	multiFileObject := make([]*entities.FileObjectWithBlob, len(multiFile))

	for k, v := range multiFile {

		file, err := v.Open()
		if err != nil {
			return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		fileExt := strings.ReplaceAll(strings.ToLower(filepath.Ext(v.Filename)), ".", "")
		fileName := strings.ReplaceAll(v.Filename, filepath.Ext(v.Filename), "")
		fileSize := v.Size
		contentType := v.Header.Get("Content-Type")

		multiFileObject[k] = &entities.FileObjectWithBlob{
			Alt:         fileName,
			Ext:         fileExt,
			Size:        fileSize,
			ContentType: contentType,
			File:        file,
			BlobID:      blobIDs[k],
		}
	}

	return multiFileObject, nil
}
