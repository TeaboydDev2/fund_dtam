package routes

import (
	"dtam-fund-cms-backend/infrastructure/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func StatViewRoutes(
	handler *handler.StatViewHandler,
) (app *fiber.App) {
	app = fiber.New()
	app.Get("/web-view", handler.GetCountWebView)
	app.Patch("/web-view/increase", handler.IncreaseWebView)
	return
}
