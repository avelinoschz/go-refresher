# Testing Ex02: Handler Tests with httptest

Goal:

- test a handler without starting a real server
- validate the status code and JSON body

Prompt:

Write tests for `catalogHandler`.

Minimum cases:

- `400` when `sku` is missing
- `404` when the product does not exist
- `200` with valid JSON when the product exists

Tip:

- use `httptest.NewRequest`
- use `httptest.NewRecorder`

Command:

```bash
go run ./go-refresher/03-testing/ex02
```

## Run tests

```bash
go test ./go-refresher/03-testing/ex02
```
