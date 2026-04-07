# HTTP Ex06: Path Params + Strict JSON Body

Goal:

- practice `r.PathValue` with the stdlib router
- decode JSON defensively with `json.Decoder`
- reject unknown fields and oversized request bodies

Prompt:

Complete `PUT /catalog/{sku}/price`.

Expected body:

```json
{
  "price": 30
}
```

Rules:

- use `r.PathValue("sku")` to extract the product SKU from the route
- if the SKU does not exist, respond with `404`
- if `price` is less than or equal to 0, respond with `400`
- if the JSON is invalid or contains unknown fields, respond with `400`
- if the body exceeds the allowed size, respond with `413`
- if everything is valid, update the product price and respond with `200` and the updated product JSON

Command:

```bash
go run ./02-http-json/ex06/main.go
```

## Run tests

```bash
go test ./02-http-json/ex06
```
