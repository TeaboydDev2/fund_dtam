package handler

import (
	"context"
	"dtam-fund-cms-backend/domain/ports"
	fiber_helper "dtam-fund-cms-backend/infrastructure/fiber/helper"
	"dtam-fund-cms-backend/infrastructure/fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(
	userService ports.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (us *UserHandler) CreateUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Minute)
	defer cancel()

	newUser := new(model.CreateUser)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	file, err := fiber_helper.UploadFileHandler(c, "profile_picture")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	user := model.ToEntity(newUser)

	if err := us.userService.CreateUser(ctx, user, file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user create successfully",
	})
}

func (us *UserHandler) GetUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), 15*time.Second)
	defer cancel()

	id := c.Params("id")

	user, picProfile, err := us.userService.GetUser(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": model.ToResponse(user, picProfile),
	})
}
