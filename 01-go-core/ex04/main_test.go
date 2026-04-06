package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestNotFoundError_Error(t *testing.T) {
	t.Parallel()
	err := &NotFoundError{SKU: "HAMMER-001"}
	want := "product not found: HAMMER-001"
	if err.Error() != want {
		t.Fatalf("got %q, want %q", err.Error(), want)
	}
}

func TestNotFoundError_Is(t *testing.T) {
	t.Parallel()
	err := &NotFoundError{SKU: "HAMMER-001"}
	if !errors.Is(err, ErrNotFound) {
		t.Fatal("errors.Is(err, ErrNotFound) should be true")
	}
	if errors.Is(err, ErrConflict) {
		t.Fatal("errors.Is(err, ErrConflict) should be false")
	}
}

func TestNotFoundError_As(t *testing.T) {
	t.Parallel()
	var original *NotFoundError = &NotFoundError{SKU: "NAILS-050"}
	// Wrap it to simulate a layer of context.
	wrapped := fmt.Errorf("store lookup: %w", original)

	var target *NotFoundError
	if !errors.As(wrapped, &target) {
		t.Fatal("errors.As should extract *NotFoundError from wrapped error")
	}
	if target.SKU != "NAILS-050" {
		t.Fatalf("got SKU %q, want %q", target.SKU, "NAILS-050")
	}
}

func TestValidationError_Error(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		err    *ValidationError
		want   string
	}{
		{"sku empty", &ValidationError{Field: "sku", Reason: "cannot be empty"}, "validation: sku: cannot be empty"},
		{"price negative", &ValidationError{Field: "price", Reason: "must be positive"}, "validation: price: must be positive"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.err.Error() != tc.want {
				t.Fatalf("got %q, want %q", tc.err.Error(), tc.want)
			}
		})
	}
}

func TestValidationError_Is(t *testing.T) {
	t.Parallel()
	err := &ValidationError{Field: "sku", Reason: "cannot be empty"}
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatal("errors.Is(err, ErrInvalidInput) should be true")
	}
}

func TestValidationError_As(t *testing.T) {
	t.Parallel()
	original := &ValidationError{Field: "name", Reason: "too long"}
	wrapped := fmt.Errorf("service: %w", original)

	var target *ValidationError
	if !errors.As(wrapped, &target) {
		t.Fatal("errors.As should extract *ValidationError")
	}
	if target.Field != "name" {
		t.Fatalf("got Field %q, want %q", target.Field, "name")
	}
}

func TestProduct_String(t *testing.T) {
	t.Parallel()
	p := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}
	s := p.String()
	// Must contain name, sku, and price — exact format is flexible but all three must appear.
	for _, sub := range []string{"Hammer", "HAMMER-001", "25"} {
		if !strings.Contains(s, sub) {
			t.Fatalf("String() = %q, expected it to contain %q", s, sub)
		}
	}
}

func TestProduct_StringVerbsViaFmt(t *testing.T) {
	t.Parallel()
	p := Product{SKU: "NAILS-050", Name: "Nails", Price: 5}
	// %v and %s should both trigger Stringer.
	sv := fmt.Sprintf("%v", p)
	ss := fmt.Sprintf("%s", p)
	if sv == "" {
		t.Fatal("fmt.Sprintf with verb v produced empty string")
	}
	if ss == "" {
		t.Fatal("fmt.Sprintf with verb s produced empty string")
	}
}

func TestCatalogStore_SaveAndGet(t *testing.T) {
	t.Parallel()
	store := NewCatalogStore()
	p := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}

	if err := store.Save(p); err != nil {
		t.Fatalf("Save returned unexpected error: %v", err)
	}

	got, err := store.GetBySKU("HAMMER-001")
	if err != nil {
		t.Fatalf("GetBySKU returned unexpected error: %v", err)
	}
	if got.SKU != p.SKU || got.Name != p.Name {
		t.Fatalf("got %+v, want %+v", got, p)
	}
}

func TestCatalogStore_Save_ValidationError(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		product Product
		field   string
	}{
		{"empty sku", Product{SKU: "", Name: "Hammer", Price: 25}, "sku"},
		{"empty name", Product{SKU: "HAMMER-001", Name: "", Price: 25}, "name"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			store := NewCatalogStore()
			err := store.Save(tc.product)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !errors.Is(err, ErrInvalidInput) {
				t.Fatalf("expected ErrInvalidInput, got %v", err)
			}
			var ve *ValidationError
			if !errors.As(err, &ve) {
				t.Fatalf("expected *ValidationError, got %T", err)
			}
			if ve.Field != tc.field {
				t.Fatalf("expected Field=%q, got %q", tc.field, ve.Field)
			}
		})
	}
}

func TestCatalogStore_Save_Conflict(t *testing.T) {
	t.Parallel()
	store := NewCatalogStore()
	p := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}

	if err := store.Save(p); err != nil {
		t.Fatalf("first Save failed: %v", err)
	}
	err := store.Save(p)
	if err == nil {
		t.Fatal("expected conflict error, got nil")
	}
	if !errors.Is(err, ErrConflict) {
		t.Fatalf("expected ErrConflict, got %v", err)
	}
}

func TestCatalogStore_GetBySKU_NotFound(t *testing.T) {
	t.Parallel()
	store := NewCatalogStore()
	_, err := store.GetBySKU("MISSING-999")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
	var nfe *NotFoundError
	if !errors.As(err, &nfe) {
		t.Fatalf("expected *NotFoundError, got %T", err)
	}
	if nfe.SKU != "MISSING-999" {
		t.Fatalf("expected SKU=%q, got %q", "MISSING-999", nfe.SKU)
	}
}

func TestHTTPStatusFor(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		err  error
		want int
	}{
		{"nil", nil, http.StatusOK},
		{"not found", &NotFoundError{SKU: "X"}, http.StatusNotFound},
		{"conflict", ErrConflict, http.StatusConflict},
		{"invalid input", &ValidationError{Field: "sku", Reason: "empty"}, http.StatusBadRequest},
		{"wrapped not found", fmt.Errorf("svc: %w", &NotFoundError{SKU: "X"}), http.StatusNotFound},
		{"wrapped invalid", fmt.Errorf("svc: %w", &ValidationError{Field: "name", Reason: "empty"}), http.StatusBadRequest},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := HTTPStatusFor(tc.err)
			if got != tc.want {
				t.Fatalf("HTTPStatusFor(%v) = %d, want %d", tc.err, got, tc.want)
			}
		})
	}
}
