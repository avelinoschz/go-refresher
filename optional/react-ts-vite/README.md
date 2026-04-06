# Optional: React + TypeScript

This block is optional — it complements the Go track with frontend exercises.

Goal:

- avoid going cold on React if full-stack work shows up
- practice the minimum useful set of modern React + TypeScript concepts

## Setup

```bash
make react-install   # from repo root
# or:
cd optional/react-ts-vite && npm install
```

## Run exercises

```bash
cd optional/react-ts-vite
npm run dev    # opens localhost:5173
npm test       # run all tests (no backend needed — fetch is mocked)
```

## Exercises

| Exercise | Level | Concept |
|----------|-------|---------|
| ex01 | Simple | `useState<T>` — Counter with typed state |
| ex02 | Simple | Typed props, optional props, conditional rendering |
| ex03 | Intermediate | `useEffect` + fetch to Go backend, loading/error/data states |
| ex04 | Intermediate | Custom hook — extract fetch logic to `useCatalog(sku)` |
| ex05 | Intermediate | Controlled input, form handling, list rendering |

## With Go backend (ex03, ex05)

```bash
# In another terminal, from repo root:
go run ./02-http-json/ex01/_solution/main.go
# or the full session solution:
go run ./06-timed-practice/session01/_solution/main.go
```

Vite proxies `/api/*` → `localhost:8080/*`, so no CORS issues.

## Reset an exercise

```bash
# Example for ex01:
cp src/ex01/_original/Counter.tsx src/ex01/Counter.tsx
```

Each exercise has `_original/` (blank skeleton) and `_solution/` (reference).
