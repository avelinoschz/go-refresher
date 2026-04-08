package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewCatalogClientConfiguresTimeout(t *testing.T) {
	t.Parallel()

	c := NewCatalogClient("http://localhost", 5*time.Second)
	if c.client == nil {
		t.Fatal("expected client to be set")
	}
	if c.client.Timeout != 5*time.Second {
		t.Fatalf("unexpected timeout: got %v want %v", c.client.Timeout, 5*time.Second)
	}
	if c.baseURL != "http://localhost" {
		t.Fatalf("unexpected baseURL: got %q want %q", c.baseURL, "http://localhost")
	}
}

func TestGetProductReturnsProduct(t *testing.T) {
	t.Parallel()

	want := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("sku") != "HAMMER-001" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(want)
	}))
	defer srv.Close()

	c := NewCatalogClient(srv.URL, 5*time.Second)
	got, err := c.GetProduct(context.Background(), "HAMMER-001")
	if err != nil {
		t.Fatalf("GetProduct: %v", err)
	}

	if got != want {
		t.Fatalf("unexpected product: got %+v want %+v", got, want)
	}
}

func TestGetProductReturnsErrorOnNotFound(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()

	c := NewCatalogClient(srv.URL, 5*time.Second)
	_, err := c.GetProduct(context.Background(), "MISSING-001")
	if err == nil {
		t.Fatal("expected error for 404 response, got nil")
	}
}

func TestGetProductReturnsErrorOnCanceledContext(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(Product{})
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	c := NewCatalogClient(srv.URL, 5*time.Second)
	_, err := c.GetProduct(ctx, "HAMMER-001")
	if err == nil {
		t.Fatal("expected error for canceled context, got nil")
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

	c := NewCatalogClient(srv.URL, 5*time.Second)
	p := Product{SKU: "DRILL-002", Name: "Drill", Price: 80}
	if err := c.CreateProduct(context.Background(), p); err != nil {
		t.Fatalf("CreateProduct: %v", err)
	}
}

func TestCreateProductSetsContentTypeHeader(t *testing.T) {
	t.Parallel()

	var receivedContentType string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedContentType = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	c := NewCatalogClient(srv.URL, 5*time.Second)
	p := Product{SKU: "DRILL-002", Name: "Drill", Price: 80}
	if err := c.CreateProduct(context.Background(), p); err != nil {
		t.Fatalf("CreateProduct: %v", err)
	}

	if receivedContentType != "application/json" {
		t.Fatalf("unexpected Content-Type: got %q want %q", receivedContentType, "application/json")
	}
}

func TestCreateProductReturnsErrorOnConflict(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "conflict", http.StatusConflict)
	}))
	defer srv.Close()

	c := NewCatalogClient(srv.URL, 5*time.Second)
	p := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}
	if err := c.CreateProduct(context.Background(), p); err == nil {
		t.Fatal("expected error for 409 response, got nil")
	}
}

func TestCreateProductReturnsErrorOnCanceledContext(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	c := NewCatalogClient(srv.URL, 5*time.Second)
	p := Product{SKU: "DRILL-002", Name: "Drill", Price: 80}
	if err := c.CreateProduct(ctx, p); err == nil {
		t.Fatal("expected error for canceled context, got nil")
	}
}
