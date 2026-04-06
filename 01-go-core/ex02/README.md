# Go Core Ex02: Composition and Errors

Goal:

- practice simple composition
- use errors to stop invalid flow
- get used to separating `service` and `store`

Prompt:

Complete `DeploymentPlanner`.

It must receive:

- environment
- service name
- desired replicas

It must validate:

- allowed environment: `dev`, `stage`, `prod`
- non-empty service name
- positive replicas
- in `prod`, a minimum of 2 replicas

Output:

- a `DeploymentPlan` ready to print or use later

Command:

```bash
go run ./01-go-core/ex02/main.go
```

## Run tests

```bash
go test ./01-go-core/ex02
```
