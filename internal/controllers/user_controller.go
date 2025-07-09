package controllers

import (
	"github.com/amrit713/food-delivery/internal/dto"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{userService: service}
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {

	users, err := c.userService.GetAllUsers()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return ctx.JSON(fiber.Map{
		"status":     fiber.StatusAccepted,
		"total_user": len(users),
		"data":       users,
	})
}

func (c *UserController) EditMe(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	var input dto.UpdateUserInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	err := c.userService.EditUser(user, &input)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}
