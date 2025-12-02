package routes

import (
	"dtam-fund-cms-backend/infrastructure/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func OtherServiceRoutes(
	otherServiceHandler *handler.OtherServiceHandler,
) *fiber.App {

	app := fiber.New()

	app.Post("/", otherServiceHandler.CreateOtherService)
	app.Get("/", otherServiceHandler.GetOtherServiceList)
	app.Get("/:id/service", otherServiceHandler.GetOtherService)
	app.Put("/:id/edit-service", otherServiceHandler.EditService)
	app.Patch("/", otherServiceHandler.EditStatus)
	app.Delete("/:id/delete", otherServiceHandler.DeleteService)

	return app
}
