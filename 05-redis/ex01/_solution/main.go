package main

import (
	"context"
	"errors"
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

func (c *RedisCache) Get(ctx context.Context, key string) (string, bool, error) {
	value, err := c.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return value, true, nil
}

func (c *RedisCache) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	cache := NewRedisCache(client)
	ctx := context.Background()

	_ = cache.Set(ctx, "catalog:HAMMER-001", "Hammer", 30*time.Second)
	value, ok, err := cache.Get(ctx, "catalog:HAMMER-001")
	fmt.Println("value:", value, "found:", ok, "err:", err)
}
