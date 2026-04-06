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
	_ = fmt.Sprintf
	// TODO: implement — "product not found: <sku>"
	return ""
}

func (e *NotFoundError) Unwrap() error {
	// TODO: implement
	return nil
}

// ValidationError is returned when a product fails field validation.
// It wraps ErrInvalidInput so errors.Is(err, ErrInvalidInput) == true.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	// TODO: implement — "validation: <field>: <reason>"
	return ""
}

func (e *ValidationError) Unwrap() error {
	// TODO: implement
	return nil
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
	return ""
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
	// TODO: implement
	return nil
}

// GetBySKU looks up a product by SKU.
// Returns *NotFoundError if the SKU does not exist.
func (s *CatalogStore) GetBySKU(sku string) (Product, error) {
	// TODO: implement
	return Product{}, nil
}

// HTTPStatusFor maps a catalog error to the appropriate HTTP status code.
// Uses errors.Is to discriminate — never inspect error strings.
// Returns 200 for nil.
func HTTPStatusFor(err error) int {
	_ = http.StatusOK
	// TODO: implement
	return 0
}

func main() {}
