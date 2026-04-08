package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchProductReturnsProduct(t *testing.T) {
	t.Parallel()

	want := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Query().Get("sku") != "HAMMER-001" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(want)
	}))
	defer srv.Close()

	got, err := FetchProduct(srv.URL, "HAMMER-001")
	if err != nil {
		t.Fatalf("FetchProduct: %v", err)
	}

	if got != want {
		t.Fatalf("unexpected product: got %+v want %+v", got, want)
	}
}

func TestFetchProductReturnsErrorOnNotFound(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := FetchProduct(srv.URL, "MISSING-001")
	if err == nil {
		t.Fatal("expected error for 404 response, got nil")
	}
}

func TestFetchProductReturnsErrorOnServerError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}))
	defer srv.Close()

	_, err := FetchProduct(srv.URL, "HAMMER-001")
	if err == nil {
		t.Fatal("expected error for 500 response, got nil")
	}
}

func TestCreateProductReturnsNilOnSuccess(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		var p Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	p := Product{SKU: "DRILL-002", Name: "Drill", Price: 80}
	if err := CreateProduct(srv.URL, p); err != nil {
		t.Fatalf("CreateProduct: %v", err)
	}
}

func TestCreateProductReturnsErrorOnConflict(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "conflict", http.StatusConflict)
	}))
	defer srv.Close()

	p := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}
	if err := CreateProduct(srv.URL, p); err == nil {
		t.Fatal("expected error for 409 response, got nil")
	}
}

func TestCreateProductSendsJSONBody(t *testing.T) {
	t.Parallel()

	var received Product

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&received); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	want := Product{SKU: "DRILL-002", Name: "Drill", Price: 80}
	if err := CreateProduct(srv.URL, want); err != nil {
		t.Fatalf("CreateProduct: %v", err)
	}

	if received != want {
		t.Fatalf("unexpected body: got %+v want %+v", received, want)
	}
}
