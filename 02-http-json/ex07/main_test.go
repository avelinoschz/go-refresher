package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewHandlerHealth(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}
	if rec.Header().Get("X-Server") != "go-refresher" {
		t.Fatalf("unexpected X-Server header: %q", rec.Header().Get("X-Server"))
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body["status"] != "ok" {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestNewHandlerReady(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}

	var body map[string]bool
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if !body["ready"] {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestNewHandlerEcho(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(`{"message":"hello"}`))
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body["message"] != "hello" {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestNewHandlerEchoMethodNotAllowed(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/echo", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}

func TestRecoveryMiddlewareReturnsJSON500(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusInternalServerError)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body["error"] != "internal server error" {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestNewServerConfig(t *testing.T) {
	t.Parallel()

	handler := newHandler()
	server := newServer(":9090", handler)

	if server.Addr != ":9090" {
		t.Fatalf("unexpected addr: got %q want %q", server.Addr, ":9090")
	}
	if server.Handler == nil {
		t.Fatal("expected handler to be configured")
	}
	if server.ReadHeaderTimeout != 2*time.Second {
		t.Fatalf("unexpected ReadHeaderTimeout: got %v want %v", server.ReadHeaderTimeout, 2*time.Second)
	}
	if server.WriteTimeout != 5*time.Second {
		t.Fatalf("unexpected WriteTimeout: got %v want %v", server.WriteTimeout, 5*time.Second)
	}
	if server.IdleTimeout != 30*time.Second {
		t.Fatalf("unexpected IdleTimeout: got %v want %v", server.IdleTimeout, 30*time.Second)
	}
}

func TestServeGracefulShutdown(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	server := newServer("127.0.0.1:0", newHandler())

	done := make(chan error, 1)
	go func() {
		done <- serve(ctx, server)
	}()

	time.Sleep(20 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("serve returned error: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("serve did not shut down in time")
	}
}
