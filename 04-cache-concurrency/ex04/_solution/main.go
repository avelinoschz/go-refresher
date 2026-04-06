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
	priceCh := make(chan result[int], 1)
	stockCh := make(chan result[int], 1)

	go func() {
		price, err := fetchPrice(sku)
		priceCh <- result[int]{value: price, err: err}
	}()

	go func() {
		stock, err := fetchStock(sku)
		stockCh <- result[int]{value: stock, err: err}
	}()

	priceResult := <-priceCh
	if priceResult.err != nil {
		return ProductSnapshot{}, priceResult.err
	}

	stockResult := <-stockCh
	if stockResult.err != nil {
		return ProductSnapshot{}, stockResult.err
	}

	return ProductSnapshot{
		SKU:   sku,
		Price: priceResult.value,
		Stock: stockResult.value,
	}, nil
}

func main() {
	snapshot, err := LoadSnapshot("HAMMER-001")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("snapshot: %+v\n", snapshot)
}
