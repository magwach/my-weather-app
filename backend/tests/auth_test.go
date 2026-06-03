package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/magwach/my-weather-app/backend/internal/models"
	"github.com/magwach/my-weather-app/backend/internal/services"
)

func TestRegisterUser_Success(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	res, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Token)
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	_, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.NoError(t, err)

	_, err = auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user already exists")
}

func TestLoginUser_Success(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	_, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.NoError(t, err)

	res, err := auth.LoginUser(models.LoginRequest{
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Token)
}

func TestLoginUser_WrongPassword(t *testing.T) {
	db := setupTestDB(t)

	auth := services.NewAuthService(db)

	_, err := auth.RegisterUser(models.RegisterRequest{
		Name:     "John",
		Email:    "john@test.com",
		Password: "password123",
	})

	assert.NoError(t, err)

	_, err = auth.LoginUser(models.LoginRequest{
		Email:    "john@test.com",
		Password: "wrongpassword",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid credentials")
}