# 01-go-core/ex04 — Custom error types + fmt.Stringer

## Goal

Practicar los patrones de error de Go más allá de `errors.New` y `fmt.Errorf`:
custom error types, la cadena de `Unwrap`, `errors.Is`, `errors.As`, y la interface `fmt.Stringer`.

**Packages principales:** `errors`, `fmt`

## Prompt

El catalog service necesita comunicar fallos con precisión para que los handlers HTTP puedan mapearlos al status code correcto sin inspeccionar strings de error.

Implementa los tipos y funciones marcados con `// TODO: implement` en `main.go`.

## Expected behavior

| Función/Método | Comportamiento |
|---|---|
| `NotFoundError.Error()` | `"product not found: <sku>"` |
| `NotFoundError.Unwrap()` | retorna `ErrNotFound` |
| `ValidationError.Error()` | `"validation: <field>: <reason>"` |
| `ValidationError.Unwrap()` | retorna `ErrInvalidInput` |
| `Product.String()` | contiene name, sku y price |
| `CatalogStore.Save` | `*ValidationError` si SKU o Name vacíos; `ErrConflict` si SKU duplicado |
| `CatalogStore.GetBySKU` | `*NotFoundError` si no existe |
| `HTTPStatusFor` | 200/400/404/409 usando `errors.Is`; nunca inspeccionando strings |

## Notes

- Usa `errors.Is` en `HTTPStatusFor` — no type switches, no string comparisons.
- El `Unwrap()` es lo que permite que `errors.Is` y `errors.As` traversar la cadena.
- `fmt.Stringer` se activa cuando usas `%v` o `%s` en `fmt.Sprintf`.
- `errors.As` extrae el tipo concreto incluso cuando el error está wrapped con `fmt.Errorf("%w", ...)`.

## Run tests

```bash
go test ./01-go-core/ex04/...
```
