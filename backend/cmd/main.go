package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/magwach/my-weather-app/backend/internal/config"
	"github.com/magwach/my-weather-app/backend/internal/db"
)

func main() {
	configs := config.Load()

	_, err := db.ConnectDB(configs.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	db.ConnectRedis(configs.RedisUrl)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     configs.ClientUrl,
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	auth := app.Group("/api/auth")

	auth.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: time.Minute,
	}))

	weather := app.Group("/api/weather")
	forecast := app.Group("/api/forecast")
	favorites := app.Group("/api/favorites")

	_ = auth
	_ = weather
	_ = forecast
	_ = favorites

	log.Printf("Server running on :%s\n", configs.Port)
	log.Fatal(app.Listen(":" + configs.Port))
}
