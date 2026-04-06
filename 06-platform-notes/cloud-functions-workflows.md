# Cloud Functions and Workflows

## Short Summary

- Cloud Functions fits small, well-bounded event-driven work.
- Workflows coordinates steps across services, especially when there is orchestration, retries, or calls to multiple APIs.

## Typical Questions

### When would you avoid Cloud Functions

When the logic is no longer small, when the flow needs too much state, or when the service fits better as a longer-lived and stable API.

### When would you use Workflows

To orchestrate steps across GCP services or external APIs without embedding all that coordination into a single application.

### How would you connect this to an internal platform

You could use Workflows to automate onboarding, approvals, or operational sequences; for example, validating input, creating resources, and notifying outcomes.
