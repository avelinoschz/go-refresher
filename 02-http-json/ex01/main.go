package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CatalogService struct {
	products map[string]Product
}

func NewCatalogService() CatalogService {
	return CatalogService{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer", Price: 25},
			"NAILS-050":  {SKU: "NAILS-050", Name: "Nails", Price: 5},
		},
	}
}

func (s CatalogService) GetBySKU(sku string) (Product, bool) {
	product, ok := s.products[sku]
	return product, ok
}

func catalogHandler(service CatalogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		sku := params.Get("sku")
		if sku == "" {
			http.Error(w, "sku is required", http.StatusBadRequest)
			return
		}

		prod, found := service.GetBySKU(sku)
		if !found {
			http.Error(w, "sku not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(prod)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	service := NewCatalogService()

	http.HandleFunc("GET /catalog", catalogHandler(service))

	_ = http.ListenAndServe(":8080", nil)
}

// // Just practicing. This is a helper function to write a error response
// func writeJSONError(w http.ResponseWriter, status int, msg string) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)

// 	// using literal strings
// 	// fmt.Fprintf(w, `{"error":%q}`, msg)

// 	// or jsn marhasling
// 	resp := ErrorReponse{
// 		Error: msg,
// 	}
// 	payload, err := json.Marshal(resp)
// 	if err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "%s", payload)
// }

// type ErrorReponse struct {
// 	Error string `json:"error"`
// }
