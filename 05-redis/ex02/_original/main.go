package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, bool, error)
	Set(ctx context.Context, key, value string, ttl time.Duration) error
}

type ProductStore interface {
	GetName(ctx context.Context, sku string) (string, error)
}

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

type StaticStore struct{}

func (StaticStore) GetName(_ context.Context, sku string) (string, error) {
	products := map[string]string{
		"HAMMER-001": "Hammer",
		"DRILL-002":  "Drill",
	}
	name, ok := products[sku]
	if !ok {
		return "", errors.New("product not found")
	}
	return name, nil
}

type CatalogService struct {
	cache Cache
	store ProductStore
}

func NewCatalogService(cache Cache, store ProductStore) CatalogService {
	return CatalogService{cache: cache, store: store}
}

const cacheTTL = 30 * time.Second

// GetProductName returns the product name for sku.
// It checks the cache first; on a miss it queries the store and caches the result.
func (s CatalogService) GetProductName(ctx context.Context, sku string) (string, error) {
	_ = cacheTTL
	// TODO: implement
	// key format: "catalog:<sku>"
	return "", errors.New("not implemented")
}

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	service := NewCatalogService(NewRedisCache(client), StaticStore{})
	ctx := context.Background()

	name, err := service.GetProductName(ctx, "HAMMER-001")
	fmt.Println("name:", name, "err:", err)
}
