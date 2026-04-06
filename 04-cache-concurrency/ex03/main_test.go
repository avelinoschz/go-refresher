package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestSlowFetcherReturnsValueWhenContextStaysActive(t *testing.T) {
	t.Parallel()

	fetcher := SlowFetcher{delay: 10 * time.Millisecond}

	got, err := fetcher.Fetch(context.Background(), "catalog:HAMMER-001")
	if err != nil {
		t.Fatalf("fetch value: %v", err)
	}

	if got == "" {
		t.Fatal("expected non-empty value")
	}
}

func TestSlowFetcherReturnsContextErrorWhenCanceled(t *testing.T) {
	t.Parallel()

	fetcher := SlowFetcher{delay: 100 * time.Millisecond}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := fetcher.Fetch(ctx, "catalog:HAMMER-001")
	if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected error: got %v", err)
	}
}

func TestSlowFetcherReturnsDeadlineExceeded(t *testing.T) {
	t.Parallel()

	fetcher := SlowFetcher{delay: 100 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := fetcher.Fetch(ctx, "catalog:HAMMER-001")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected error: got %v want %v", err, context.DeadlineExceeded)
	}
}
