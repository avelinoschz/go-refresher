# Exercise 03: Fetch with Loading / Error / Data States

## Goal

Implement `CatalogSearch`: an input that fetches a product from the Go backend and renders the result.

## Task

In `CatalogSearch.tsx`, implement:

- A text input bound to a `sku` state
- On input change, fetch `/api/catalog?sku=<value>` when the value is non-empty
- Track three states: `loading`, `data` (typed as `Product | null`), and `error`
- Render:
  - `"Loading..."` while the request is in flight
  - The product's name and price on success
  - The error message on failure

## Run

```bash
# Start the Go backend first:
go run ../../02-http-json/ex01/_solution/main.go

# Then in this directory:
npm run dev
npm test    # tests use mocked fetch — no backend needed
```

## Reset

```bash
cp _original/CatalogSearch.tsx CatalogSearch.tsx
```

## Concepts

- `useEffect` — run side effects after render
- `useState` for async state: `loading`, `data`, `error`
- Typed fetch response: `const data = await res.json() as Product`
- Cleanup with `AbortController` to avoid stale state
