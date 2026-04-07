package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type PriceUpdateRequest struct {
	Price int `json:"price"`
}

type CatalogStore struct {
	products map[string]Product
}

func NewCatalogStore() *CatalogStore {
	return &CatalogStore{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer", Price: 25},
			"NAILS-050":  {SKU: "NAILS-050", Name: "Nails", Price: 5},
		},
	}
}

func (s *CatalogStore) UpdatePrice(sku string, price int) (Product, bool) {
	product, ok := s.products[sku]
	if !ok {
		return Product{}, false
	}

	product.Price = price
	s.products[sku] = product
	return product, true
}

func decodeStrictJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	_ = json.NewDecoder
	_ = http.MaxBytesReader
	_ = io.EOF
	_ = errors.New
	// TODO: implement
	// Requirements:
	// - limit the request body size
	// - use json.Decoder
	// - reject unknown fields
	// - reject trailing JSON after the first object
	return nil
}

func updatePriceHandler(store *CatalogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue
		_ = store
		// TODO: implement
		// Requirements:
		// - read sku from the route
		// - decode the JSON body strictly
		// - return 400 / 404 / 413 when appropriate
		// - return the updated product as JSON on success
	}
}

func main() {
	store := NewCatalogStore()

	http.HandleFunc("PUT /catalog/{sku}/price", updatePriceHandler(store))

	_ = http.ListenAndServe(":8080", nil)
}
