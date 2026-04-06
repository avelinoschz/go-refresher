# HTTP Ex01: GET + Query Params + JSON

Goal:

- implement a `GET` handler
- read query params
- respond with JSON

Prompt:

Complete the `GET /catalog?sku=...` endpoint.

Expected behavior:

- if `sku` is missing, respond with `400`
- if it does not exist, respond with `404`
- if it exists, respond with `200` and the product JSON

Notes:

- use the in-memory catalog that is already included
- you do not need a database
- keep the response simple and clear

Command:

```bash
go run ./02-http-json/ex01/main.go
```

## Run tests

```bash
go test ./02-http-json/ex01
```
