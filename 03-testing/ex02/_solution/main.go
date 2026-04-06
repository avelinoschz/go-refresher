package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

type CatalogService struct {
	products map[string]Product
}

func NewCatalogService() CatalogService {
	return CatalogService{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer"},
		},
	}
}

func catalogHandler(service CatalogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sku := r.URL.Query().Get("sku")
		if sku == "" {
			http.Error(w, "missing sku", http.StatusBadRequest)
			return
		}

		product, ok := service.products[sku]
		if !ok {
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

func main() {
	service := NewCatalogService()
	_ = http.ListenAndServe(":8080", catalogHandler(service))
}
