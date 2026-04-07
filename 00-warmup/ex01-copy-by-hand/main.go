package main

import (
	"errors"
	"fmt"
)

type Product struct {
	SKU      string
	Quantity int
}

type Inventory struct {
	products []Product
}

func (i *Inventory) AddProduct(sku string, quantity int) error {
	if sku == "" {
		return errors.New("sku is required")
	}

	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	for index := range i.products {
		if i.products[index].SKU == sku {
			i.products[index].Quantity += quantity
			return nil
		}
	}

	i.products = append(i.products, Product{
		SKU:      sku,
		Quantity: quantity,
	})

	return nil
}

func (i *Inventory) RemoveProduct(sku string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	for index := range i.products {
		product := &i.products[index]
		if product.SKU != sku {
			continue
		}

		if product.Quantity < quantity {
			return errors.New("insufficient stock")
		}

		product.Quantity -= quantity
		return nil
	}

	return errors.New("product not found")
}

func (i *Inventory) FindBySKU(sku string) (Product, bool) {
	for _, product := range i.products {
		if product.SKU == sku {
			return product, true
		}
	}

	return Product{}, false
}

func (i *Inventory) TotalQuantity() int {
	total := 0
	for _, product := range i.products {
		total += product.Quantity
	}

	return total
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
