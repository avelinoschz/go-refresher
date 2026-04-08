# HTTP Ex09: Configured HTTP Client

Goal:

- practice building a reusable HTTP client struct
- configure `http.Client` with an explicit timeout
- use `http.NewRequestWithContext` to propagate context cancellation
- set request headers explicitly

Prompt:

Complete `CatalogClient` and its methods in `main.go`.

Implement:

- `NewCatalogClient`
- `GetProduct`
- `CreateProduct`

Rules:

- configure `http.Client{Timeout: timeout}` in `NewCatalogClient`
- use `http.NewRequestWithContext` to attach the context to every request
- perform requests with `c.client.Do`
- always close the response body with `defer resp.Body.Close()`
- set `Content-Type: application/json` explicitly on POST requests
- return an error if the context is canceled or the status code is unexpected

## Run tests

```bash
go test ./02-http-json/ex09
```
