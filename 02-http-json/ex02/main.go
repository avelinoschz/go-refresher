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
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var prod Product
		err := json.NewDecoder(r.Body).Decode(&prod)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if prod.SKU == "" || prod.Name == "" || prod.Price <= 0 {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if store.Exists(prod.SKU) {
			http.Error(w, "product already exists", http.StatusConflict)
			return
		}

		if err := store.Save(prod); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	store := NewCatalogStore()

	http.HandleFunc("POST /catalog", createCatalogHandler(store))

	_ = http.ListenAndServe(":8080", nil)
}
