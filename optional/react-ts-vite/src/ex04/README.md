# Exercise 04: Custom Hook

## Goal

Extract the fetch logic from ex03 into a reusable `useCatalog(sku)` hook.

## Task

In `useCatalog.ts`, implement:

- `useCatalog(sku: string): UseCatalogResult`
- The hook manages `loading`, `data`, and `error` state
- Fetches `/api/catalog?sku=<sku>` whenever `sku` changes (skip empty string)
- Returns `{ loading, data, error }`

The `UseCatalogResult` type is already defined — implement the hook body.

## Run

```bash
npm test
```

## Reset

```bash
cp _original/useCatalog.ts useCatalog.ts
```

## Concepts

- Custom hooks: any function starting with `use` that calls other hooks
- Explicit return type annotations
- Reusing hook logic across multiple components
