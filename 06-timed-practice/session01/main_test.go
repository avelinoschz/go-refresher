package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatalogHandlerMissingSKU(t *testing.T) {
	t.Parallel()

	service := NewCatalogService(NewMemoryCache(), NewStaticStore())
	handler := catalogHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/catalog", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCatalogHandlerUnknownSKU(t *testing.T) {
	t.Parallel()

	service := NewCatalogService(NewMemoryCache(), NewStaticStore())
	handler := catalogHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/catalog?sku=UNKNOWN-999", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusNotFound)
	}
}

func TestCatalogHandlerReturnsProductJSON(t *testing.T) {
	t.Parallel()

	service := NewCatalogService(NewMemoryCache(), NewStaticStore())
	handler := catalogHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/catalog?sku=HAMMER-001", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}

	var product Product
	if err := json.NewDecoder(rec.Body).Decode(&product); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	if product.SKU != "HAMMER-001" {
		t.Fatalf("unexpected sku: got %q want %q", product.SKU, "HAMMER-001")
	}
	if product.Price != 25 {
		t.Fatalf("unexpected price: got %d want %d", product.Price, 25)
	}
}

func TestCatalogHandlerSecondRequestHitsCache(t *testing.T) {
	t.Parallel()

	cache := NewMemoryCache()
	service := NewCatalogService(cache, NewStaticStore())
	handler := catalogHandler(service)

	for i := range 2 {
		req := httptest.NewRequest(http.MethodGet, "/catalog?sku=DRILL-002", nil)
		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("request %d: unexpected status code: got %d want %d", i+1, rec.Code, http.StatusOK)
		}
	}

	// Verify the product was cached after the first request.
	if _, ok := cache.Get("DRILL-002"); !ok {
		t.Fatal("expected product to be cached after first request")
	}
}
