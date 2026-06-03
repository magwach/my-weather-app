package main

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/magwach/my-weather-app/backend/internal/config"
	"github.com/magwach/my-weather-app/backend/internal/db"
	"github.com/magwach/my-weather-app/backend/internal/handlers"
	"github.com/magwach/my-weather-app/backend/internal/middlewares"
	"github.com/magwach/my-weather-app/backend/internal/services"
)

func main() {

	cfg := config.Load()

	pg, err := db.ConnectDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	db.ConnectRedis(cfg.RedisUrl)

	authService := services.NewAuthService(pg)
	weatherService := services.NewWeatherService(cfg.OpenWeatherApiKey)
	favoritesService := services.NewFavoritesService(pg, weatherService)

	authHandler := handlers.NewAuthHandler(authService)
	weatherHandler := handlers.NewWeatherHandler(weatherService)
	favoritesHandler := handlers.NewFavoritesHandler(favoritesService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "internal server error"

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			return c.Status(code).JSON(fiber.Map{
				"error": message,
			})
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.ClientUrl,
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	api := app.Group("/api")

	auth := api.Group("/auth")

	auth.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: time.Minute,
	}))

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	protected := api.Use(middlewares.AuthMiddleware)

	protected.Get("/weather/:city", weatherHandler.GetCurrentWeather)
	protected.Get("/forecast/:city", weatherHandler.GetForecast)

	protected.Get("/favorites", favoritesHandler.GetFavorites)
	protected.Post("/favorites", favoritesHandler.AddFavorite)
	protected.Delete("/favorites/:city", favoritesHandler.RemoveFavorite)

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}