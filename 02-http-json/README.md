# 02 HTTP JSON

This is the most hands-on block for real backend HTTP work.

Priorities:

- `net/http`
- `encoding/json`
- query params
- path params with the stdlib router
- server bootstrap and explicit `http.Server` timeouts
- `ServeMux` route wiring and middleware composition
- body decoding
- strict JSON decoding (`DisallowUnknownFields`, single-object bodies, size limits)
- simple validation
- consistent JSON responses
- introduce `context.Context` without overcomplicating things

Rule:

- avoid frameworks
- aim for code you could explain in 2-3 minutes

## Quick recall: graceful shutdown

A good stdlib pattern to remember for interviews and backend practice:

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

srv := &http.Server{
    Addr:              ":8080",
    Handler:           handler,
    ReadHeaderTimeout: 2 * time.Second,
    WriteTimeout:      5 * time.Second,
    IdleTimeout:       30 * time.Second,
}

go func() {
    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _ = srv.Shutdown(shutdownCtx)
}()

if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
    return err
}
```

## Run tests

```bash
go test ./02-http-json/...
```
