package main

import (
	"context"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
)

func newTestClient(t *testing.T) *redis.Client {
	t.Helper()
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{Addr: addr})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Skipf("Redis not available at %s — run 'make redis-up' first: %v", addr, err)
	}
	t.Cleanup(func() { client.Close() })
	return client
}

func TestCatalogServiceReturnsCachedValue(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	ctx := context.Background()

	key := "catalog:HAMMER-001"
	t.Cleanup(func() { client.Del(ctx, key) })

	// Pre-populate the cache with a different name to confirm the cache path is taken.
	client.Set(ctx, key, "CachedHammer", 0)

	cache := NewRedisCache(client)
	service := NewCatalogService(cache, StaticStore{})

	name, err := service.GetProductName(ctx, "HAMMER-001")
	if err != nil {
		t.Fatalf("get product name: %v", err)
	}
	if name != "CachedHammer" {
		t.Fatalf("unexpected name: got %q want %q (cache should have been used)", name, "CachedHammer")
	}
}

func TestCatalogServiceFetchesFromStoreOnMissAndCaches(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	ctx := context.Background()

	key := "catalog:DRILL-002"
	t.Cleanup(func() { client.Del(ctx, key) })

	// Ensure key does not exist.
	client.Del(ctx, key)

	cache := NewRedisCache(client)
	service := NewCatalogService(cache, StaticStore{})

	name, err := service.GetProductName(ctx, "DRILL-002")
	if err != nil {
		t.Fatalf("get product name: %v", err)
	}
	if name != "Drill" {
		t.Fatalf("unexpected name: got %q want %q", name, "Drill")
	}

	// Verify the value was stored in Redis.
	cached, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("expected value to be cached in Redis: %v", err)
	}
	if cached != "Drill" {
		t.Fatalf("unexpected cached value: got %q want %q", cached, "Drill")
	}
}

func TestCatalogServiceReturnsErrorForUnknownSKU(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	ctx := context.Background()

	cache := NewRedisCache(client)
	service := NewCatalogService(cache, StaticStore{})

	if _, err := service.GetProductName(ctx, "UNKNOWN-999"); err == nil {
		t.Fatal("expected error for unknown SKU")
	}
}
