# 05-redis

Este bloque es la progresión natural de `04-cache-concurrency`: reemplaza el cache in-memory por Redis real usando `go-redis/v9`.

## Ejercicios

| # | Ejercicio | Concepto clave |
| - | - | - |
| ex01 | Redis Client Wrapper | `GET`/`SET`, manejo de `redis.Nil` como miss |
| ex02 | Cache-Aside con Redis | key format, TTL de 30s, cache hit/miss |
| ex03 | Graceful Degradation | `context.WithTimeout`, fallback al store |

## Requisitos

- Docker (para levantar Redis)
- Go 1.22+

## Setup

```bash
# Levantar Redis
make redis-up

# Validar la conexión antes de correr los tests
redis-cli PING
# Debe responder: PONG

# Detener Redis al terminar
make redis-down
```

## Correr los tests

```bash
# Un ejercicio a la vez
go test -race ./05-redis/ex01/
go test -race ./05-redis/ex02/
go test -race ./05-redis/ex03/

# Todo el bloque
go test -race ./05-redis/...
```

> Los tests requieren Redis corriendo en `localhost:6379`. Si fallan con "connection refused", verifica que `make redis-up` haya levantado el contenedor.
