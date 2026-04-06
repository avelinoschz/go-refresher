package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONWritesStatusAndBody(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()

	err := writeJSON(rec, http.StatusCreated, map[string]string{
		"status": "ok",
	})
	if err != nil {
		t.Fatalf("write json: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusCreated)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("unexpected response body: got %v", body)
	}
}

func TestWriteErrorWritesConsistentJSONBody(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()

	writeError(rec, http.StatusBadRequest, "invalid request")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusBadRequest)
	}

	var body ErrorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	if body.Error != "invalid request" {
		t.Fatalf("unexpected error response: got %+v", body)
	}
}

func TestHealthHandlerReturnsExpectedPayload(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	want := map[string]string{
		"status":  "ok",
		"service": "catalog-api",
	}

	if body["status"] != want["status"] || body["service"] != want["service"] {
		t.Fatalf("unexpected health response: got %v want %v", body, want)
	}
}
