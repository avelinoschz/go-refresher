package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// newTestServer is a helper that starts a test HTTP server.
// handler receives each request; the caller is responsible for closing the server.
func newTestServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return srv
}

func TestFetchPrice_Success(t *testing.T) {
	t.Parallel()
	want := PriceResponse{SKU: "HAMMER-001", PriceCents: 2500}
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	})

	client := NewPricingClient(srv.URL, 0)
	got, err := client.FetchPrice(context.Background(), "HAMMER-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.SKU != want.SKU || got.PriceCents != want.PriceCents {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestFetchPrice_UpstreamError_404(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})

	client := NewPricingClient(srv.URL, 0)
	_, err := client.FetchPrice(context.Background(), "MISSING")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrUpstream) {
		t.Fatalf("errors.Is(err, ErrUpstream) should be true, got: %v", err)
	}
}

func TestFetchPrice_UpstreamError_StatusCode(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	})

	client := NewPricingClient(srv.URL, 0)
	_, err := client.FetchPrice(context.Background(), "X")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var ue *UpstreamError
	if !errors.As(err, &ue) {
		t.Fatalf("expected *UpstreamError, got %T: %v", err, err)
	}
	if ue.StatusCode != http.StatusInternalServerError {
		t.Fatalf("StatusCode: got %d, want %d", ue.StatusCode, http.StatusInternalServerError)
	}
}

func TestFetchPrice_UpstreamError_BodyCaptured(t *testing.T) {
	t.Parallel()
	body := "something went wrong upstream"
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(body))
	})

	client := NewPricingClient(srv.URL, 0)
	_, err := client.FetchPrice(context.Background(), "X")

	var ue *UpstreamError
	if !errors.As(err, &ue) {
		t.Fatalf("expected *UpstreamError, got %T", err)
	}
	if ue.Body == "" {
		t.Fatal("UpstreamError.Body should contain the response body")
	}
}

func TestFetchPrice_ContextCancelled(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		// Simulate a slow server.
		<-r.Context().Done()
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately before making the request

	client := NewPricingClient(srv.URL, 5*time.Second)
	_, err := client.FetchPrice(ctx, "X")
	if err == nil {
		t.Fatal("expected error for cancelled context, got nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled in error chain, got: %v", err)
	}
}

func TestFetchPrice_ClientTimeout(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		// Never respond — force client timeout.
		<-r.Context().Done()
	})

	client := NewPricingClient(srv.URL, 50*time.Millisecond)
	_, err := client.FetchPrice(context.Background(), "X")
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded in error chain, got: %v", err)
	}
}

func TestUpstreamError_Error(t *testing.T) {
	t.Parallel()
	e := &UpstreamError{StatusCode: 503, Body: "service unavailable"}
	msg := e.Error()
	for _, sub := range []string{"503", "service unavailable"} {
		if len(msg) == 0 {
			t.Fatal("Error() returned empty string")
		}
		found := false
		for i := 0; i+len(sub) <= len(msg); i++ {
			if msg[i:i+len(sub)] == sub {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Error() = %q does not contain %q", msg, sub)
		}
	}
}

func TestUpstreamError_Unwrap(t *testing.T) {
	t.Parallel()
	e := &UpstreamError{StatusCode: 404}
	if !errors.Is(e, ErrUpstream) {
		t.Fatal("errors.Is(*UpstreamError, ErrUpstream) should be true")
	}
}
