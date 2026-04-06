# Go Core Ex03: In-Memory Service

Goal:

- introduce a small interface that actually helps
- separate `service` and `store`
- practice business-level errors

Prompt:

Implement a small application onboarding flow.

You need:

- `AppStore` with `Save` and `GetByName`
- `InMemoryAppStore`
- `AppService.Register`

Rules:

- `Name` and `OwnerTeam` are required
- duplicate names are not allowed
- `Register` must delegate persistence to the store

Goal:

- make the code feel like a small version of a real backend service

Command:

```bash
go run ./go-refresher/01-go-core/ex03
```

## Run tests

```bash
go test ./go-refresher/01-go-core/ex03
```
