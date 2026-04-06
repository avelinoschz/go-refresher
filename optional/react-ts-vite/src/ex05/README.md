# Exercise 05: Controlled Input and Search History

## Goal

Implement `CatalogList`: a controlled search form that displays a list of products looked up by SKU.

## Task

In `CatalogList.tsx`, implement:

- A controlled `<input>` bound to a `query` state
- A `<form>` with an `onSubmit` handler
- On submit: look up the SKU using `useCatalog` and add the result to a `results` list
- Render the list of past results (each with its sku and name)
- Clear the input after a successful submit
- Show an inline error if the SKU is not found

## Run

```bash
# Start the Go backend first:
go run ../../02-http-json/ex01/_solution/main.go

npm run dev
npm test    # tests use mocked fetch
```

## Reset

```bash
cp _original/CatalogList.tsx CatalogList.tsx
```

## Concepts

- Controlled inputs: `value` + `onChange`
- `onSubmit` with `e.preventDefault()`
- Rendering lists with `key`
- Composing custom hooks
