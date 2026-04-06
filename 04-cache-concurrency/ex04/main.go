package main

import (
	"errors"
	"fmt"
)

type ProductSnapshot struct {
	SKU   string
	Price int
	Stock int
}

type result[T any] struct {
	value T
	err   error
}

func fetchPrice(sku string) (int, error) {
	if sku == "" {
		return 0, errors.New("sku is required")
	}

	return 25, nil
}

func fetchStock(sku string) (int, error) {
	if sku == "" {
		return 0, errors.New("sku is required")
	}

	return 12, nil
}

func LoadSnapshot(sku string) (ProductSnapshot, error) {
	// TODO: implement
	return ProductSnapshot{}, errors.New("not implemented")
}

func main() {
	snapshot, err := LoadSnapshot("HAMMER-001")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("snapshot: %+v\n", snapshot)
}
