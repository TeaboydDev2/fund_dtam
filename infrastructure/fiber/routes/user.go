package routes

import (
	"dtam-fund-cms-backend/infrastructure/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(
	userHandler *handler.UserHandler,
) *fiber.App {

	app := fiber.New()

	app.Post("/", userHandler.CreateUser)
	app.Get("/:id", userHandler.GetUser)

	return app
}
