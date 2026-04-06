# Warmup Ex02: Slice + Map Summary

Goal:

- practice data transformation
- iterate over slices
- accumulate results in a map
- return multiple values

Prompt:

Implement `BuildDepartmentSummary`.

Input:

- a list of employees

Output:

- a `map[string]int` with the total number of employees per department
- an integer with the total number of active employees

Rules:

- ignore employees with `Department == ""`
- only count employees with `Active == true` in `activeTotal`
- the `map` must include every valid department even if it only has one employee

Done criteria:

- it compiles
- it uses a single main loop
- it does not use external packages

Command:

```bash
go run ./go-refresher/00-warmup/ex02
```

## Run tests

```bash
go test ./go-refresher/00-warmup/ex02
```
