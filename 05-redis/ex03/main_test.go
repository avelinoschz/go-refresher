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

func TestGetWithFallbackReturnsCachedValue(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	ctx := context.Background()

	key := "catalog:HAMMER-001"
	t.Cleanup(func() { client.Del(ctx, key) })

	client.Set(ctx, key, "CachedHammer", 0)

	service := NewCatalogService(client, StaticStore{})
	name, err := service.GetWithFallback(ctx, "HAMMER-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "CachedHammer" {
		t.Fatalf("unexpected name: got %q want %q", name, "CachedHammer")
	}
}

func TestGetWithFallbackUsesStoreOnCacheMiss(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	ctx := context.Background()

	key := "catalog:DRILL-002"
	t.Cleanup(func() { client.Del(ctx, key) })

	client.Del(ctx, key)

	service := NewCatalogService(client, StaticStore{})
	name, err := service.GetWithFallback(ctx, "DRILL-002")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "Drill" {
		t.Fatalf("unexpected name: got %q want %q", name, "Drill")
	}
}

func TestGetWithFallbackUsesStoreWhenContextCancelled(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)

	// Use an already-cancelled context to simulate Redis being unreachable.
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	service := NewCatalogService(client, StaticStore{})

	// The fallback store should still work using the original (cancelled) ctx
	// is passed to it — but our store ignores ctx, so it should succeed.
	name, err := service.GetWithFallback(ctx, "HAMMER-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "Hammer" {
		t.Fatalf("unexpected name: got %q want %q", name, "Hammer")
	}
}
