package routes

import (
	"dtam-fund-cms-backend/infrastructure/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func EBookRoutes(
	EBookHandler *handler.EBookHandler,
) (app *fiber.App) {

	app = fiber.New()
	app.Post("/", EBookHandler.CreateEBook)
	app.Get("/", EBookHandler.GetEBookList)
	app.Get("/:id", EBookHandler.GetEBook)
	app.Put("/:id", EBookHandler.EditEBook)
	app.Delete("/:id", EBookHandler.DeleteEBook)

	return
}
