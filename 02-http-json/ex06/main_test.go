package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestMux(store *CatalogStore) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /catalog/{sku}/price", updatePriceHandler(store))
	return mux
}

func TestUpdatePriceHandlerUpdatesProductFromPathValue(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPut, "/catalog/HAMMER-001/price", bytes.NewBufferString(`{"price":30}`))
	rec := httptest.NewRecorder()

	newTestMux(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}

	var got Product
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	if got.SKU != "HAMMER-001" || got.Price != 30 {
		t.Fatalf("unexpected product: got %+v", got)
	}
}

func TestUpdatePriceHandlerReturnsNotFoundForUnknownSKU(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPut, "/catalog/MISSING-001/price", bytes.NewBufferString(`{"price":30}`))
	rec := httptest.NewRecorder()

	newTestMux(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusNotFound)
	}
}

func TestUpdatePriceHandlerRejectsInvalidPrice(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPut, "/catalog/HAMMER-001/price", bytes.NewBufferString(`{"price":0}`))
	rec := httptest.NewRecorder()

	newTestMux(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestUpdatePriceHandlerRejectsUnknownField(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPut, "/catalog/HAMMER-001/price", bytes.NewBufferString(`{"price":30,"currency":"USD"}`))
	rec := httptest.NewRecorder()

	newTestMux(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestUpdatePriceHandlerRejectsTrailingJSON(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	req := httptest.NewRequest(http.MethodPut, "/catalog/HAMMER-001/price", bytes.NewBufferString(`{"price":30}{"price":40}`))
	rec := httptest.NewRecorder()

	newTestMux(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestUpdatePriceHandlerRejectsTooLargeBody(t *testing.T) {
	t.Parallel()

	store := NewCatalogStore()
	body := `{"price":30,"note":"` + strings.Repeat("x", 100) + `"}`
	req := httptest.NewRequest(http.MethodPut, "/catalog/HAMMER-001/price", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	newTestMux(store).ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusRequestEntityTooLarge)
	}
}
