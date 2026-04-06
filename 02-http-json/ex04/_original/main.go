package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// CentPrice represents a price stored as an integer number of cents.
// It serializes to/from a JSON decimal string: 1299 <-> "12.99".
type CentPrice int

func (p CentPrice) MarshalJSON() ([]byte, error) {
	_ = fmt.Sprintf
	// TODO: implement — produce a JSON string like "12.99"
	// Hint: fmt.Sprintf("%.2f", float64(p)/100) then wrap with quotes.
	return nil, nil
}

func (p *CentPrice) UnmarshalJSON(data []byte) error {
	// TODO: implement — parse a JSON string "12.99" into cents (1299).
	// Hint: strip quotes, parse float, multiply by 100.
	return nil
}

// InternalProduct is the full admin representation of a product.
// CostCents is omitted from JSON output when zero.
type InternalProduct struct {
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     CentPrice `json:"price"`
	CostCents int       `json:"cost_cents,omitempty"`
	Available bool      `json:"available"`
}

// PublicProduct is the storefront view.
// CostCents is never included in JSON output (json:"-").
// InStock mirrors Available.
type PublicProduct struct {
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     CentPrice `json:"price"`
	InStock   bool      `json:"in_stock"`
	CostCents int       `json:"-"`
}

// ToPublic converts an InternalProduct to a PublicProduct.
func ToPublic(p InternalProduct) PublicProduct {
	// TODO: implement
	return PublicProduct{}
}

// DecodeProduct decodes a single InternalProduct from r using json.Decoder.
// Do not use json.Unmarshal(io.ReadAll(r), ...).
func DecodeProduct(r io.Reader) (InternalProduct, error) {
	_ = json.NewDecoder
	// TODO: implement
	return InternalProduct{}, nil
}

// EncodePublicProduct encodes p to w using json.Encoder.
// Do not use json.Marshal + w.Write.
func EncodePublicProduct(w io.Writer, p PublicProduct) error {
	// TODO: implement
	return nil
}

func main() {}
