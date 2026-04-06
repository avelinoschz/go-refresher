# Optional: Redis with Docker

This block is optional and stays outside the critical path.

Goal:

- run Redis locally
- practice `GET` and `SET`
- connect a small version of the cache exercises

Suggested path:

1. start Redis with Docker Compose
2. validate the connection with `redis-cli`
3. add a minimal cache version using `go-redis/v9`

If we do it:

- reuse one of the `04-cache-concurrency` exercises
- do not replace the in-memory cache; only compare against it
