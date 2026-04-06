# Exercise 01: Redis Client Wrapper

## Goal

Implement a `RedisCache` struct that wraps `*redis.Client` and exposes a clean interface for get and set operations.

## Task

Implement:

- `Get(ctx, key) (string, bool, error)` — return `false` when the key does not exist (`redis.Nil`), not an error
- `Set(ctx, key, value string, ttl time.Duration) error`

## Run

```bash
make redis-up
go test ./05-redis/ex01/
```

## Notes

- `redis.Nil` is the sentinel error returned by go-redis when a key is missing — treat it as a miss, not a failure.
- Always pass `ctx` to every Redis call.
