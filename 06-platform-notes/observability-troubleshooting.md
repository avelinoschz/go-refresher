# Observability and Troubleshooting

## Short Summary

In a platform/backend role, it matters a lot to isolate whether the problem comes from code, configuration, networking, an external dependency, or deployment.

## Things You Should Mention Naturally

- structured logs
- basic metrics for latency, error rate, and throughput
- request IDs or trace IDs
- health checks
- timeouts and retries used with intent

## Typical Questions

### How would you investigate a slow service in production

I would start by reviewing metrics and logs to narrow down whether the issue is broad or isolated, then check external dependencies, per-endpoint latency, and recent deployment or configuration changes.

### What would you add to a new service

- structured logs
- request metrics
- a health endpoint
- timeouts on external clients
- basic dashboards and alerts
