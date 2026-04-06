package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var ErrUpstream = errors.New("upstream error")

type UpstreamError struct {
	StatusCode int
	Body       string
}

func (e *UpstreamError) Error() string {
	return fmt.Sprintf("upstream error: status %d: %s", e.StatusCode, e.Body)
}

func (e *UpstreamError) Unwrap() error {
	return ErrUpstream
}

type PriceResponse struct {
	SKU        string `json:"sku"`
	PriceCents int    `json:"price_cents"`
}

type PricingClient struct {
	client  *http.Client
	baseURL string
}

func NewPricingClient(baseURL string, timeout time.Duration) *PricingClient {
	if timeout == 0 {
		timeout = 5 * time.Second
	}
	return &PricingClient{
		client:  &http.Client{Timeout: timeout},
		baseURL: baseURL,
	}
}

func (c *PricingClient) FetchPrice(ctx context.Context, sku string) (PriceResponse, error) {
	url := fmt.Sprintf("%s/prices/%s", c.baseURL, sku)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return PriceResponse{}, fmt.Errorf("build request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return PriceResponse{}, err
	}
	defer resp.Body.Close()

	limited := io.LimitReader(resp.Body, 1<<20)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(limited)
		snippet := string(body)
		if len(snippet) > 256 {
			snippet = snippet[:256]
		}
		return PriceResponse{}, &UpstreamError{
			StatusCode: resp.StatusCode,
			Body:       snippet,
		}
	}

	var price PriceResponse
	if err := json.NewDecoder(limited).Decode(&price); err != nil {
		return PriceResponse{}, fmt.Errorf("decode pricing response: %w", err)
	}
	return price, nil
}

func main() {}
