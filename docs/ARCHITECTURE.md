## Architecture

This template follows Domain-Driven Design and Hexagonal Architecture.

### Layers
- Domain: Entities and Domain Errors. Pure Go, no framework.
- Domain Ports: Hexagonal ports under `internal/domain/port`:
  - `in/`: inbound use case interfaces exposed to adapters
  - `out/`: outbound repository/external interfaces required by application
- Application: Implements inbound ports, depends on outbound ports.
- Adapters: Inbound (e.g., HTTP, background jobs) depend on inbound ports; Outbound (e.g., DB, external) implement outbound ports.

### Dependencies
- Domain: depends on nothing.
- Domain Ports: depend only on Domain.
- Application: depends on Domain and Domain Ports (`port/in`, `port/out`).
- Adapters:
  - Inbound adapters depend on `port/in` (call use cases).
  - Outbound adapters implement `port/out`.
- Composition root: `cmd/service` wires use cases and adapters.

### HTTP Flow Example
1. Request -> Inbound Adapter (`internal/adapters/inbound/http`)
2. Handler validates and calls inbound port (`internal/domain/port/in`) implemented by app (`internal/app`)
3. Use-case depends on outbound port (`internal/domain/port/out`) to persist/fetch
4. Outbound adapter implements outbound port (`internal/adapters/outbound/...`)
5. Response mapped and returned

### Error Mapping
Domain errors map to HTTP status codes in the HTTP adapter.

### Testing Strategy
- Unit-test domain and application layers with in-memory ports.
- Integration-test adapters with real infrastructure.
- Blackbox/API tests at the boundary.


