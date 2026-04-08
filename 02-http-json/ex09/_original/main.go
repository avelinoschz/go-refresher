package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// CatalogClient is an HTTP client scoped to a catalog service base URL.
type CatalogClient struct {
	baseURL string
	client  *http.Client
}

// NewCatalogClient returns a CatalogClient that uses a dedicated http.Client
// with the given timeout.
func NewCatalogClient(baseURL string, timeout time.Duration) *CatalogClient {
	_ = http.Client{}
	// TODO: implement
	return &CatalogClient{}
}

// GetProduct calls GET <baseURL>/catalog?sku=<sku> using http.NewRequestWithContext.
// It returns an error if the context is canceled or the response status is not 200.
func (c *CatalogClient) GetProduct(ctx context.Context, sku string) (Product, error) {
	_ = http.NewRequestWithContext
	_ = io.ReadAll
	_ = json.NewDecoder
	_ = fmt.Errorf
	// TODO: implement
	// Requirements:
	// - build the request with ctx using http.NewRequestWithContext
	// - perform the request using c.client.Do
	// - close the response body
	// - return an error if the status is not 200
	// - decode the JSON body into a Product and return it
	return Product{}, nil
}

// CreateProduct calls POST <baseURL>/catalog with the JSON-encoded product.
// It sets the Content-Type header explicitly to "application/json".
// It returns an error if the context is canceled or the response status is not 201.
func (c *CatalogClient) CreateProduct(ctx context.Context, p Product) error {
	_ = bytes.NewBuffer
	_ = json.Marshal
	// TODO: implement
	// Requirements:
	// - marshal p to JSON
	// - build the request with ctx using http.NewRequestWithContext
	// - set the Content-Type header to "application/json"
	// - perform the request using c.client.Do
	// - close the response body
	// - return an error if the status is not 201
	return nil
}

func main() {}
