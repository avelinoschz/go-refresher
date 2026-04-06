package main

import (
	"context"
	"fmt"
	"time"
)

type SlowFetcher struct {
	delay time.Duration
}

func (f SlowFetcher) Fetch(ctx context.Context, key string) (string, error) {
	_ = key
	// TODO: implement
	return "", nil
}

func main() {
	fetcher := SlowFetcher{delay: 2 * time.Second}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	value, err := fetcher.Fetch(ctx, "catalog:HAMMER-001")
	fmt.Println("value:", value)
	fmt.Println("error:", err)
}
