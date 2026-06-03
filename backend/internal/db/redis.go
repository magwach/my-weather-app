package db

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis(redisURL string) *redis.Client {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Invalid Redis URL: %v", err)
	}

	RedisClient = redis.NewClient(opt)

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")

	return RedisClient
}

func GetCache(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", redis.Nil
	}
	return val, err
}

func SetCache(key string, value string, ttl time.Duration) error {
	return RedisClient.Set(ctx, key, value, ttl).Err()
}
