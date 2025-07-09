package middlewares

import (
	"slices"

	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/repositories"
	"github.com/amrit713/food-delivery/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected(userRepo *repositories.UserRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		jwtTokenString := ctx.Cookies("jwt_token")

		if jwtTokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"error":      "unauthorized",
			})
		}

		token, err := jwt.Parse(jwtTokenString, func(t *jwt.Token) (any, error) {
			return utils.JwtSecrect, nil
		})

		if err != nil || !token.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"error":      "unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Failed to parse claims",
			})
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
		}

		user, err := userRepo.GetUserByID(userID)

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}

func RoleBase(allowedRoles []models.Role) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*models.User)

		if slices.Contains(allowedRoles, user.Role) {
			return ctx.Next()
		}

		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}
}
