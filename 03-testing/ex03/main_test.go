package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProductHandlerRejectsInvalidJSON(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":`))
	rec := httptest.NewRecorder()

	createProductHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCreateProductHandlerRejectsInvalidPayload(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":"","name":"Hammer","price":25}`))
	rec := httptest.NewRecorder()

	createProductHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCreateProductHandlerCreatesProduct(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":"HAMMER-001","name":"Hammer","price":25}`))
	rec := httptest.NewRecorder()

	createProductHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusCreated)
	}

	if !store.Exists("HAMMER-001") {
		t.Fatal("expected product to be stored")
	}
}

func TestCreateProductHandlerRejectsDuplicateSKU(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	store.Save(Product{
		SKU:   "HAMMER-001",
		Name:  "Hammer",
		Price: 25,
	})

	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":"HAMMER-001","name":"Hammer","price":25}`))
	rec := httptest.NewRecorder()

	createProductHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusConflict)
	}
}
