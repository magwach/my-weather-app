package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/magwach/my-weather-app/backend/internal/config"
)

func AuthMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "missing authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == authHeader {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid token format",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid or expired token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)

	c.Locals("user_id", userID)

	return c.Next()
}
