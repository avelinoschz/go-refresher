# HTTP Ex08: Simple HTTP Client

Goal:

- practice making HTTP requests with the default client
- handle response status codes explicitly
- read and close response bodies correctly
- encode and decode JSON over HTTP

Prompt:

Complete the client functions in `main.go`.

Implement:

- `FetchProduct`
- `CreateProduct`

Rules:

- use `http.Get` for the GET request
- use `http.Post` for the POST request
- always close the response body with `defer resp.Body.Close()`
- return an error if the status code is not the expected one
- decode the response body with `json.NewDecoder`

## Run tests

```bash
go test ./02-http-json/ex08
```
