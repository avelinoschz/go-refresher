# 02 HTTP JSON

This is the most hands-on block for real backend HTTP work.

Priorities:

- `net/http`
- `encoding/json`
- query params
- path params with the stdlib router
- body decoding
- strict JSON decoding (`DisallowUnknownFields`, single-object bodies, size limits)
- simple validation
- consistent JSON responses
- introduce `context.Context` without overcomplicating things

Rule:

- avoid frameworks
- aim for code you could explain in 2-3 minutes

## Run tests

```bash
go test ./02-http-json/...
```
