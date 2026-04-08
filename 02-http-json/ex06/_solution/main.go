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
	r.Body = http.MaxBytesReader(w, r.Body, 64)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must contain a single JSON object")
	}

	return nil
}

func updatePriceHandler(store *CatalogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sku := r.PathValue("sku")
		if sku == "" {
			http.Error(w, "sku is required", http.StatusBadRequest)
			return
		}

		var input PriceUpdateRequest
		if err := decodeStrictJSON(w, r, &input); err != nil {
			var maxErr *http.MaxBytesError
			if errors.As(err, &maxErr) {
				http.Error(w, "request body too large", http.StatusRequestEntityTooLarge)
				return
			}

			http.Error(w, "invalid json body", http.StatusBadRequest)
			return
		}

		if input.Price <= 0 {
			http.Error(w, "price must be greater than 0", http.StatusBadRequest)
			return
		}

		product, ok := store.UpdatePrice(sku, input.Price)
		if !ok {
			http.Error(w, "sku not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	store := NewCatalogStore()

	http.HandleFunc("PUT /catalog/{sku}/price", updatePriceHandler(store))

	_ = http.ListenAndServe(":8080", nil)
}
