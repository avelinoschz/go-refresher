# Testing Ex01: Table-Driven Tests

Goal:

- practice tests for pure functions
- get comfortable with the table-driven style

Prompt:

Implement `NormalizeTags` and then write tests.

Expected behavior:

- surrounding spaces are trimmed
- everything is converted to lowercase
- empty strings are ignored
- the result must not contain duplicates

Minimum cases:

- empty input
- tags with uppercase letters and spaces
- duplicate tags

Command:

```bash
go run ./go-refresher/03-testing/ex01
```

## Run tests

```bash
go test ./go-refresher/03-testing/ex01
```
