package fiber_helper

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}
