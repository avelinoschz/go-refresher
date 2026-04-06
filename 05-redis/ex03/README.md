# Exercise 03: Graceful Degradation with Context Timeout

## Goal

Implement a `GetWithFallback` function that bounds the Redis call with a short context timeout and falls back to the store if Redis is unavailable or too slow.

## Task

Implement `GetWithFallback(ctx context.Context, sku string) (string, error)`:

- try Redis with a 100ms child context timeout
- if the Redis call fails for any reason (timeout, connection error), fall back to the store
- if the store also fails, return the error

## Run

```bash
make redis-up
go test ./05-redis/ex03/
```

## Notes

- Use `context.WithTimeout` to create the child context for the Redis call.
- A cancelled or timed-out context from the caller should still cancel the Redis child context.
- This pattern is common in production: Redis is best-effort; the store is the source of truth.
