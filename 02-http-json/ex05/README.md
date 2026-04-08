# 02-http-json/ex05 — Middleware chain + context.WithValue

## Goal

Practice HTTP middleware composition and the use of `context.WithValue` to propagate cross-cutting values (such as a request ID) through the call chain.

**Main packages:** `net/http`, `context`

## Prompt

The team needs every request to have a unique traceable ID in logs and in the response header. Implement two middlewares and a router that applies them.

Implement the types and functions marked with `// TODO: implement` in `main.go`.

## Expected behavior

| Function | Behavior |
| - | - |
| `RequestIDFromContext` | extracts the ID from the context; `""` if not present |
| `WithRequestID` | generates a unique ID, stores it in ctx, writes `X-Request-ID` header |
| `WithLogger` | writes `"METHOD PATH id=<id> status=<code>\n"` to `out` per request |
| `statusRecorder.WriteHeader` | captures the code without losing the real response |
| `NewRouter` | registers 3 routes; applies both middlewares to `getProduct` and `createProduct` |

## Notes

- `context.WithValue(ctx, requestIDKey, id)` — use the `contextKey` type to avoid collisions with other packages.
- `ctx.Value(requestIDKey)` returns `any`; you must type-assert to `string`.
- The logger needs the status code *after* the handler responds → use `statusRecorder` to intercept `WriteHeader`.
- If the handler does not call `WriteHeader` explicitly, the default status is 200 — initialize `statusRecorder.code = http.StatusOK`.
- Middlewares compose outside-in: `WithRequestID(WithLogger(handler))` — the ID is already in the context when it reaches the logger.

## Run tests

```bash
go test ./02-http-json/ex05/...
```
