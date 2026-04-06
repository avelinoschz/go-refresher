# Cache Ex02: Fetch with Fallback

Goal:

- simulate a service that checks the cache first and then falls back to a slower source
- practice `cache hit` and `cache miss` flow

Prompt:

Complete `CatalogService.GetProductName`.

Expected behavior:

- if the value exists in the cache, return it immediately
- if it does not exist, use `ProductFetcher`
- if the fetch succeeds, store the result in the cache
- if the fetch fails, return an error

Goal:

- make the flow very clear when someone reads it

Command:

```bash
go run ./go-refresher/04-cache-concurrency/ex02
```

## Run tests

```bash
go test -race ./go-refresher/04-cache-concurrency/ex02
```
