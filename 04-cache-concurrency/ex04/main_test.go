package main

import "testing"

func TestLoadSnapshotCombinesPriceAndStock(t *testing.T) {
	t.Parallel()

	got, err := LoadSnapshot("HAMMER-001")
	if err != nil {
		t.Fatalf("load snapshot: %v", err)
	}

	want := ProductSnapshot{
		SKU:   "HAMMER-001",
		Price: 25,
		Stock: 12,
	}

	if got != want {
		t.Fatalf("unexpected snapshot: got %+v want %+v", got, want)
	}
}

func TestLoadSnapshotReturnsErrorWhenInputIsInvalid(t *testing.T) {
	t.Parallel()

	if _, err := LoadSnapshot(""); err == nil {
		t.Fatal("expected error for empty sku")
	}
}
