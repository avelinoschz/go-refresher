package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatalogStoreSavePersistsProduct(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	product := Product{
		SKU:   "DRILL-002",
		Name:  "Drill",
		Price: 80,
	}

	if err := store.Save(product); err != nil {
		t.Fatalf("save product: %v", err)
	}

	got, ok := store.products["DRILL-002"]
	if !ok {
		t.Fatal("expected product to be stored")
	}

	if got != product {
		t.Fatalf("unexpected product: got %+v want %+v", got, product)
	}
}

func TestCreateCatalogHandlerRejectsInvalidPayload(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":"","name":"Drill","price":80}`))
	rec := httptest.NewRecorder()

	createCatalogHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestCreateCatalogHandlerRejectsDuplicateSKU(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":"HAMMER-001","name":"Hammer","price":25}`))
	rec := httptest.NewRecorder()

	createCatalogHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusConflict)
	}
}

func TestCreateCatalogHandlerCreatesProduct(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPost, "/catalog", bytes.NewBufferString(`{"sku":"DRILL-002","name":"Drill","price":80}`))
	rec := httptest.NewRecorder()

	createCatalogHandler(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusCreated)
	}

	if _, ok := store.products["DRILL-002"]; !ok {
		t.Fatal("expected product to be stored")
	}
}
