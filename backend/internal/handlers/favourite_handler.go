package handlers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/magwach/my-weather-app/backend/internal/services"
)

type FavoritesHandler struct {
	Service *services.FavoritesService
}

func NewFavoritesHandler(s *services.FavoritesService) *FavoritesHandler {
	return &FavoritesHandler{Service: s}
}

func (h *FavoritesHandler) AddFavorite(c *fiber.Ctx) error {

	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var body struct {
		City string `json:"city"`
	}

	if err := c.BodyParser(&body); err != nil || body.City == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.Service.AddFavorite(userID, body.City)
	if err != nil {

		switch err.Error() {

		case "favorites limit reached, maximum 3 cities allowed":
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})

		case "city not found":
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})

		case "city already in favorites":
			return c.Status(409).JSON(fiber.Map{
				"error": err.Error(),
			})

		default:
			log.Println(err)
			return c.Status(500).JSON(fiber.Map{
				"error": "internal server error",
			})
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "favorite added",
	})
}

func (h *FavoritesHandler) GetFavorites(c *fiber.Ctx) error {

	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	result, err := h.Service.GetFavorites(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch favorites",
		})
	}

	return c.JSON(result)
}

func (h *FavoritesHandler) RemoveFavorite(c *fiber.Ctx) error {

	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	city := c.Params("city")
	if city == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "city is required",
		})
	}

	err := h.Service.RemoveFavorite(userID, city)
	if err != nil {

		if errors.Is(err, errors.New("city not found in favorites")) {
			return c.Status(404).JSON(fiber.Map{
				"error": "city not found in favorites",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "failed to remove favorite",
		})
	}

	return c.JSON(fiber.Map{
		"message": "favorite removed",
	})
}
