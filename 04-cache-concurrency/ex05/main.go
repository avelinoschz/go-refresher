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

// ErrUpstream is the sentinel error for non-2xx upstream responses.
// Callers use errors.Is(err, ErrUpstream) to detect upstream failures.
var ErrUpstream = errors.New("upstream error")

// UpstreamError is returned when the pricing API responds with a non-2xx
// status code. It wraps ErrUpstream.
type UpstreamError struct {
	StatusCode int
	Body       string // first 256 bytes of response body, for debugging
}

func (e *UpstreamError) Error() string {
	_ = fmt.Sprintf
	// TODO: implement — "upstream error: status <code>: <body>"
	return ""
}

func (e *UpstreamError) Unwrap() error {
	// TODO: implement
	return nil
}

// PriceResponse is what the upstream pricing API returns.
type PriceResponse struct {
	SKU        string `json:"sku"`
	PriceCents int    `json:"price_cents"`
}

// PricingClient calls an upstream pricing API.
type PricingClient struct {
	client  *http.Client
	baseURL string
}

// NewPricingClient constructs a PricingClient.
// If timeout is zero, 5 seconds is used.
func NewPricingClient(baseURL string, timeout time.Duration) *PricingClient {
	if timeout == 0 {
		timeout = 5 * time.Second
	}
	return &PricingClient{
		client:  &http.Client{Timeout: timeout},
		baseURL: baseURL,
	}
}

// FetchPrice calls GET <baseURL>/prices/<sku> with the provided context.
//
// Rules:
//   - Use http.NewRequestWithContext (never http.NewRequest).
//   - Read the response body with io.LimitReader (max 1 MB) and always close it.
//   - Non-2xx responses → return *UpstreamError wrapping ErrUpstream.
//   - JSON decode error → wrap with fmt.Errorf("decode pricing response: %w", err).
func (c *PricingClient) FetchPrice(ctx context.Context, sku string) (PriceResponse, error) {
	_ = json.NewDecoder
	_ = io.LimitReader
	// TODO: implement
	return PriceResponse{}, nil
}

func main() {}
