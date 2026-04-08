package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return err
	}

	return nil
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	_ = writeJSON(w, statusCode, ErrorResponse{
		Error: message,
	})
}

func getServiceName() string {
	return "catalog-api"
}

// func healthHandler(logger *slog.Logger) http.HandlerFunc {
func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// reqLogger := logger.With(
		// 	"method", r.Method,
		// 	"path", r.URL.Path,
		// 	"remote_addr", r.RemoteAddr,
		// )
		// reqLogger.Info("health check received")

		if r.Method != http.MethodGet {
			// reqLogger.Error("request failed",
			// 	"status_code", http.StatusMethodNotAllowed,
			// 	"error", "method not allowed")
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		if err := writeJSON(w, http.StatusOK, HealthResponse{
			Status:  "ok",
			Service: getServiceName(),
		}); err != nil {
			// reqLogger.Error("failed to encode health",
			// 	"status_code", http.StatusInternalServerError,
			// 	"error", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
		}

		// reqLogger.Info("health response sent")
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Info becomes the Minimum level to log , disables Debug level basically
	}))

	// http.HandleFunc("GET /health", healthHandler(logger))
	http.HandleFunc("GET /health", healthHandler())

	logger.Info("server starting", "addr", ":8080")
	_ = http.ListenAndServe(":8080", nil)
}
