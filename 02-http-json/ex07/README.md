# HTTP Ex07: Server Bootstrap + ServeMux + Middleware

Goal:

- practice building an `*http.Server`
- wire routes using `http.NewServeMux`
- compose middleware around the mux
- implement graceful shutdown with `signal.NotifyContext`
- keep the handlers simple and infrastructure-focused

Prompt:

Complete the server setup in `main.go`.

Implement:

- `chain`
- `withServerHeader`
- `withRecovery`
- `newHandler`
- `newServer`
- `serve`

Expected routes:

- `GET /health` → `200` with JSON `{ "status": "ok" }`
- `GET /ready` → `200` with JSON `{ "ready": true }`
- `POST /echo` → reads JSON body and writes it back unchanged
- `GET /panic` → intentionally panics so the recovery middleware can turn it into a `500` JSON error

Rules:

- use `http.NewServeMux()`
- use method-based stdlib patterns such as `"GET /health"`
- set `X-Server: go-refresher` on all responses through middleware
- if a handler panics, return `500` with a JSON error body
- configure an `http.Server` with explicit timeouts
- use `signal.NotifyContext` and `server.Shutdown` with a timeout for graceful shutdown

Command:

```bash
go run ./02-http-json/ex07/main.go
```

## Run tests

```bash
go test ./02-http-json/ex07
```
