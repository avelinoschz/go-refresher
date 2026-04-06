package main

import (
	"fmt"
	"sync"
)

type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]string),
	}
}

func (c *MemoryCache) Get(key string) (string, bool) {
	// TODO: implement
	return "", false
}

func (c *MemoryCache) Set(key, value string) {
	// TODO: implement
}

func main() {
	cache := NewMemoryCache()
	cache.Set("catalog:HAMMER-001", "Hammer")
	value, ok := cache.Get("catalog:HAMMER-001")
	fmt.Println("found:", ok, "value:", value)
}
