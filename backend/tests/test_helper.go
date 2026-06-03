package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/magwach/my-weather-app/backend/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	_ = godotenv.Load("../.env")
}

func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	pool, err := db.ConnectDB(os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatalf("failed to connect test db: %v", err)
	}

	redisURL := os.Getenv("TEST_REDIS_URL")
	if redisURL == "" {
		t.Fatal("TEST_REDIS_URL is not set")
	}

	db.ConnectRedis(redisURL)

	_, err = pool.Exec(
		context.Background(),
		"TRUNCATE users, favorites RESTART IDENTITY CASCADE",
	)
	if err != nil {
		t.Fatalf("failed to truncate tables: %v", err)
	}

	return pool
}