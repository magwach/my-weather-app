package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl       string
	RedisUrl          string
	JwtSecret         string
	OpenWeatherApiKey string
	Port              string
	ClientUrl         string
}

func Load() *Config {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	ClientUrl := os.Getenv("CLIENT_URL")

	if ClientUrl == "" {
		log.Fatal("CLIENT_URL is not set")
	}

	RedisUrl := os.Getenv("REDIS_URL")

	if RedisUrl == "" {
		log.Fatal("REDIS_URL is not set")
	}

	DatabaseUrl := os.Getenv("DATABASE_URL")

	if DatabaseUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	JwtSecret := os.Getenv("JWT_SECRET")

	if JwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	OpenWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")

	if OpenWeatherApiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY is not set")
	}

	Port := os.Getenv("PORT")

	if Port == "" {
		log.Fatal("PORT is not set")
	}

	return &Config{
		DatabaseUrl:       DatabaseUrl,
		RedisUrl:          RedisUrl,
		JwtSecret:         JwtSecret,
		OpenWeatherApiKey: OpenWeatherApiKey,
		Port:              Port,
		ClientUrl:         ClientUrl,
	}
}
