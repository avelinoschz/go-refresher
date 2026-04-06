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
	cache        *redis.Client
	store        ProductStore
	cacheTimeout time.Duration
}

func NewCatalogService(cache *redis.Client, store ProductStore) CatalogService {
	return CatalogService{
		cache:        cache,
		store:        store,
		cacheTimeout: 100 * time.Millisecond,
	}
}

func (s CatalogService) GetWithFallback(ctx context.Context, sku string) (string, error) {
	cacheCtx, cancel := context.WithTimeout(ctx, s.cacheTimeout)
	defer cancel()

	name, err := s.cache.Get(cacheCtx, "catalog:"+sku).Result()
	if err == nil {
		return name, nil
	}
	if !errors.Is(err, redis.Nil) {
		// Redis error or timeout — fall back to store using the original ctx.
		return s.store.GetName(ctx, sku)
	}

	// Cache miss — go to store and populate cache.
	name, err = s.store.GetName(ctx, sku)
	if err != nil {
		return "", err
	}

	s.cache.Set(ctx, "catalog:"+sku, name, 30*time.Second)
	return name, nil
}

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	service := NewCatalogService(client, StaticStore{})
	ctx := context.Background()

	name, err := service.GetWithFallback(ctx, "HAMMER-001")
	fmt.Println("name:", name, "err:", err)
}
