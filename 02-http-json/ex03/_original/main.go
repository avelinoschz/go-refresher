package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	// TODO: implement
	return nil
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	// TODO: implement
}

func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder
		// TODO: implement
	}
}

func main() {
	http.HandleFunc("GET /health", healthHandler())

	_ = http.ListenAndServe(":8080", nil)
}
