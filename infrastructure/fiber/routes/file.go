package routes

import (
	"dtam-fund-cms-backend/infrastructure/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func FileObjectRoutes(
	fileObjectHandler *handler.FileObjectHandler,
) *fiber.App {

	app := fiber.New()

	app.Get("/", fileObjectHandler.Dowload)

	return app
}
