package main

import "testing"

func TestInventoryAddProductCreatesAndAggregatesStock(t *testing.T) {
	t.Parallel()

	inventory := &Inventory{}

	if err := inventory.AddProduct("HAMMER-001", 5); err != nil {
		t.Fatalf("add initial product: %v", err)
	}

	if err := inventory.AddProduct("HAMMER-001", 2); err != nil {
		t.Fatalf("add duplicate product: %v", err)
	}

	product, found := inventory.FindBySKU("HAMMER-001")
	if !found {
		t.Fatal("expected product to be found")
	}

	if product.Quantity != 7 {
		t.Fatalf("unexpected quantity: got %d want %d", product.Quantity, 7)
	}
}

func TestInventoryAddProductValidatesInput(t *testing.T) {
	t.Parallel()

	inventory := &Inventory{}

	if err := inventory.AddProduct("", 1); err == nil {
		t.Fatal("expected error for empty sku")
	}

	if err := inventory.AddProduct("HAMMER-001", 0); err == nil {
		t.Fatal("expected error for non-positive quantity")
	}
}

func TestInventoryRemoveProductUpdatesStock(t *testing.T) {
	t.Parallel()

	inventory := &Inventory{}
	_ = inventory.AddProduct("NAILS-050", 50)

	if err := inventory.RemoveProduct("NAILS-050", 10); err != nil {
		t.Fatalf("remove product: %v", err)
	}

	product, found := inventory.FindBySKU("NAILS-050")
	if !found {
		t.Fatal("expected product to be found")
	}

	if product.Quantity != 40 {
		t.Fatalf("unexpected quantity: got %d want %d", product.Quantity, 40)
	}
}

func TestInventoryRemoveProductReturnsErrors(t *testing.T) {
	t.Parallel()

	inventory := &Inventory{}
	_ = inventory.AddProduct("HAMMER-001", 5)

	if err := inventory.RemoveProduct("HAMMER-001", 0); err == nil {
		t.Fatal("expected error for non-positive quantity")
	}

	if err := inventory.RemoveProduct("HAMMER-001", 10); err == nil {
		t.Fatal("expected error for insufficient stock")
	}

	if err := inventory.RemoveProduct("MISSING-001", 1); err == nil {
		t.Fatal("expected error for missing product")
	}
}

func TestInventoryTotalQuantity(t *testing.T) {
	t.Parallel()

	inventory := &Inventory{}
	_ = inventory.AddProduct("HAMMER-001", 5)
	_ = inventory.AddProduct("NAILS-050", 50)
	_ = inventory.RemoveProduct("NAILS-050", 10)

	if got := inventory.TotalQuantity(); got != 45 {
		t.Fatalf("unexpected total quantity: got %d want %d", got, 45)
	}
}
