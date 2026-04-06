package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CatalogStore struct {
	products map[string]Product
}

func NewCatalogStore() *CatalogStore {
	return &CatalogStore{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer", Price: 25},
		},
	}
}

func (s *CatalogStore) Save(product Product) error {
	s.products[product.SKU] = product
	return nil
}

func (s *CatalogStore) Exists(sku string) bool {
	_, ok := s.products[sku]
	return ok
}

func createCatalogHandler(store *CatalogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if product.SKU == "" || product.Name == "" || product.Price <= 0 {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if store.Exists(product.SKU) {
			http.Error(w, "product already exists", http.StatusConflict)
			return
		}

		store.Save(product)
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	store := NewCatalogStore()

	http.HandleFunc("POST /catalog", createCatalogHandler(store))

	_ = http.ListenAndServe(":8080", nil)
}
