# GitHub Actions

## Short Summary

A basic pipeline for this kind of role should validate quickly, fail early, and separate build, test, and deploy into clear steps.

## Base Pipeline You Should Be Able to Explain

- checkout the repository
- set up Go
- cache modules
- `go test ./...`
- build
- optional extra analysis
- deploy conditioned on branch or environment

## Typical Questions

### What would you pay attention to in CI

- fast feedback
- small, readable jobs
- well-managed secrets
- promoting immutable artifacts

### How would you use AI tools without losing quality

I would use Copilot or similar tools to speed up implementation, but I would keep tests, code review, and pipeline policies as the safety net.
