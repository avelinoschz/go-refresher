package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CatalogClient struct {
	baseURL string
	client  *http.Client
}

func NewCatalogClient(baseURL string, timeout time.Duration) *CatalogClient {
	return &CatalogClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: timeout},
	}
}

func (c *CatalogClient) GetProduct(ctx context.Context, sku string) (Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/catalog?sku="+sku, nil)
	if err != nil {
		return Product{}, fmt.Errorf("build request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Product{}, fmt.Errorf("do request: %w", err)
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

func (c *CatalogClient) CreateProduct(ctx context.Context, p Product) error {
	body, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("marshal product: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/catalog", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}

func main() {}
