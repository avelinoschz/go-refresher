package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMemoryCacheSetAndGet(t *testing.T) {
	t.Parallel()

	cache := NewMemoryCache()
	cache.Set("catalog:HAMMER-001", "Hammer")

	got, ok := cache.Get("catalog:HAMMER-001")
	if !ok {
		t.Fatal("expected key to be found")
	}

	if got != "Hammer" {
		t.Fatalf("unexpected value: got %q want %q", got, "Hammer")
	}
}

func TestMemoryCacheGetReturnsFalseForMissingKey(t *testing.T) {
	t.Parallel()

	cache := NewMemoryCache()

	if _, ok := cache.Get("missing"); ok {
		t.Fatal("expected missing key lookup to return false")
	}
}

func TestMemoryCacheSupportsConcurrentAccess(t *testing.T) {
	t.Parallel()

	cache := NewMemoryCache()
	var wg sync.WaitGroup

	for i := 0; i < 25; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			key := fmt.Sprintf("key-%d", i)
			cache.Set(key, "value")
			if got, ok := cache.Get(key); !ok || got != "value" {
				t.Errorf("unexpected cache lookup for %s: got %q found=%v", key, got, ok)
			}
		}()
	}

	wg.Wait()
}
