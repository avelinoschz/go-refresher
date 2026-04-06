package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

// Get returns the value for key. Returns (value, true, nil) on hit,
// ("", false, nil) when the key does not exist, and ("", false, err) on error.
func (c *RedisCache) Get(ctx context.Context, key string) (string, bool, error) {
	// TODO: implement
	// hint: errors.Is(err, redis.Nil) means the key is missing, not a failure
	return "", false, nil
}

// Set stores value under key with the given TTL.
func (c *RedisCache) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	// TODO: implement
	return nil
}

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	cache := NewRedisCache(client)
	ctx := context.Background()

	_ = cache.Set(ctx, "catalog:HAMMER-001", "Hammer", 30*time.Second)
	value, ok, err := cache.Get(ctx, "catalog:HAMMER-001")
	fmt.Println("value:", value, "found:", ok, "err:", err)
}
