# Go Refresher

This repository contains a structured Go skills refresher with a backend/platform focus.

Goals:

- regain speed writing Go from memory
- practice HTTP, JSON, tests, caching, and concurrency
- revisit platform-engineering concepts

Usage rules:

1. Move in order. Each block assumes the previous one.
2. `00-warmup/ex01-copy-by-hand` is the only fully solved exercise.
3. Everything after that uses starter code or small skeletons.
4. Try to solve each exercise on your own before bringing it back for review.
5. If you get stuck, read the exercise README before consulting external documentation.

Helpful commands:

```bash
go run ./00-warmup/ex01-copy-by-hand/main.go
go test ./...

# Redis exercises (05-redis) require a running Redis instance:
make redis-up
go test ./05-redis/...
make redis-down
```

## Redoing exercises

Each completed exercise keeps a blank copy of the implementation under `_original/` (ignored by `go test ./...`). Use the Makefile to reset or restore an exercise:

```bash
# Reset to blank state — backs up current solution as main.go.bak
make reset EX=00-warmup/ex01-copy-by-hand

# Restore the backed-up solution
make restore EX=00-warmup/ex01-copy-by-hand
```

Conventions:

- Write documentation and comments in English.
- Keep code identifiers in English.
- In backend exercises, aim for a consistent shape:
  - `handler`
  - `service`
  - `store` or `cache`
- Use the standard library by default.
- `05-redis` is the only block that uses external dependencies (`go-redis/v9`).
