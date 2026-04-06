package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatalogHandlerReturnsBadRequestWhenSKUMissing(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/catalog", nil)
	rec := httptest.NewRecorder()

	catalogHandler(NewCatalogService()).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCatalogHandlerReturnsNotFoundForUnknownSKU(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/catalog?sku=MISSING-001", nil)
	rec := httptest.NewRecorder()

	catalogHandler(NewCatalogService()).ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusNotFound)
	}
}

func TestCatalogHandlerReturnsProductJSON(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/catalog?sku=HAMMER-001", nil)
	rec := httptest.NewRecorder()

	catalogHandler(NewCatalogService()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}

	var got Product
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	want := Product{
		SKU:   "HAMMER-001",
		Name:  "Hammer",
		Price: 25,
	}

	if got != want {
		t.Fatalf("unexpected product: got %+v want %+v", got, want)
	}
}
