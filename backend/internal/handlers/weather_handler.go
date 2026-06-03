package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/magwach/my-weather-app/backend/internal/services"
)

type WeatherHandler struct {
	Service *services.WeatherService
}

func NewWeatherHandler(s *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{Service: s}
}

func (h *WeatherHandler) GetCurrentWeather(c *fiber.Ctx) error {

	city := c.Params("city")

	if city == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "city is required",
		})
	}

	result, err := h.Service.GetCurrentWeather(city)
	if err != nil {

		if err.Error() == "city not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": "city not found",
			})
		}

		if err.Error() == "weather service unavailable" {
			return c.Status(502).JSON(fiber.Map{
				"error": "weather service unavailable",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(result)
}

func (h *WeatherHandler) GetForecast(c *fiber.Ctx) error {

	city := c.Params("city")

	if city == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "city is required",
		})
	}

	result, err := h.Service.GetForecast(city)
	if err != nil {

		if err.Error() == "city not found" {
			return c.Status(404).JSON(fiber.Map{
				"error": "city not found",
			})
		}

		if err.Error() == "weather service unavailable" {
			return c.Status(502).JSON(fiber.Map{
				"error": "weather service unavailable",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(result)
}

func (h *WeatherHandler) GetCurrentWeatherByCoords(c *fiber.Ctx) error {

	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "lat and lon are required",
		})
	}

	result, err := h.Service.GetCurrentWeatherByCoords(lat, lon)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func (h *WeatherHandler) GetForecastByCoords(c *fiber.Ctx) error {

	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "lat and lon are required",
		})
	}

	result, err := h.Service.GetForecastByCoords(lat, lon)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}
