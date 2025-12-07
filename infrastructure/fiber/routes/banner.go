package routes

import (
	"dtam-fund-cms-backend/infrastructure/fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func BannerRoutes(
	bannerHandler *handler.BannerHandler,
) (app *fiber.App) {

	app = fiber.New()
	app.Post("/", bannerHandler.CreateBanner)
	app.Get("/", bannerHandler.GetBannerList)
	app.Get("/:id", bannerHandler.GetBanner)
	app.Put("/:id", bannerHandler.EditBanner)
	app.Patch("/patch-position", bannerHandler.EditPosition)
	app.Delete("/:id", bannerHandler.DeleteBanner)

	return
}
