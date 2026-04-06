package main

import (
	"errors"
	"testing"
)

func TestCatalogServiceReturnsValueFromCache(t *testing.T) {
	t.Parallel()

	cache := NewMemoryCache()
	cache.Set("catalog:HAMMER-001", "Hammer")

	service := NewCatalogService(cache, stubFetcher{})

	got, err := service.GetProductName("HAMMER-001")
	if err != nil {
		t.Fatalf("get product name: %v", err)
	}

	if got != "Hammer" {
		t.Fatalf("unexpected product name: got %q want %q", got, "Hammer")
	}
}

func TestCatalogServiceFetchesAndCachesOnMiss(t *testing.T) {
	t.Parallel()

	cache := NewMemoryCache()
	fetcher := stubFetcher{name: "Hammer"}
	service := NewCatalogService(cache, fetcher)

	got, err := service.GetProductName("HAMMER-001")
	if err != nil {
		t.Fatalf("get product name: %v", err)
	}

	if got != "Hammer" {
		t.Fatalf("unexpected product name: got %q want %q", got, "Hammer")
	}

	cachedValue, ok := cache.Get("catalog:HAMMER-001")
	if !ok {
		t.Fatal("expected fetched value to be cached")
	}

	if cachedValue != "Hammer" {
		t.Fatalf("unexpected cached value: got %q want %q", cachedValue, "Hammer")
	}
}

func TestCatalogServiceReturnsFetcherError(t *testing.T) {
	t.Parallel()

	service := NewCatalogService(NewMemoryCache(), stubFetcher{err: errors.New("product not found")})

	if _, err := service.GetProductName("MISSING-001"); err == nil {
		t.Fatal("expected fetcher error")
	}
}

type stubFetcher struct {
	name string
	err  error
}

func (s stubFetcher) FetchName(string) (string, error) {
	if s.err != nil {
		return "", s.err
	}

	return s.name, nil
}
