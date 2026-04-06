# Patrones de cache

Referencia de los patrones más comunes. El bloque implementa **Cache-Aside** en ex02 y **Graceful Degradation** en ex03.

---

## Cache-Aside (Lazy Loading)

> "El caller es responsable de poblar el cache."

Es el patrón implementado en `ex02`. La aplicación maneja el cache explícitamente: primero consulta Redis, y solo si hay miss va al store.

```text
┌────────┐    1. GET key     ┌───────┐
│ caller │ ───────────────►  │ Redis │
│        │ ◄───────────────  │       │
│        │   2a. HIT: value  └───────┘
│        │
│        │   2b. MISS        ┌───────┐
│        │ ───────────────►  │ Store │
│        │ ◄───────────────  │  (DB) │
│        │   3. value        └───────┘
│        │
│        │   4. SET key      ┌───────┐
│        │ ───────────────►  │ Redis │
└────────┘                   └───────┘
```

**Implementación en ex02:**

```go
func (s CatalogService) GetProductName(ctx context.Context, sku string) (string, error) {
    // 1. Check cache
    if name, ok, _ := s.cache.Get(ctx, "catalog:"+sku); ok {
        return name, nil
    }
    // 2. Miss: fetch from store
    name, err := s.store.GetName(sku)
    if err != nil {
        return "", err
    }
    // 3. Populate cache
    _ = s.cache.Set(ctx, "catalog:"+sku, name, 30*time.Second)
    return name, nil
}
```

**Trade-offs:**

| | |
| - | - |
| El cache solo contiene datos que realmente se pidieron | Primera petición de cada key es lenta (cache miss) |
| Fácil de implementar y razonar | Si el cache cae, el store absorbe toda la carga hasta que se reconstruye |
| La app tolera un cache caído (lee del store) | Posible **cache stampede**: muchos callers hacen miss al mismo tiempo y saturan el store |
| Funciona bien con TTL corto | Datos en cache pueden quedar stale hasta que expiren |

**Cuándo usarlo:** Es el default para la mayoría de los casos de lectura. Si no tienes un requisito específico, empieza aquí.

---

## Read-Through

> "El cache maneja el fetch automáticamente."

La diferencia con Cache-Aside es que el cache mismo es responsable de ir al store en caso de miss. La aplicación solo habla con el cache.

```text
┌────────┐    GET key    ┌────────────────────┐    fetch   ┌───────┐
│ caller │ ────────────► │   Cache (library)  │ ─────────► │ Store │
│        │ ◄──────────── │   + loader fn      │ ◄───────── │       │
└────────┘    value      └────────────────────┘            └───────┘
```

La "loader function" se registra en el cliente de cache y se invoca automáticamente en cada miss.

**Trade-offs:**

| | |
| - | - |
| El código de la aplicación queda más limpio (no maneja misses) | El cache es un punto de fallo: si cae, todo falla |
| Consistencia: un solo lugar que decide cómo poblar el cache | Más difícil de razonar sobre errores (la lógica está escondida en el loader) |
| Bueno cuando muchos callers acceden a los mismos datos | Requiere soporte explícito en el cliente de cache o una capa extra |

**Cuándo usarlo:** Cuando quieres ocultar la lógica de cache a los callers y el cache es confiable. Menos común con Redis puro; más habitual con librerías como `ristretto` o `groupcache`.

---

## Write-Through

> "Se escribe en cache y en store al mismo tiempo."

Cada escritura actualiza Redis y el store en la misma operación. El cache nunca queda stale porque siempre tiene la versión más reciente.

```text
┌────────┐    SET key    ┌───────┐    write   ┌───────┐
│ caller │ ────────────► │ Redis │ ─────────► │ Store │
│        │ ◄──────────── │       │ ◄───────── │       │
└────────┘    ok         └───────┘    ok      └───────┘
```

**Trade-offs:**

| | |
| - | - |
| El cache siempre tiene datos frescos | Cada escritura paga doble latencia (Redis + store) |
| No hay stale data ni necesidad de TTL corto | Se cachean datos que quizás nunca se lean (write sin read posterior) |
| Lecturas subsiguientes siempre son cache hit | Si el store falla, la escritura completa falla (no hay degradación parcial) |

**Cuándo usarlo:** Cuando los datos se leen frecuentemente después de cada escritura y la consistencia es crítica. Típico en perfiles de usuario o configuraciones.

---

## Write-Behind (Write-Back)

> "Se escribe en cache primero; el store se actualiza de forma asíncrona."

La escritura confirma inmediatamente al caller después de actualizar Redis. Un proceso background drena las escrituras pendientes al store.

```text
┌────────┐    SET key    ┌───────┐
│ caller │ ────────────► │ Redis │  ◄─── confirmación inmediata
└────────┘               └───┬───┘
                              │  async (batch / debounce)
                              ▼
                          ┌───────┐
                          │ Store │
                          └───────┘
```

**Trade-offs:**

| | |
| - | - |
| Escrituras muy rápidas (no bloquean en el store) | **Riesgo de pérdida de datos**: si Redis cae antes del flush, las escrituras se pierden |
| Permite batching: múltiples escrituras → un solo INSERT/UPDATE | Mucho más complejo de implementar correctamente |
| Ideal para workloads write-heavy | El store puede quedar temporalmente desincronizado |

**Cuándo usarlo:** Contadores de visitas, métricas de uso, logs de actividad — casos donde perder algún dato es aceptable. **Evitarlo para datos transaccionales.**

---

## Comparativa

| Patrón | Quién puebla el cache | Consistencia | Complejidad | Tolerancia a fallo del cache |
| - | - | - | - | - |
| Cache-Aside | La aplicación (en lectura) | Eventual (TTL) | Baja | Alta — cae al store |
| Read-Through | El cache (loader fn) | Eventual (TTL) | Media | Baja — falla si cae el cache |
| Write-Through | La aplicación (en escritura) | Fuerte | Media | Alta — cache es warm siempre |
| Write-Behind | La aplicación (async) | Eventual (lag) | Alta | Baja — riesgo de pérdida |

---

## Graceful Degradation (ex03)

No es un patrón de escritura sino una estrategia de resiliencia para **cualquiera** de los patrones anteriores. La idea: el cache es best-effort, el store es la fuente de verdad.

```go
func (s Service) GetWithFallback(ctx context.Context, sku string) (string, error) {
    // Acota el tiempo que le damos a Redis
    cacheCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
    defer cancel()

    if name, ok, err := s.cache.Get(cacheCtx, "catalog:"+sku); err == nil && ok {
        return name, nil
    }
    // Redis lento, caído, o miss → vamos directo al store
    return s.store.GetName(sku)
}
```

**Por qué 100ms:** En un servicio con un SLA de ~200ms, dedicar más de la mitad al cache es demasiado. El timeout evita que un Redis lento degrade toda la respuesta.

**El patrón aplica a todos:** Cache-Aside, Read-Through y Write-Through pueden (y deben) tener degradación graceful en producción.
