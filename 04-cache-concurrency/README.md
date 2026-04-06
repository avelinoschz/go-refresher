# 04 Cache and Concurrency

This block is about two things:

- solving classic backend problems with in-memory state
- getting comfortable with basic concurrency when it actually helps

Priorities:

- `map`
- `sync.RWMutex`
- `context.Context`
- `goroutine`
- `channel`
- `select`

Rule:

- aim for solutions that are clean and readable
- avoid fancy patterns if a simpler solution works

## Run tests

```bash
go test -race ./go-refresher/04-cache-concurrency/...
```
