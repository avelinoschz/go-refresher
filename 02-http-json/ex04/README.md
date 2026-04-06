# 02-http-json/ex04 — Custom JSON marshaling

## Goal

Profundizar en `encoding/json` más allá de los tags básicos: marshalers custom, `omitempty`, el tag `-`, y streaming con `json.Decoder`/`json.Encoder`.

**Packages principales:** `encoding/json`, `fmt`

## Prompt

El catalog API expone productos a dos audiencias: admins (todos los campos, incluyendo costo) y el storefront público (precio y disponibilidad; el costo nunca debe filtrarse).

Adicionalmente, el precio se almacena en centavos internamente (`CentPrice`) pero debe serializarse como decimal string en JSON (`"12.99"`).

Implementa los tipos y funciones marcados con `// TODO: implement` en `main.go`.

## Expected behavior

| Tipo/Función | Comportamiento |
|---|---|
| `CentPrice.MarshalJSON` | `CentPrice(1299)` → `"12.99"` (string JSON con comillas) |
| `CentPrice.UnmarshalJSON` | `"12.99"` → `CentPrice(1299)` |
| `InternalProduct` con `CostCents=0` | `cost_cents` omitido del JSON |
| `InternalProduct` con `CostCents>0` | `cost_cents` incluido en el JSON |
| `PublicProduct` | `cost_cents` nunca aparece (tag `json:"-"`) |
| `ToPublic` | convierte correctamente; `InStock` = `Available` |
| `DecodeProduct` | usa `json.Decoder`, no `json.Unmarshal(io.ReadAll(...))` |
| `EncodePublicProduct` | usa `json.Encoder`, no `json.Marshal` + `w.Write` |

## Notes

- `omitempty` omite el campo cuando su valor es el zero value del tipo (`0`, `false`, `""`).
- El tag `json:"-"` excluye el campo por completo — no aparece ni en marshal ni en unmarshal.
- Para `MarshalJSON`, debes producir bytes JSON válidos (el string con sus comillas incluidas).
- `json.Encoder` agrega `\n` al final — los tests lo tienen en cuenta.
- Usa `fmt.Sprintf("%.2f", float64(p)/100)` para formatear centavos como decimal.

## Run tests

```bash
go test ./02-http-json/ex04/...
```
