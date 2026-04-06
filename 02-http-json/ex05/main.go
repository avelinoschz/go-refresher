package main

import (
	"context"
	"io"
	"net/http"
)

// contextKey is an unexported type used as context key to avoid collisions
// with other packages that also use context.WithValue.
type contextKey string

const requestIDKey contextKey = "request_id"

// RequestIDFromContext extracts the request ID injected by WithRequestID.
// Returns "" if no request ID is present in the context.
func RequestIDFromContext(ctx context.Context) string {
	_ = requestIDKey
	// TODO: implement — ctx.Value(requestIDKey)
	return ""
}

// WithRequestID is a middleware that:
//  1. Generates a short unique ID for each request (e.g. random hex or counter).
//  2. Stores it in the request context using requestIDKey.
//  3. Sets the "X-Request-ID" response header before calling next.
func WithRequestID(next http.Handler) http.Handler {
	_ = context.WithValue
	// TODO: implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// statusRecorder wraps http.ResponseWriter to capture the written status code.
// It is used by WithLogger to log the response code after the handler runs.
type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (r *statusRecorder) WriteHeader(code int) {
	// TODO: implement — store code, then delegate to the embedded ResponseWriter
}

// WithLogger is a middleware that writes one log line per request to w:
//
//	"<METHOD> <PATH> id=<request_id> status=<code>\n"
//
// It must call next.ServeHTTP first so it can capture the status code.
// Use statusRecorder to intercept WriteHeader.
// The request ID comes from RequestIDFromContext.
func WithLogger(out io.Writer, next http.Handler) http.Handler {
	_ = io.Writer(out)
	// TODO: implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// NewRouter builds a *http.ServeMux with:
//   - GET /products/{sku}  → getProduct
//   - POST /products       → createProduct
//   - GET /health          → inline handler returning {"status":"ok"}
//
// Both getProduct and createProduct are wrapped with WithRequestID and
// WithLogger (in that order, outermost first).
func NewRouter(getProduct, createProduct http.HandlerFunc, logOut io.Writer) *http.ServeMux {
	// TODO: implement
	return http.NewServeMux()
}

func main() {}
