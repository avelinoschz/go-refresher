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
	// TODO: implement
	return Product{}, errors.New("not implemented")
}

func catalogHandler(service CatalogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder
		// TODO: implement
	}
}

func main() {
	service := NewCatalogService(NewMemoryCache(), NewStaticStore())
	_ = http.ListenAndServe(":8080", catalogHandler(service))
}
