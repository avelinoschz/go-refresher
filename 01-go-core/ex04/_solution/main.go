package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrInvalidInput = errors.New("invalid input")
)

type NotFoundError struct {
	SKU string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("product not found: %s", e.SKU)
}

func (e *NotFoundError) Unwrap() error {
	return ErrNotFound
}

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s: %s", e.Field, e.Reason)
}

func (e *ValidationError) Unwrap() error {
	return ErrInvalidInput
}

type Product struct {
	SKU   string
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("%s (sku=%s, price=%d)", p.Name, p.SKU, p.Price)
}

type CatalogStore struct {
	products map[string]Product
}

func NewCatalogStore() *CatalogStore {
	return &CatalogStore{products: make(map[string]Product)}
}

func (s *CatalogStore) Save(p Product) error {
	if p.SKU == "" {
		return &ValidationError{Field: "sku", Reason: "cannot be empty"}
	}
	if p.Name == "" {
		return &ValidationError{Field: "name", Reason: "cannot be empty"}
	}
	if _, exists := s.products[p.SKU]; exists {
		return ErrConflict
	}
	s.products[p.SKU] = p
	return nil
}

func (s *CatalogStore) GetBySKU(sku string) (Product, error) {
	p, ok := s.products[sku]
	if !ok {
		return Product{}, &NotFoundError{SKU: sku}
	}
	return p, nil
}

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
