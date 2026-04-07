package main

import "fmt"

type Product struct {
	SKU      string
	Quantity int
}

type Inventory struct {
	products []Product
}

func (i *Inventory) AddProduct(sku string, quantity int) error {
	// TODO: implement
	return nil
}

func (i *Inventory) RemoveProduct(sku string, quantity int) error {
	// TODO: implement
	return nil
}

func (i *Inventory) FindBySKU(sku string) (Product, bool) {
	// TODO: implement
	return Product{}, false
}

func (i *Inventory) TotalQuantity() int {
	// TODO: implement
	return 0
}

func main() {
	inventory := &Inventory{}

	must(inventory.AddProduct("HAMMER-001", 5))
	must(inventory.AddProduct("NAILS-050", 50))
	must(inventory.AddProduct("HAMMER-001", 2))
	must(inventory.RemoveProduct("NAILS-050", 10))

	product, found := inventory.FindBySKU("HAMMER-001")
	fmt.Println("found hammer:", found)
	fmt.Printf("hammer stock: %+v\n", product)
	fmt.Println("total quantity:", inventory.TotalQuantity())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
