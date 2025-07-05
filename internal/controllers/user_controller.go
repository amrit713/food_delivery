package controllers

import "github.com/gofiber/fiber/v2"

func GetAllUsers(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Fetched users successfully",
		"data":    []string{"User1", "User2"},
	})
}
