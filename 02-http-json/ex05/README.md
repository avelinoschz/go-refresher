# 02-http-json/ex05 — Middleware chain + context.WithValue

## Goal

Practicar la composición de middleware HTTP y el uso de `context.WithValue` para propagar valores cross-cutting (como un request ID) a través de la cadena de llamadas.

**Packages principales:** `net/http`, `context`

## Prompt

El equipo necesita que cada request tenga un ID único trazable en los logs y en el header de respuesta. Implementa dos middlewares y un router que los aplique.

Implementa los tipos y funciones marcados con `// TODO: implement` en `main.go`.

## Expected behavior

| Función | Comportamiento |
|---|---|
| `RequestIDFromContext` | extrae el ID del contexto; `""` si no hay |
| `WithRequestID` | genera ID único, lo guarda en ctx, escribe `X-Request-ID` header |
| `WithLogger` | escribe `"METHOD PATH id=<id> status=<code>\n"` a `out` por request |
| `statusRecorder.WriteHeader` | captura el código sin perder la respuesta real |
| `NewRouter` | registra 3 rutas; aplica ambos middlewares a `getProduct` y `createProduct` |

## Notes

- `context.WithValue(ctx, requestIDKey, id)` — usa el tipo `contextKey` para evitar colisiones con otros packages.
- `ctx.Value(requestIDKey)` retorna `any`; debes hacer type assertion a `string`.
- El logger necesita el status code *después* de que el handler responde → usa `statusRecorder` para interceptar `WriteHeader`.
- Si el handler no llama `WriteHeader` explícitamente, el status default es 200 — inicializa `statusRecorder.code = http.StatusOK`.
- Los middlewares se componen afuera hacia adentro: `WithRequestID(WithLogger(handler))` — el ID ya está en el contexto cuando llega al logger.

## Run tests

```bash
go test ./02-http-json/ex05/...
```
