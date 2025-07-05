package routes

import (
	"github.com/amrit713/bank-api/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func userRoute(api fiber.Router) {

	userApi := api.Group("/users")

	userApi.Get("/", controllers.GetAllUsers)
}
