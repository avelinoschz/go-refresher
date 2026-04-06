package main

import (
	"errors"
	"fmt"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type ProductFetcher interface {
	FetchName(sku string) (string, error)
}

type StaticFetcher struct{}

func (StaticFetcher) FetchName(sku string) (string, error) {
	if sku == "HAMMER-001" {
		return "Hammer", nil
	}

	return "", errors.New("product not found")
}

type MemoryCache struct {
	data map[string]string
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]string),
	}
}

func (c *MemoryCache) Get(key string) (string, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *MemoryCache) Set(key, value string) {
	c.data[key] = value
}

type CatalogService struct {
	cache   Cache
	fetcher ProductFetcher
}

func NewCatalogService(cache Cache, fetcher ProductFetcher) CatalogService {
	return CatalogService{
		cache:   cache,
		fetcher: fetcher,
	}
}

func (s CatalogService) GetProductName(sku string) (string, error) {
	// TODO: implement
	return "", errors.New("not implemented")
}

func main() {
	service := NewCatalogService(NewMemoryCache(), StaticFetcher{})
	name, err := service.GetProductName("HAMMER-001")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("product:", name)
}
