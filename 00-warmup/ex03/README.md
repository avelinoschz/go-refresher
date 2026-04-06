# Warmup Ex03: Pointer Receivers + Interface

Goal:

- practice pointer receivers
- refresh small interfaces
- separate state logic from console output

Prompt:

Complete `TaskList` so that:

- `Add` adds a non-empty task
- `Complete` marks a task as completed by title
- `Stats` returns total and completed counts
- `ConsolePrinter` prints the summary using the `StatsReporter` interface

Rules:

- if the title is empty, return an error
- if `Complete` cannot find the task, return an error
- avoid overengineering

Command:

```bash
go run ./go-refresher/00-warmup/ex03
```

## Run tests

```bash
go test ./go-refresher/00-warmup/ex03
```
