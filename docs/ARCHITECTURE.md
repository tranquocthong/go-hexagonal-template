## Architecture

This template follows Domain-Driven Design and Hexagonal Architecture, applying CQRS in the application layer.

### Layers
- Domain: Entities and Domain Errors. Pure Go, no framework.
- Domain Ports: Hexagonal ports under `internal/domain/port`:
  - `in/`: inbound use case interfaces exposed to adapters
  - `out/`: outbound repository/external interfaces required by application
- Application: Implements inbound ports, depends on outbound ports. Organized as CQRS with `command/` and `query/` handlers, bundled by an `Application` aggregator.
- Adapters: Inbound (e.g., HTTP, background jobs) depend on inbound ports; Outbound (e.g., DB, external) implement outbound ports.

### Dependencies
- Domain: depends on nothing.
- Domain Ports: depend only on Domain.
- Application: depends on Domain and Domain Ports (`port/in`, `port/out`).
- Adapters:
  - Inbound adapters depend on `port/in` (call use cases).
  - Outbound adapters implement `port/out`.
- Composition root: `cmd/service` wires use cases and adapters.

### HTTP Flow Example (CQRS)
1. Request -> Inbound Adapter (`internal/adapters/inbound/http`)
2. Handler validates and calls application `Queries` or `Commands` via the `Application` aggregator (`internal/app`)
3. Handlers depend on outbound ports (`internal/domain/port/out`) to persist/fetch
4. Outbound adapter implements outbound port (`internal/adapters/outbound/...`)
5. Response mapped and returned

### Error Mapping
Domain errors map to HTTP status codes in the HTTP adapter.

### Testing Strategy
- Unit-test domain and application layers with in-memory ports.
- Integration-test adapters with real infrastructure.
- Blackbox/API tests at the boundary.

### Application Layer Structure
```
internal/app/
  app.go                 # Application aggregator (bundles Commands/Queries)
  command/               # Commands (write-side)
    create_greeting.go
    ...
  query/                 # Queries (read-side)
    get_greeting.go
    list_greetings.go
  service/               # Application/domain services
    entity_service.go
    converts/
      doc.go
```

`Application` wiring example:
```startLine:endLine:internal/app/app.go
// See file for full details; this shows the aggregator fields
type Application struct {
    Queries  Queries
    Commands Commands
}
```

### Project Structure

```
golang-template-prj/
├── cmd/
│   └── service/
│       └── main.go              # Composition root - wires all dependencies
│
├── internal/                    # Private application code
│   ├── domain/                  # Domain layer (entities + domain errors)
│   │   ├── error.go             # Domain-specific errors
│   │   ├── greeting.go          # Greeting entity/aggregate
│   │   ├── user.go              # User entity
│   │   └── port/                # Hexagonal ports
│   │       ├── in/              # Inbound ports (use cases)
│   │       │   └── greeting_usecases.go
│   │       └── out/             # Outbound ports (repositories/external)
│   │           └── greeting_repository.go
│   │
│   ├── app/                     # Application layer (CQRS)
│   │   ├── app.go               # Application aggregator
│   │   ├── command/             # Write-side handlers
│   │   │   └── create_greeting.go
│   │   ├── query/               # Read-side handlers
│   │   │   ├── get_greeting.go
│   │   │   └── list_greetings.go
│   │   └── service/             # Application services
│   │       ├── entity_service.go
│   │       └── converts/
│   │           └── doc.go       # Domain-DTO conversions
│   │
│   └── adapters/                # Adapters layer
│       ├── inbound/             # Inbound adapters
│       │   └── http/            # HTTP adapter
│       │       ├── server.go    # HTTP server setup
│       │       ├── auth.go      # Authentication handlers
│       │       └── server_extras.go
│       └── outbound/            # Outbound adapters
│           └── memory/          # In-memory repository implementation
│               └── greeting_repo.go
│
├── pkg/                         # Public/shared packages
│   ├── config/
│   │   └── config.go            # Environment configuration
│   ├── logger/
│   │   └── logger.go            # Structured logging (slog)
│   └── auth/
│       └── jwt/
│           └── jwt.go           # JWT token helpers
│
├── docs/                        # Documentation
│   ├── GETTING_STARTED.md       # Quick start guide
│   ├── ARCHITECTURE.md          # Architecture documentation
│   └── API_REFERENCE.md         # API endpoints reference
│
├── bin/                         # Compiled binaries (gitignored)
│
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── Makefile                     # Build and development tasks
├── Dockerfile                   # Container image definition
├── docker-compose.yml           # Docker composition
├── env.example                  # Example environment variables
├── .gitignore                   # Git ignore patterns
├── .dockerignore                # Docker ignore patterns
└── README.md                    # Project overview
```

### Directory Responsibilities

#### `/cmd/service`
- Composition root where all dependencies are wired together
- Creates instances of adapters and injects them into the application layer
- Starts the HTTP server

#### `/internal/domain`
- **Pure domain logic** - no external dependencies
- Entities, value objects, and domain errors
- Business rules and invariants
- Domain ports define interfaces for use cases (in/) and repositories (out/)

#### `/internal/app`
- Implements inbound ports (use cases)
- Depends on outbound ports for data access
- Organized using CQRS pattern:
  - **Commands**: Write operations (create, update, delete)
  - **Queries**: Read operations (get, list)
- Application services for cross-cutting domain logic

#### `/internal/adapters`
- **Inbound adapters**: Entry points (HTTP, CLI, gRPC, etc.)
  - Call application use cases through inbound ports
  - Handle protocol-specific concerns (HTTP routing, middleware, error mapping)
- **Outbound adapters**: Infrastructure implementations
  - Implement outbound ports (repositories, external services)
  - Handle persistence, caching, external APIs

#### `/pkg`
- Reusable packages that could be extracted to separate libraries
- No business logic - only technical utilities
- Examples: config loading, logging, authentication helpers

### Dependency Flow
```
Inbound Adapter → Inbound Port → Application (Use Case) → Outbound Port → Outbound Adapter
    (HTTP)           (interface)       (command/query)      (interface)      (repository)
```


