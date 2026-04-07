# 02-http-json/ex04 — Custom JSON marshaling

## Goal

Go deeper into `encoding/json` beyond basic tags: custom marshalers, `omitempty`, the `-` tag, and streaming with `json.Decoder`/`json.Encoder`.

**Main packages:** `encoding/json`, `fmt`

## Prompt

The catalog API exposes products to two audiences: admins (all fields, including cost) and the public storefront (price and availability; cost must never be leaked).

Additionally, prices are stored internally in cents (`CentPrice`) but must be serialized as a decimal string in JSON (`"12.99"`).

Implement the types and functions marked with `// TODO: implement` in `main.go`.

## Expected behavior

| Type/Function | Behavior |
|---|---|
| `CentPrice.MarshalJSON` | `CentPrice(1299)` → `"12.99"` (JSON string with quotes) |
| `CentPrice.UnmarshalJSON` | `"12.99"` → `CentPrice(1299)` |
| `InternalProduct` with `CostCents=0` | `cost_cents` omitted from JSON |
| `InternalProduct` with `CostCents>0` | `cost_cents` included in JSON |
| `PublicProduct` | `cost_cents` never appears (tag `json:"-"`) |
| `ToPublic` | converts correctly; `InStock` = `Available` |
| `DecodeProduct` | uses `json.Decoder`, not `json.Unmarshal(io.ReadAll(...))` |
| `EncodePublicProduct` | uses `json.Encoder`, not `json.Marshal` + `w.Write` |

## Notes

- `omitempty` omits the field when its value is the zero value of the type (`0`, `false`, `""`).
- The `json:"-"` tag excludes the field entirely — it does not appear in marshal or unmarshal.
- For `MarshalJSON`, you must produce valid JSON bytes (the string including its quotes).
- `json.Encoder` appends `\n` at the end — tests account for this.
- Use `fmt.Sprintf("%.2f", float64(p)/100)` to format cents as a decimal.

## Run tests

```bash
go test ./02-http-json/ex04/...
```
