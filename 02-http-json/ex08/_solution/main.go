package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func FetchProduct(baseURL, sku string) (Product, error) {
	resp, err := http.Get(baseURL + "/catalog?sku=" + sku)
	if err != nil {
		return Product{}, fmt.Errorf("get product: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Product{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var p Product
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return Product{}, fmt.Errorf("decode product: %w", err)
	}

	return p, nil
}

func CreateProduct(baseURL string, p Product) error {
	body, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("marshal product: %w", err)
	}

	resp, err := http.Post(baseURL+"/catalog", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("post product: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}

func main() {}
