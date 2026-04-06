# Concurrency Ex04: Basic Fan-In

Goal:

- coordinate two concurrent lookups
- use channels without making the solution more complex than necessary

Prompt:

You have two provider functions:

- one fetches price
- one fetches stock

Complete `LoadSnapshot` so that it:

- runs both at the same time
- waits for both results
- combines everything into a `ProductSnapshot`
- returns an error if either call fails

Restriction:

- do not use `sync.WaitGroup`
- solve it with channels because that is the skill we want to practice here

Command:

```bash
go run ./04-cache-concurrency/ex04/main.go
```

## Run tests

```bash
go test -race ./04-cache-concurrency/ex04
```
