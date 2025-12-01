package routes

import (
	"fund_dtam/infrastructure/fiber/handler"

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
