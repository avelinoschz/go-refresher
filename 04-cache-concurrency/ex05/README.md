# 04-cache-concurrency/ex05 — HTTP client con context + error wrapping

## Goal

Practicar el lado cliente de `net/http`: hacer requests con `context`, manejar timeouts, leer el body de forma segura, y construir una cadena de errores rica con `errors` y `fmt`.

**Packages principales:** `net/http`, `context`, `errors`

## Prompt

El catalog service necesita obtener precios desde un upstream externo. Implementa un `PricingClient` que respete el contexto del caller, capture errores de upstream con suficiente contexto para hacer debugging, y exponga tipos de error que los callers puedan inspeccionar con `errors.Is` / `errors.As`.

Los tests usan `httptest.NewServer` para simular el upstream — no se hace ninguna llamada de red real.

Implementa los tipos y funciones marcados con `// TODO: implement` en `main.go`.

## Expected behavior

| Función/Método | Comportamiento |
|---|---|
| `UpstreamError.Error()` | `"upstream error: status <code>: <body>"` |
| `UpstreamError.Unwrap()` | retorna `ErrUpstream` |
| `FetchPrice` — 2xx | decodifica el JSON y retorna `PriceResponse` |
| `FetchPrice` — non-2xx | retorna `*UpstreamError` con el status code y primeros 256 bytes del body |
| `FetchPrice` — ctx cancelado | retorna error que wrappea `context.Canceled` |
| `FetchPrice` — timeout del client | retorna error que wrappea `context.DeadlineExceeded` |

## Notes

- Usa `http.NewRequestWithContext(ctx, ...)` — nunca `http.NewRequest`.
- Siempre cierra `resp.Body` con `defer resp.Body.Close()`.
- Usa `io.LimitReader(resp.Body, 1<<20)` para acotar el body antes de leerlo.
- El timeout del `http.Client` y el deadline del context son independientes — el que dispara primero gana.
- Cuando el contexto está cancelado antes del request, `http.Client.Do` retorna un error que ya wrappea `context.Canceled` — no necesitas wrapearlo manualmente.
- Para errores de decode usa: `fmt.Errorf("decode pricing response: %w", err)`.

## Run tests

```bash
go test ./04-cache-concurrency/ex05/...
```
