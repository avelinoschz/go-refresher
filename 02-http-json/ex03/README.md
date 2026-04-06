# HTTP Ex03: Consistent Error Handling

Goal:

- avoid repeating too much code across handlers
- create small helpers for JSON responses

Prompt:

Complete:

- `writeJSON`
- `writeError`
- the `GET /health` handler

Response for `GET /health`:

```json
{
  "status": "ok",
  "service": "catalog-api"
}
```

When a serialization error happens:

- respond with `500`
- return a consistent error JSON body

Command:

```bash
go run ./go-refresher/02-http-json/ex03
```

## Run tests

```bash
go test ./go-refresher/02-http-json/ex03
```
