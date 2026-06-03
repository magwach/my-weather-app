package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/magwach/my-weather-app/backend/internal/models"
	"github.com/magwach/my-weather-app/backend/internal/services"
)

func TestAddFavorite_LimitEnforced(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	_, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})
	assert.NoError(t, err)

	user, err := auth.GetUserByEmail("john@test.com")
	assert.NoError(t, err)

	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	assert.NotEmpty(t, apiKey, "OPEN_WEATHER_API_KEY is required for tests")

	weather := services.NewWeatherService(apiKey)

	favorites := services.NewFavoritesService(db, weather)

	assert.NoError(t, favorites.AddFavorite(user.ID, "Nairobi"))
	assert.NoError(t, favorites.AddFavorite(user.ID, "Mombasa"))
	assert.NoError(t, favorites.AddFavorite(user.ID, "Kisumu"))

	err = favorites.AddFavorite(user.ID, "Nakuru")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "favorites limit reached")
}

func TestAddFavorite_Duplicate(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	_, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})
	assert.NoError(t, err)

	user, err := auth.GetUserByEmail("john@test.com")
	assert.NoError(t, err)

	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	assert.NotEmpty(t, apiKey, "OPEN_WEATHER_API_KEY is required for tests")

	weather := services.NewWeatherService(apiKey)

	favorites := services.NewFavoritesService(db, weather)

	assert.NoError(t, favorites.AddFavorite(user.ID, "Nairobi"))

	err = favorites.AddFavorite(user.ID, "Nairobi")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "city already in favorites")
}
