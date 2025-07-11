package routes

import (
	"github.com/amrit713/food-delivery/internal/controllers"
	"github.com/amrit713/food-delivery/internal/db"
	"github.com/amrit713/food-delivery/internal/middlewares"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/repositories"
	"github.com/amrit713/food-delivery/internal/services"
	"github.com/gofiber/fiber/v2"
)

func userRoute(api fiber.Router) {

	userApi := api.Group("/users")

	userRepo := repositories.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	//user route
	userApi.Get("/", middlewares.Protected(userRepo), middlewares.RoleBase([]models.Role{models.RoleAdmin}),
		userController.GetAllUsers)
	userApi.Put("/me", middlewares.Protected(userRepo), userController.EditMe)

	//auth route
	userApi.Post("/sign-up", authController.Register)
	userApi.Post("/login", authController.Login)
	userApi.Post("/logout", authController.Logout)
}
