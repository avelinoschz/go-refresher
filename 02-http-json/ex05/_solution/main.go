package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func RequestIDFromContext(ctx context.Context) string {
	id, _ := ctx.Value(requestIDKey).(string)
	return id
}

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := fmt.Sprintf("%08x", rand.Uint32())
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.code = code
	r.ResponseWriter.WriteHeader(code)
}

func WithLogger(out io.Writer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, code: http.StatusOK}
		next.ServeHTTP(rec, r)
		id := RequestIDFromContext(r.Context())
		fmt.Fprintf(out, "%s %s id=%s status=%d\n", r.Method, r.URL.Path, id, rec.code)
	})
}

func NewRouter(getProduct, createProduct http.HandlerFunc, logOut io.Writer) *http.ServeMux {
	mux := http.NewServeMux()

	wrap := func(h http.Handler) http.Handler {
		return WithRequestID(WithLogger(logOut, h))
	}

	mux.Handle("GET /products/{sku}", wrap(getProduct))
	mux.Handle("POST /products", wrap(createProduct))
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	return mux
}

func main() {}
