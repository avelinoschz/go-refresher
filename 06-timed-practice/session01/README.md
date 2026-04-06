# Session 01: Catalog Endpoint with Cache

Target duration:

- 45 to 60 minutes

Problem:

Implement an HTTP service with this endpoint:

- `GET /catalog?sku=...`

Behavior:

- validate `sku`
- check the cache first
- if it is missing, query a simulated primary source
- if the product exists, respond with JSON and store the result in the cache
- if it does not exist, respond with `404`

Constraints:

- use only the standard library
- do not use frameworks
- keep a reasonable separation between handler, service, and cache

Done criteria:

- code compiles
- clear handler
- correct errors and status codes
- clear cache hit/miss behavior
- basic tests included

Follow-up topics:

- how this would change if it ran on Cloud Run
- what observability you would add
- where you would place timeouts and why
