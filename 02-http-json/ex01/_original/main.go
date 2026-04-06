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

type CatalogService struct {
	products map[string]Product
}

func NewCatalogService() CatalogService {
	return CatalogService{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer", Price: 25},
			"NAILS-050":  {SKU: "NAILS-050", Name: "Nails", Price: 5},
		},
	}
}

func (s CatalogService) GetBySKU(sku string) (Product, bool) {
	product, ok := s.products[sku]
	return product, ok
}

func catalogHandler(service CatalogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = service
		_ = json.NewEncoder
		// TODO: implement
	}
}

func main() {
	service := NewCatalogService()

	http.HandleFunc("GET /catalog", catalogHandler(service))

	_ = http.ListenAndServe(":8080", nil)
}
