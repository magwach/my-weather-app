package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/magwach/my-weather-app/backend/internal/middlewares"
	"github.com/magwach/my-weather-app/backend/internal/models"
	"github.com/magwach/my-weather-app/backend/internal/services"
)

func TestProtectedRoute_NoToken(t *testing.T) {
	app := fiber.New()

	app.Get(
		"/api/favorites",
		middlewares.AuthMiddleware,
		func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		},
	)

	req := httptest.NewRequest("GET", "/api/favorites", nil)

	res, _ := app.Test(req)

	assert.Equal(t, 401, res.StatusCode)
}

func TestProtectedRoute_InvalidToken(t *testing.T) {
	app := fiber.New()

	app.Get(
		"/api/favorites",
		middlewares.AuthMiddleware,
		func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		},
	)

	req := httptest.NewRequest("GET", "/api/favorites", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")

	res, _ := app.Test(req)

	assert.Equal(t, 401, res.StatusCode)
}

func TestProtectedRoute_ValidToken(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	registerRes, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.NoError(t, err)

	app := fiber.New()

	app.Get(
		"/api/favorites",
		middlewares.AuthMiddleware,
		func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		},
	)

	req := httptest.NewRequest("GET", "/api/favorites", nil)

	req.Header.Set(
		"Authorization",
		"Bearer "+registerRes.Token,
	)

	res, _ := app.Test(req)

	assert.Equal(t, 200, res.StatusCode)
}