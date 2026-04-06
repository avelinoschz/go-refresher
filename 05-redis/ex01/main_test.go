package main

import (
	"context"
	"os"
	"testing"
	"time"

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

func TestRedisCacheSetAndGet(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	cache := NewRedisCache(client)
	ctx := context.Background()

	key := "test:ex01:set-get"
	t.Cleanup(func() { client.Del(ctx, key) })

	if err := cache.Set(ctx, key, "Hammer", 30*time.Second); err != nil {
		t.Fatalf("set: %v", err)
	}

	value, ok, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if !ok {
		t.Fatal("expected key to be found")
	}
	if value != "Hammer" {
		t.Fatalf("unexpected value: got %q want %q", value, "Hammer")
	}
}

func TestRedisCacheGetMissingKey(t *testing.T) {
	t.Parallel()
	client := newTestClient(t)
	cache := NewRedisCache(client)
	ctx := context.Background()

	_, ok, err := cache.Get(ctx, "test:ex01:missing-key-xyz-abc")
	if err != nil {
		t.Fatalf("unexpected error on missing key: %v", err)
	}
	if ok {
		t.Fatal("expected key to not be found")
	}
}

func TestRedisCacheGetAfterTTLExpiry(t *testing.T) {
	client := newTestClient(t)
	cache := NewRedisCache(client)
	ctx := context.Background()

	key := "test:ex01:ttl"
	t.Cleanup(func() { client.Del(ctx, key) })

	if err := cache.Set(ctx, key, "temporary", 1*time.Second); err != nil {
		t.Fatalf("set: %v", err)
	}

	// Expire the key immediately via Redis to avoid sleeping in tests.
	client.Expire(ctx, key, 0)

	_, ok, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("unexpected error after expiry: %v", err)
	}
	if ok {
		t.Fatal("expected key to be expired")
	}
}
