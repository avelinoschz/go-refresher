# Concurrency Ex03: Context Cancellation

Goal:

- practice `context.Context`
- remember how to stop work when the caller no longer wants to wait

Prompt:

Complete `SlowFetcher.Fetch`.

Expected behavior:

- if the context is canceled before the delay finishes, return `context.Canceled` or `context.DeadlineExceeded`
- if it finishes in time, return the value

Hint:

- use `select`
- combine `time.After` with `<-ctx.Done()`

Command:

```bash
go run ./go-refresher/04-cache-concurrency/ex03
```

## Run tests

```bash
go test -race ./go-refresher/04-cache-concurrency/ex03
```
