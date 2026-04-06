package main

import (
	"testing"
	"time"
)

func TestBuildResponseWithoutMeta(t *testing.T) {
	t.Parallel()

	product := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}
	resp := BuildResponse(product, false, "store", time.Now())

	if resp.SKU != "HAMMER-001" {
		t.Fatalf("unexpected sku: got %q want %q", resp.SKU, "HAMMER-001")
	}
	if resp.Source != "" {
		t.Fatalf("expected empty source when includeMeta=false, got %q", resp.Source)
	}
	if resp.ServedAt != "" {
		t.Fatalf("expected empty served_at when includeMeta=false, got %q", resp.ServedAt)
	}
}

func TestBuildResponseWithMetaFromCache(t *testing.T) {
	t.Parallel()

	product := Product{SKU: "HAMMER-001", Name: "Hammer", Price: 25}
	now := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)
	resp := BuildResponse(product, true, "cache", now)

	if resp.Source != "cache" {
		t.Fatalf("unexpected source: got %q want %q", resp.Source, "cache")
	}
	if resp.ServedAt != "2026-01-15T12:00:00Z" {
		t.Fatalf("unexpected served_at: got %q want %q", resp.ServedAt, "2026-01-15T12:00:00Z")
	}
}

func TestBuildResponseWithMetaFromStore(t *testing.T) {
	t.Parallel()

	product := Product{SKU: "DRILL-002", Name: "Drill", Price: 80}
	now := time.Now()
	resp := BuildResponse(product, true, "store", now)

	if resp.Source != "store" {
		t.Fatalf("unexpected source: got %q want %q", resp.Source, "store")
	}
	if resp.ServedAt == "" {
		t.Fatal("expected non-empty served_at when includeMeta=true")
	}
	if resp.Price != 80 {
		t.Fatalf("unexpected price: got %d want %d", resp.Price, 80)
	}
}
