package controllers

import (
	"github.com/amrit713/food-delivery/internal/dto"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/services"
	"github.com/gofiber/fiber/v2"
)

type RestaurantController struct {
	restaurantService *services.RestaurantService
}

func NewRestaurantController(s *services.RestaurantService) *RestaurantController {
	return &RestaurantController{restaurantService: s}
}

func (c *RestaurantController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)

	var input dto.ResturantInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	restaurnat, err := c.restaurantService.Create(&input, user)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Restaurant created successfully",
		"data":    restaurnat,
	})

}

func (c *RestaurantController) GetAllResturant(ctx *fiber.Ctx) error {
	resturants, err := c.restaurantService.GetAllResturant()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return ctx.JSON(fiber.Map{
		"data":              resturants,
		"status":            fiber.StatusOK,
		"total_restaurants": len(resturants),
	})
}

func (c *RestaurantController) EditResturant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var input dto.ResturantInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user := ctx.Locals("user").(*models.User)

	updated, err := c.restaurantService.UpdateRestaurant(id, &input, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Resturant updated successfully",
		"restaurant": updated,
	})

}

func (c *RestaurantController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := ctx.Locals("user").(*models.User)

	err := c.restaurantService.DeleteRestaurant(id, user.ID)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Restaurant deleted successfully",
	})
}
