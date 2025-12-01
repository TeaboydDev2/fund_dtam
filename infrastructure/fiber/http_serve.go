package fiber

import (
	"context"
	"fmt"
	"fund_dtam/config"
	"fund_dtam/infrastructure/fiber/handler"
	default_router "fund_dtam/infrastructure/fiber/helper"
	"fund_dtam/infrastructure/fiber/routes"
	mongodb "fund_dtam/infrastructure/mongo"
	"fund_dtam/infrastructure/mongo/repository"
	"fund_dtam/service"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start(ctx context.Context, cfg config.HTTP, mongo *mongodb.MongoClient) error {

	app := fiber.New(fiber.Config{
		BodyLimit: cfg.BodyLimit * 1024 * 1024,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     cfg.AllowedOrigin,
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
			AllowCredentials: true,
		},
	))
	app.Use(limiter.New(limiter.Config{
		Max:        2000,
		Expiration: 1 * time.Second,
	}))

	userRepository := repository.NewUserRepository(mongo)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	app.Mount("/users", routes.UserRoutes(userHandler))

	app.Get("/healt-check", default_router.HealthCheck)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})

	go func() {
		if err := app.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
			log.Printf("Server closed: %v", err)
		}
	}()

	log.Printf("Server running on port %s", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Graceful shutdown...")

	return app.ShutdownWithContext(ctx)
}
