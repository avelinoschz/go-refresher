package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

// ---- CentPrice ----

func TestCentPrice_MarshalJSON(t *testing.T) {
	t.Parallel()
	cases := []struct {
		cents CentPrice
		want  string // expected JSON token including surrounding quotes
	}{
		{1299, `"12.99"`},
		{500, `"5.00"`},
		{0, `"0.00"`},
		{1, `"0.01"`},
	}
	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			b, err := json.Marshal(tc.cents)
			if err != nil {
				t.Fatalf("MarshalJSON error: %v", err)
			}
			if string(b) != tc.want {
				t.Fatalf("got %s, want %s", b, tc.want)
			}
		})
	}
}

func TestCentPrice_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input string
		want  CentPrice
	}{
		{`"12.99"`, 1299},
		{`"5.00"`, 500},
		{`"0.01"`, 1},
		{`"0.00"`, 0},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()
			var p CentPrice
			if err := json.Unmarshal([]byte(tc.input), &p); err != nil {
				t.Fatalf("UnmarshalJSON error: %v", err)
			}
			if p != tc.want {
				t.Fatalf("got %d, want %d", p, tc.want)
			}
		})
	}
}

func TestCentPrice_RoundTrip(t *testing.T) {
	t.Parallel()
	original := CentPrice(4250)
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var restored CentPrice
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if restored != original {
		t.Fatalf("round-trip failed: got %d, want %d", restored, original)
	}
}

// ---- omitempty and "-" tags ----

func TestInternalProduct_OmitEmptyCostCents(t *testing.T) {
	t.Parallel()
	p := InternalProduct{SKU: "X", Name: "Y", Price: 100, CostCents: 0, Available: true}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if strings.Contains(string(b), "cost_cents") {
		t.Fatalf("cost_cents should be omitted when zero, got: %s", b)
	}
}

func TestInternalProduct_IncludesCostCentsWhenNonZero(t *testing.T) {
	t.Parallel()
	p := InternalProduct{SKU: "X", Name: "Y", Price: 100, CostCents: 800, Available: true}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if !strings.Contains(string(b), "cost_cents") {
		t.Fatalf("cost_cents should be present when non-zero, got: %s", b)
	}
}

func TestPublicProduct_CostCentsAlwaysHidden(t *testing.T) {
	t.Parallel()
	p := PublicProduct{SKU: "X", Name: "Y", Price: 100, InStock: true, CostCents: 9999}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if strings.Contains(string(b), "cost") {
		t.Fatalf("cost should never appear in PublicProduct JSON, got: %s", b)
	}
}

// ---- ToPublic ----

func TestToPublic(t *testing.T) {
	t.Parallel()
	internal := InternalProduct{
		SKU:       "HAMMER-001",
		Name:      "Hammer",
		Price:     2500,
		CostCents: 1200,
		Available: true,
	}
	pub := ToPublic(internal)

	if pub.SKU != internal.SKU {
		t.Fatalf("SKU: got %q, want %q", pub.SKU, internal.SKU)
	}
	if pub.Price != internal.Price {
		t.Fatalf("Price: got %d, want %d", pub.Price, internal.Price)
	}
	if pub.InStock != internal.Available {
		t.Fatalf("InStock: got %v, want %v", pub.InStock, internal.Available)
	}
	// CostCents should not leak through (it won't appear in JSON due to "-" tag,
	// but the struct field itself should be zero or ignored).
	b, _ := json.Marshal(pub)
	if strings.Contains(string(b), "cost") {
		t.Fatalf("cost must not appear in public JSON: %s", b)
	}
}

// ---- Streaming ----

func TestDecodeProduct(t *testing.T) {
	t.Parallel()
	input := `{"sku":"NAILS-050","name":"Nails","price":"5.00","cost_cents":200,"available":true}`
	p, err := DecodeProduct(strings.NewReader(input))
	if err != nil {
		t.Fatalf("DecodeProduct error: %v", err)
	}
	if p.SKU != "NAILS-050" {
		t.Fatalf("SKU: got %q, want %q", p.SKU, "NAILS-050")
	}
	if p.Price != 500 {
		t.Fatalf("Price: got %d, want 500", p.Price)
	}
	if p.CostCents != 200 {
		t.Fatalf("CostCents: got %d, want 200", p.CostCents)
	}
}

func TestDecodeProduct_InvalidJSON(t *testing.T) {
	t.Parallel()
	_, err := DecodeProduct(strings.NewReader(`{invalid`))
	if err == nil {
		t.Fatal("expected error for malformed JSON, got nil")
	}
}

func TestEncodePublicProduct(t *testing.T) {
	t.Parallel()
	p := PublicProduct{SKU: "HAMMER-001", Name: "Hammer", Price: 2500, InStock: true}
	var buf bytes.Buffer
	if err := EncodePublicProduct(&buf, p); err != nil {
		t.Fatalf("EncodePublicProduct error: %v", err)
	}

	var decoded PublicProduct
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("could not decode output: %v — raw: %s", err, buf.String())
	}
	if decoded.SKU != p.SKU {
		t.Fatalf("SKU: got %q, want %q", decoded.SKU, p.SKU)
	}
	if decoded.Price != p.Price {
		t.Fatalf("Price: got %d, want %d", decoded.Price, p.Price)
	}
}
