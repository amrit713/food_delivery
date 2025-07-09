package main

import (
	"log"

	"github.com/amrit713/food-delivery/config"
	"github.com/amrit713/food-delivery/internal/db"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/routes"

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

	extQuery := `CREATE EXTENSION IF NOT EXISTS pgcrypto;`
	if err := db.DB.Exec(extQuery).Error; err != nil {
		log.Fatal("failed to create pgcrypto extension:", err)
	}

	// ✅ Create ENUM type for user roles
	enumQuery := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
			CREATE TYPE user_role AS ENUM ('admin', 'user', 'rider');
		END IF;
	END $$;
	`
	if err := db.DB.Exec(enumQuery).Error; err != nil {
		log.Fatal("failed to create enum:", err)
	}

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
