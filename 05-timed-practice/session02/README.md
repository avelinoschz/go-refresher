# Session 02: Add a New Requirement

Target duration:

- 45 to 60 minutes

Problem:

Take the previous session and add a new requirement:

- support `GET /catalog?sku=...&include=meta`

When `include=meta`, the response must include:

- `source`: `cache` or `store`
- `served_at`: a timestamp in RFC3339 format

Constraints:

- do not break the base product format
- try to extend the solution without rewriting everything

Done criteria:

- small but clean change
- no excessive logic duplication
- tests adapted to the new response

Follow-up topics:

- how you would version the API if this kept growing
- how to avoid coupling the handler too tightly to the response format
