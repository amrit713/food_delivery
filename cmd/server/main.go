package main

import (
	"log"

	"github.com/amrit713/bank-api/config"
	"github.com/amrit713/bank-api/internal/db"
	"github.com/amrit713/bank-api/internal/models"
	"github.com/amrit713/bank-api/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	//fiber instance initialization
	app := fiber.New(fiber.Config{
		AppName: "Food Delivery API",
	})

	db.ConnectDB()
	db.DB.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto;")

	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Restaurant{},
		&models.MenuItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Review{},
		&models.Notification{},
		&models.Session{},
	)

	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	routes.SetupRoutes(app)

	PORT := config.LoadEnv("PORT", ":4000")

	//start server on port
	app.Listen(PORT)
}
