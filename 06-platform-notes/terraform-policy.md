# Terraform, OPA, Rego, CEL

## Short Summary

- Terraform is used to declare infrastructure in a repeatable way.
- The important part is not memorizing syntax, but understanding modules, variables, state, and drift.
- OPA/Rego and CEL help express rules or policies to validate configurations and deployments.

## Typical Questions

### What is drift

It is the difference between what is declared and what actually exists in infrastructure.

### What are modules for

To reuse patterns and reduce copy/paste in infrastructure as code.

### Where would policy-as-code fit

Before or during deployment, to block configurations that do not meet security, naming, networking, or compliance standards.

### How to approach OPA/Rego or CEL if you are not an expert

Not a specialist, but the goal is clear: expressing declarative rules to validate configurations. The concrete syntax is learnable while keeping the evaluation flow clear.
