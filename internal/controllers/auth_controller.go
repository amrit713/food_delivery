package controllers

import (
	"log"
	"time"

	"github.com/amrit713/food-delivery/internal/dto"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var input models.User

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	user, err := c.authService.Register(&input)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// set cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt_token",
		Value:    user.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		Path:     "/",
	})

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"code":   fiber.StatusOK, "data": dto.AuthResponse{User: user.User, Token: user.Token},
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var input dto.LoginInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}

	log.Println("inputs", input)

	user, err := c.authService.Login(&input)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// set cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt_token",
		Value:    user.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		Path:     "/",
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"code":   fiber.StatusOK, "data": dto.AuthResponse{User: user.User, Token: user.Token},
	})
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt_token",
		Value:    "",                             // Clear the value
		Expires:  time.Now().Add(-1 * time.Hour), // Expire the cookie
		MaxAge:   -1,                             // Expire immediately
		Path:     "/",                            // Match original cookie path
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout sucessfully",
	})
}
