package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// stubHandler is a minimal handler used in middleware tests.
// It writes the given status code and records whether it was called.
func stubHandler(code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	}
}

// ---- RequestIDFromContext ----

func TestRequestIDFromContext_Missing(t *testing.T) {
	t.Parallel()
	id := RequestIDFromContext(t.Context())
	if id != "" {
		t.Fatalf("expected empty string, got %q", id)
	}
}

// ---- WithRequestID ----

func TestWithRequestID_SetsHeader(t *testing.T) {
	t.Parallel()
	handler := WithRequestID(stubHandler(http.StatusOK))
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	id := rec.Header().Get("X-Request-ID")
	if id == "" {
		t.Fatal("X-Request-ID header should be set")
	}
}

func TestWithRequestID_IDAvailableInContext(t *testing.T) {
	t.Parallel()
	var capturedID string
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedID = RequestIDFromContext(r.Context())
	})

	handler := WithRequestID(inner)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if capturedID == "" {
		t.Fatal("request ID should be accessible from context inside handler")
	}
	if rec.Header().Get("X-Request-ID") != capturedID {
		t.Fatalf("header %q != context value %q", rec.Header().Get("X-Request-ID"), capturedID)
	}
}

func TestWithRequestID_UniquePerRequest(t *testing.T) {
	t.Parallel()
	handler := WithRequestID(stubHandler(http.StatusOK))
	ids := make(map[string]struct{})
	for range 5 {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		id := rec.Header().Get("X-Request-ID")
		if _, dup := ids[id]; dup {
			t.Fatalf("duplicate request ID %q", id)
		}
		ids[id] = struct{}{}
	}
}

// ---- WithLogger ----

func TestWithLogger_WritesOneLine(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	// Wrap with WithRequestID first so the logger can read the ID from context.
	handler := WithRequestID(WithLogger(&buf, stubHandler(http.StatusOK)))

	req := httptest.NewRequest(http.MethodGet, "/products/HAMMER-001", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 1 {
		t.Fatalf("expected 1 log line, got %d: %q", len(lines), buf.String())
	}
}

func TestWithLogger_LineContainsMethodAndPath(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	handler := WithRequestID(WithLogger(&buf, stubHandler(http.StatusCreated)))

	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	line := buf.String()
	for _, sub := range []string{"POST", "/products"} {
		if !strings.Contains(line, sub) {
			t.Fatalf("log line %q does not contain %q", line, sub)
		}
	}
}

func TestWithLogger_LineContainsRequestID(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	var capturedID string
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedID = RequestIDFromContext(r.Context())
		w.WriteHeader(http.StatusOK)
	})
	handler := WithRequestID(WithLogger(&buf, inner))

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)

	if capturedID == "" {
		t.Fatal("could not capture request ID")
	}
	if !strings.Contains(buf.String(), capturedID) {
		t.Fatalf("log line %q does not contain request ID %q", buf.String(), capturedID)
	}
}

// ---- NewRouter ----

func TestNewRouter_GetProducts(t *testing.T) {
	t.Parallel()
	called := false
	getProduct := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})
	mux := NewRouter(getProduct, stubHandler(http.StatusCreated), &bytes.Buffer{})

	req := httptest.NewRequest(http.MethodGet, "/products/HAMMER-001", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if !called {
		t.Fatal("GET /products/{sku} should have reached the getProduct handler")
	}
}

func TestNewRouter_PostProducts(t *testing.T) {
	t.Parallel()
	called := false
	createProduct := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusCreated)
	})
	mux := NewRouter(stubHandler(http.StatusOK), createProduct, &bytes.Buffer{})

	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if !called {
		t.Fatal("POST /products should have reached the createProduct handler")
	}
}

func TestNewRouter_HealthCheck(t *testing.T) {
	t.Parallel()
	mux := NewRouter(stubHandler(http.StatusOK), stubHandler(http.StatusCreated), &bytes.Buffer{})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /health status: got %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "ok") {
		t.Fatalf("GET /health body should contain 'ok', got: %s", rec.Body.String())
	}
}

func TestNewRouter_UnknownRoute(t *testing.T) {
	t.Parallel()
	mux := NewRouter(stubHandler(http.StatusOK), stubHandler(http.StatusCreated), &bytes.Buffer{})

	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("unknown route status: got %d, want 404", rec.Code)
	}
}

func TestNewRouter_MiddlewareApplied(t *testing.T) {
	t.Parallel()
	var logBuf bytes.Buffer
	mux := NewRouter(stubHandler(http.StatusOK), stubHandler(http.StatusCreated), &logBuf)

	req := httptest.NewRequest(http.MethodGet, "/products/NAILS-050", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	// Middleware should set X-Request-ID header.
	if rec.Header().Get("X-Request-ID") == "" {
		t.Fatal("middleware should set X-Request-ID header on routed requests")
	}
	// Logger should write at least one line.
	if logBuf.Len() == 0 {
		t.Fatal("middleware should log the request")
	}
}
