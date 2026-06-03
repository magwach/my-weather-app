package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/magwach/my-weather-app/backend/internal/models"
	"github.com/magwach/my-weather-app/backend/internal/services"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}


func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	res, err := h.Service.RegisterUser(req)
	if err != nil {
		if err == services.ErrUserExists {
			return c.Status(409).JSON(fiber.Map{
				"error": "email already exists",
			})
		}

		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(res)
}


func (h *AuthHandler) Login(c *fiber.Ctx) error {

	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	res, err := h.Service.LoginUser(req)
	if err != nil {
		if err == services.ErrInvalidCreds {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(res)
}