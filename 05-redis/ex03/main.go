package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductStore interface {
	GetName(ctx context.Context, sku string) (string, error)
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
	cache  *redis.Client
	store  ProductStore
	cacheTimeout time.Duration
}

func NewCatalogService(cache *redis.Client, store ProductStore) CatalogService {
	return CatalogService{
		cache:        cache,
		store:        store,
		cacheTimeout: 100 * time.Millisecond,
	}
}

// GetWithFallback tries to read from Redis using a short timeout.
// If Redis is unavailable or too slow, it falls back to the store.
func (s CatalogService) GetWithFallback(ctx context.Context, sku string) (string, error) {
	_ = s.cacheTimeout
	// TODO: implement
	// 1. create a child context with s.cacheTimeout deadline
	// 2. try s.cache.Get with the child context
	// 3. on success (including redis.Nil miss), use the result or fall through
	// 4. on any error (including DeadlineExceeded), fall back to s.store.GetName(ctx, sku)
	return "", errors.New("not implemented")
}

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	service := NewCatalogService(client, StaticStore{})
	ctx := context.Background()

	name, err := service.GetWithFallback(ctx, "HAMMER-001")
	fmt.Println("name:", name, "err:", err)
}
