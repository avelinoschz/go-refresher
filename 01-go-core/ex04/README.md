# 01-go-core/ex04 — Custom error types + fmt.Stringer

## Goal

Practice Go error patterns beyond `errors.New` and `fmt.Errorf`:
custom error types, the `Unwrap` chain, `errors.Is`, `errors.As`, and the `fmt.Stringer` interface.

**Main packages:** `errors`, `fmt`

## Prompt

The catalog service needs to communicate failures precisely so HTTP handlers can map them to the correct status code without inspecting error strings.

Implement the types and functions marked with `// TODO: implement` in `main.go`.

## Expected behavior

| Function/Method | Behavior |
| - | - |
| `NotFoundError.Error()` | `"product not found: <sku>"` |
| `NotFoundError.Unwrap()` | returns `ErrNotFound` |
| `ValidationError.Error()` | `"validation: <field>: <reason>"` |
| `ValidationError.Unwrap()` | returns `ErrInvalidInput` |
| `Product.String()` | contains name, sku and price |
| `CatalogStore.Save` | `*ValidationError` if SKU or Name are empty; `ErrConflict` if SKU is duplicate |
| `CatalogStore.GetBySKU` | `*NotFoundError` if it does not exist |
| `HTTPStatusFor` | 200/400/404/409 using `errors.Is`; never inspecting strings |

## Notes

- Use `errors.Is` in `HTTPStatusFor` — no type switches, no string comparisons.
- `Unwrap()` is what allows `errors.Is` and `errors.As` to traverse the chain.
- `fmt.Stringer` is triggered when you use `%v` or `%s` in `fmt.Sprintf`.
- `errors.As` extracts the concrete type even when the error is wrapped with `fmt.Errorf("%w", ...)`.

## Run tests

```bash
go test ./01-go-core/ex04/...
```
