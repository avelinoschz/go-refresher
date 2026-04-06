package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Cache interface {
	Get(key string) (Product, bool)
	Set(key string, value Product)
}

type ProductStore interface {
	GetBySKU(sku string) (Product, error)
}

type MemoryCache struct {
	data map[string]Product
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]Product),
	}
}

func (c *MemoryCache) Get(key string) (Product, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *MemoryCache) Set(key string, value Product) {
	c.data[key] = value
}

type StaticStore struct {
	products map[string]Product
}

func NewStaticStore() StaticStore {
	return StaticStore{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer", Price: 25},
			"DRILL-002":  {SKU: "DRILL-002", Name: "Drill", Price: 80},
		},
	}
}

func (s StaticStore) GetBySKU(sku string) (Product, error) {
	product, ok := s.products[sku]
	if !ok {
		return Product{}, errors.New("product not found")
	}

	return product, nil
}

type CatalogService struct {
	cache Cache
	store ProductStore
}

func NewCatalogService(cache Cache, store ProductStore) CatalogService {
	return CatalogService{
		cache: cache,
		store: store,
	}
}

func (s CatalogService) GetProduct(sku string) (Product, error) {
	if product, ok := s.cache.Get(sku); ok {
		return product, nil
	}

	product, err := s.store.GetBySKU(sku)
	if err != nil {
		return Product{}, err
	}

	s.cache.Set(sku, product)
	return product, nil
}

func catalogHandler(service CatalogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sku := r.URL.Query().Get("sku")
		if sku == "" {
			http.Error(w, "missing sku", http.StatusBadRequest)
			return
		}

		product, err := service.GetProduct(sku)
		if err != nil {
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

func main() {
	service := NewCatalogService(NewMemoryCache(), NewStaticStore())
	_ = http.ListenAndServe(":8080", catalogHandler(service))
}
