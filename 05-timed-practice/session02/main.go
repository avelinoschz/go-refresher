package main

import "time"

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CatalogResponse struct {
	Product
	Source   string `json:"source,omitempty"`
	ServedAt string `json:"served_at,omitempty"`
}

func BuildResponse(product Product, includeMeta bool, source string, now time.Time) CatalogResponse {
	// TODO: implement
	return CatalogResponse{}
}

func main() {}
