# Cache Ex01: In-Memory Cache

Goal:

- build a concurrency-safe cache
- practice `sync.RWMutex`

Prompt:

Complete `MemoryCache`.

It needs:

- `Get(key string) (string, bool)`
- `Set(key, value string)`

Rules:

- use `map[string]string`
- protect access with `sync.RWMutex`
- do not add TTL yet

Done criteria:

- safe concurrent reads
- safe writes
- short and clear code

Command:

```bash
go run ./go-refresher/04-cache-concurrency/ex01
```

## Run tests

```bash
go test -race ./go-refresher/04-cache-concurrency/ex01
```
