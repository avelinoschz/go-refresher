# Exercise 02: Typed Props and Conditional Rendering

## Goal

Implement a `ProductCard` component that receives a `Product` as a prop and optionally renders a badge.

## Task

In `ProductCard.tsx`, implement:

- Render the product's `sku`, `name`, and `price`
- Accept an optional `badge` string prop
- If `badge` is provided, render it visibly (e.g. `[New]`)
- If `badge` is not provided, render nothing for it

## Run

```bash
npm run dev
npm test
```

## Reset

```bash
cp _original/ProductCard.tsx ProductCard.tsx
```

## Concepts

- `interface Props` — typed component props
- `React.FC<Props>` — component with typed props
- Optional props with `?`
- Conditional rendering with `&&`
