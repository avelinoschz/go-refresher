# Cloud Run vs GKE

## Short Summary

- Cloud Run works very well for stateless HTTP services that need fast delivery with low operational friction.
- GKE offers more power and more complexity; choose it when you need fine-grained control over networking, sidecars, scheduling, or more specialized workloads.
- In an internal platform, Cloud Run often speeds up onboarding, while GKE tends to show up when teams need more advanced capabilities or Kubernetes standardization.

## Typical Questions

### When would you choose Cloud Run over GKE

When the service is stateless, has clear HTTP traffic, does not need deep cluster control, and we want to minimize operations.

### When would you choose GKE

When we need native Kubernetes features, sidecars, DaemonSets, complex policies, long-lived workloads, or advanced network configuration.

### Risks or tradeoffs of Cloud Run

- less fine-grained runtime control
- cold starts depending on traffic patterns
- limits and a more opinionated model

### Summary

For a migration platform, Cloud Run helps teams move faster because it reduces operational load. If an application needs more advanced Kubernetes capabilities or more infrastructure control, then GKE becomes the better option.
