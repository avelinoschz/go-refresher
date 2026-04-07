# 04-cache-concurrency/ex05 — HTTP client with context + error wrapping

## Goal

Practice the client side of `net/http`: making requests with `context`, handling timeouts, reading the body safely, and building a rich error chain with `errors` and `fmt`.

**Main packages:** `net/http`, `context`, `errors`

## Prompt

The catalog service needs to fetch prices from an external upstream. Implement a `PricingClient` that respects the caller's context, captures upstream errors with enough context for debugging, and exposes error types that callers can inspect with `errors.Is` / `errors.As`.

Tests use `httptest.NewServer` to simulate the upstream — no real network calls are made.

Implement the types and functions marked with `// TODO: implement` in `main.go`.

## Expected behavior

| Function/Method | Behavior |
|---|---|
| `UpstreamError.Error()` | `"upstream error: status <code>: <body>"` |
| `UpstreamError.Unwrap()` | returns `ErrUpstream` |
| `FetchPrice` — 2xx | decodes the JSON and returns `PriceResponse` |
| `FetchPrice` — non-2xx | returns `*UpstreamError` with the status code and first 256 bytes of the body |
| `FetchPrice` — ctx canceled | returns error that wraps `context.Canceled` |
| `FetchPrice` — client timeout | returns error that wraps `context.DeadlineExceeded` |

## Notes

- Use `http.NewRequestWithContext(ctx, ...)` — never `http.NewRequest`.
- Always close `resp.Body` with `defer resp.Body.Close()`.
- Use `io.LimitReader(resp.Body, 1<<20)` to cap the body before reading it.
- The `http.Client` timeout and the context deadline are independent — whichever fires first wins.
- When the context is canceled before the request, `http.Client.Do` returns an error that already wraps `context.Canceled` — you do not need to wrap it manually.
- For decode errors use: `fmt.Errorf("decode pricing response: %w", err)`.

## Run tests

```bash
go test ./04-cache-concurrency/ex05/...
```
