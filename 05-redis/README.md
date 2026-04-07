# 05-redis

This block is the natural progression from `04-cache-concurrency`: it replaces the in-memory cache with real Redis using `go-redis/v9`.

## Exercises

| # | Exercise | Key concept |
| - | - | - |
| ex01 | Redis Client Wrapper | `GET`/`SET`, handling `redis.Nil` as a miss |
| ex02 | Cache-Aside with Redis | key format, 30s TTL, cache hit/miss |
| ex03 | Graceful Degradation | `context.WithTimeout`, fallback to store |

## Requirements

- Docker (to run Redis)
- Go 1.22+

## Setup

```bash
# Start Redis
make redis-up

# Validate the connection before running tests
redis-cli PING
# Should respond: PONG

# Stop Redis when done
make redis-down
```

## Run tests

```bash
# One exercise at a time
go test -race ./05-redis/ex01/
go test -race ./05-redis/ex02/
go test -race ./05-redis/ex03/

# Full block
go test -race ./05-redis/...
```

> Tests require Redis running on `localhost:6379`. If they fail with "connection refused", verify that `make redis-up` started the container.
