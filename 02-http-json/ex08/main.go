package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// FetchProduct calls GET <baseURL>/catalog?sku=<sku> and decodes the JSON
// response into a Product. It returns an error if the response status is not
// 200 OK.
func FetchProduct(baseURL, sku string) (Product, error) {
	_ = http.Get
	_ = io.ReadAll
	_ = json.NewDecoder
	// TODO: implement
	// Requirements:
	// - build the URL with the sku query param
	// - perform the GET request
	// - close the response body
	// - return an error if the status is not 200
	// - decode the JSON body into a Product and return it
	return Product{}, nil
}

// CreateProduct calls POST <baseURL>/catalog with the JSON-encoded product.
// It returns an error if the response status is not 201 Created.
func CreateProduct(baseURL string, p Product) error {
	_ = bytes.NewBuffer
	_ = json.Marshal
	_ = fmt.Errorf
	// TODO: implement
	// Requirements:
	// - marshal p to JSON
	// - POST to <baseURL>/catalog with the JSON body
	// - close the response body
	// - return an error if the status is not 201
	return nil
}

func main() {}
