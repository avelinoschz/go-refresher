# Exercise 02: Cache-Aside with Redis

## Goal

Implement a `CatalogService` that uses Redis as a cache layer in front of a simulated product store.

## Task

Implement `GetProductName(ctx, sku) (string, error)`:

- check Redis first (key format: `catalog:<sku>`)
- on miss: fetch from the store, store the result in Redis with a 30-second TTL
- propagate errors from both Redis and the store

## Run

```bash
make redis-up
go test ./05-redis/ex02/
```

## Notes

- The `Cache` interface is already defined — implement it or reuse your solution from ex01.
- Tests verify both the cache hit path and the cache miss + populate path.
