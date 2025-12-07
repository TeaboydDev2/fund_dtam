package fiber

import (
	"context"
	"dtam-fund-cms-backend/config"
	"dtam-fund-cms-backend/infrastructure/fiber/handler"
	default_router "dtam-fund-cms-backend/infrastructure/fiber/helper"
	"dtam-fund-cms-backend/infrastructure/fiber/routes"
	"dtam-fund-cms-backend/infrastructure/logger"
	minio_obj "dtam-fund-cms-backend/infrastructure/minio"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"
	"dtam-fund-cms-backend/infrastructure/mongo/repository"
	"dtam-fund-cms-backend/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start(
	ctx context.Context,
	cfg config.Container,
	logger *logger.ZeroLogger,
	mongo *mongodb.MongoClient,
	minio *minio_obj.MinioClient,
) error {

	app := fiber.New(fiber.Config{
		BodyLimit: cfg.HTTP.BodyLimit * 1024 * 1024,
	})

	app.Use(recover.New())
	app.Use(logger.APILogger())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     cfg.HTTP.AllowedOrigin,
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
			AllowCredentials: true,
		},
	))

	app.Use(limiter.New(limiter.Config{
		Max:        2000,
		Expiration: 1 * time.Second,
	}))

	// wired //

	fileRepository := minio_obj.NewMinioRepository(minio)
	otherServiceRepository := repository.NewOtherServiceRepository(mongo)
	statViewRepository := repository.NewStatViewRepository(mongo)
	ebookRepository := repository.NewEBookRepository(mongo)
	bannerRepository := repository.NewBannerRepository(mongo)

	otherService := service.NewOtherService(otherServiceRepository, fileRepository, cfg.Minio, logger)
	fileService := service.NewFileObjectService(fileRepository)
	statViewService := service.NewStatViewServiceService(statViewRepository)
	ebookService := service.NewEBookService(ebookRepository, fileRepository, cfg.Minio)
	bannerService := service.NewBannerService(bannerRepository, fileRepository, logger, cfg.Minio)

	otherServiceHandler := handler.NewOtherServiceHandler(logger, otherService, fileService)
	fileHandler := handler.NewFileObjectHandler(fileService)
	statViewHandler := handler.NewStatViewHandler(statViewService)
	ebookHandler := handler.NewEBookHandler(ebookService)
	bannerHandler := handler.NewBannerHandler(bannerService)

	dtam := app.Group(fmt.Sprintf("/dtam-fund/%s", cfg.HTTP.Prefix))
	dtam.Mount("/other-service", routes.OtherServiceRoutes(otherServiceHandler))
	dtam.Mount("/file", routes.FileObjectRoutes(fileHandler))
	dtam.Mount("/stat", routes.StatViewRoutes(statViewHandler))
	dtam.Mount("/ebook", routes.EBookRoutes(ebookHandler))
	dtam.Mount("/banner", routes.BannerRoutes(bannerHandler))

	app.Get("/health-check", default_router.HealthCheck)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})

	go func() {
		if err := app.Listen(fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)); err != nil {
			log.Printf("Server closed: %v", err)
		}
	}()

	log.Printf("Server running on port %s", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Graceful shutdown...")

	return app.ShutdownWithContext(ctx)
}
