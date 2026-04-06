package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CentPrice int

func (p CentPrice) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%.2f"`, float64(p)/100)
	return []byte(s), nil
}

func (p *CentPrice) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("CentPrice: cannot parse %q: %w", s, err)
	}
	*p = CentPrice(f * 100)
	return nil
}

type InternalProduct struct {
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     CentPrice `json:"price"`
	CostCents int       `json:"cost_cents,omitempty"`
	Available bool      `json:"available"`
}

type PublicProduct struct {
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     CentPrice `json:"price"`
	InStock   bool      `json:"in_stock"`
	CostCents int       `json:"-"`
}

func ToPublic(p InternalProduct) PublicProduct {
	return PublicProduct{
		SKU:     p.SKU,
		Name:    p.Name,
		Price:   p.Price,
		InStock: p.Available,
	}
}

func DecodeProduct(r io.Reader) (InternalProduct, error) {
	var p InternalProduct
	if err := json.NewDecoder(r).Decode(&p); err != nil {
		return InternalProduct{}, err
	}
	return p, nil
}

func EncodePublicProduct(w io.Writer, p PublicProduct) error {
	return json.NewEncoder(w).Encode(p)
}

func main() {}
