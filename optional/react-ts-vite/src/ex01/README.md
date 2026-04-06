# Exercise 01: useState with Types

## Goal

Implement a `Counter` component using `useState<number>`.

## Task

In `Counter.tsx`, implement:

- A numeric state initialized to `0`
- An **Increment** button that adds 1
- A **Decrement** button that subtracts 1
- A **Reset** button that sets the count back to `0`
- Display the current count

## Run

```bash
npm run dev       # view in browser at localhost:5173
npm test          # run tests
```

## Reset

```bash
cp _original/Counter.tsx Counter.tsx
```

## Concepts

- `useState<number>` — typed state
- `React.FC` — functional component type
- `React.MouseEventHandler` — event handler type
