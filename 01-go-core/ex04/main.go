package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Sentinel errors — use errors.Is against these.
var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrInvalidInput = errors.New("invalid input")
)

// NotFoundError is returned when a product does not exist in the store.
// It wraps ErrNotFound so errors.Is(err, ErrNotFound) == true.
type NotFoundError struct {
	SKU string
}

func (e *NotFoundError) Error() string {
	// TODO: implement — "product not found: <sku>"
	return fmt.Sprintf("product not found: %s", e.SKU)
}

func (e *NotFoundError) Unwrap() error {
	return ErrNotFound
}

// ValidationError is returned when a product fails field validation.
// It wraps ErrInvalidInput so errors.Is(err, ErrInvalidInput) == true.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	// TODO: implement — "validation: <field>: <reason>"
	return fmt.Sprintf("validation: %s: %s", e.Field, e.Reason)
}

func (e *ValidationError) Unwrap() error {
	return ErrInvalidInput
}

// Product is a catalog item.
// It implements fmt.Stringer.
type Product struct {
	SKU   string
	Name  string
	Price int
}

func (p Product) String() string {
	// TODO: implement — "<name> (sku=<sku>, price=<price>)"
	return fmt.Sprintf("%s (sku=%s, price=%d)", p.Name, p.SKU, p.Price)
}

// CatalogStore is an in-memory product store.
type CatalogStore struct {
	products map[string]Product
}

func NewCatalogStore() *CatalogStore {
	return &CatalogStore{products: make(map[string]Product)}
}

// Save persists a product.
// Returns *ValidationError if SKU or Name is empty.
// Returns ErrConflict if a product with the same SKU already exists.
func (s *CatalogStore) Save(p Product) error {
	if p.Name == "" {
		return &ValidationError{Field: "name", Reason: "is required"}
	}

	if p.SKU == "" {
		return &ValidationError{Field: "sku", Reason: "is required"}
	}

	if _, exists := s.products[p.SKU]; exists {
		return ErrConflict
	}

	s.products[p.SKU] = p

	return nil
}

// GetBySKU looks up a product by SKU.
// Returns *NotFoundError if the SKU does not exist.
func (s *CatalogStore) GetBySKU(sku string) (Product, error) {
	p, ok := s.products[sku]
	if !ok {
		return Product{}, &NotFoundError{SKU: sku}
	}
	return p, nil
}

// HTTPStatusFor maps a catalog error to the appropriate HTTP status code.
// Uses errors.Is to discriminate — never inspect error strings.
// Returns 200 for nil.
func HTTPStatusFor(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, ErrConflict) {
		return http.StatusConflict
	}

	if errors.Is(err, ErrInvalidInput) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func main() {}
