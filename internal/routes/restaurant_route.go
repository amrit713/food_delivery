package routes

import (
	"github.com/amrit713/food-delivery/internal/controllers"
	"github.com/amrit713/food-delivery/internal/db"
	"github.com/amrit713/food-delivery/internal/middlewares"
	"github.com/amrit713/food-delivery/internal/repositories"
	"github.com/amrit713/food-delivery/internal/services"
	"github.com/gofiber/fiber/v2"
)

func restaurantRoute(api fiber.Router) {
	restaurantApi := api.Group("/restaurants")

	userRepo := repositories.NewUserRepository(db.DB)
	restaurantRepo := repositories.NewRestaurantRepository(db.DB)
	restaurantService := services.NewRestaurantService(restaurantRepo)
	restaurantController := controllers.NewRestaurantController(restaurantService)

	restaurantApi.Post("/", middlewares.Protected(userRepo), restaurantController.Create)
	restaurantApi.Get("/", restaurantController.GetAllResturant)
	restaurantApi.Put("/:id", middlewares.Protected(userRepo), restaurantController.EditResturant)
	restaurantApi.Delete("/:id", middlewares.Protected(userRepo), restaurantController.Delete)

}
