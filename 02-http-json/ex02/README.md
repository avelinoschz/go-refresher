# HTTP Ex02: POST + Decode + Validate

Goal:

- practice JSON decoding
- validate input
- respond with `201`, `400`, and `409`

Prompt:

Complete `POST /catalog`.

Expected body:

```json
{
  "sku": "DRILL-002",
  "name": "Drill",
  "price": 80
}
```

Rules:

- `sku` and `name` are required
- `price` must be greater than 0
- if the `sku` already exists, respond with `409`

Command:

```bash
go run ./02-http-json/ex02/main.go
```

## Run tests

```bash
go test ./02-http-json/ex02
```
