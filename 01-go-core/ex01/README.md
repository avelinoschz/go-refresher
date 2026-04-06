# Go Core Ex01: Parse and Validate

Goal:

- practice parsing text into structs
- split validation into small functions
- write useful errors

Prompt:

Implement `ParseServices`.

Input format:

- one line per service
- each line uses the format `name,team,replicas`

Example:

```text
catalog,platform,3
checkout,commerce,2
```

Rules:

- `name` and `team` are required
- `replicas` must be a positive integer
- if a line is invalid, return an error and stop
- ignore empty lines

Done criteria:

- clear parsing
- errors with enough context
- no external packages

Command:

```bash
go run ./go-refresher/01-go-core/ex01
```

## Run tests

```bash
go test ./go-refresher/01-go-core/ex01
```
