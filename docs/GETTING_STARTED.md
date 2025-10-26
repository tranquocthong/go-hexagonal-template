## Getting Started

This template is a production-ready Golang backend following DDD and Hexagonal Architecture (Ports & Adapters).

### Prerequisites
- Go 1.22+
- Make (optional)
- Docker (optional)

### Run locally
```bash
cd template
make tidy
make run
# open http://localhost:8080/healthz
```

### Endpoints
- GET `/healthz`: service status
- GET `/api/v1/greetings`: list seeded greetings
- GET `/api/v1/greetings/{id}`: get by id
- POST `/api/v1/greetings` (Bearer token required): create greeting
- POST `/api/v1/login`: demo login returns JWT
- GET `/api/v1/me` (Bearer token required): show subject

### Config
Environment variables (see `.env.example`):
- `APP_NAME`, `APP_ENV`, `LOG_LEVEL`
- `HTTP_ADDRESS`, timeouts
- `JWT_SECRET`, `JWT_TTL`

### Build Docker image
```bash
make docker
make docker-run
```

### Project Structure
```
cmd/service            # composition root
internal/
  domain/              # entities + domain errors
  domain/port/         # hexagonal ports
    in/                # inbound use cases
    out/               # outbound repositories/external
  app/                 # CQRS: commands and queries, wired by Application
    command/           # write handlers
    query/             # read handlers
    service/           # shared services and converts
  adapters/
    inbound/http       # HTTP handlers, middleware, server
    outbound/memory    # example repo implementation
pkg/
  config/              # env config
  logger/              # slog JSON logger
  auth/jwt             # JWT helpers
docs/                  # docs
```

### Next Steps
- Replace in-memory repositories with real adapters (e.g., PostgreSQL, MongoDB)
- Add migrations, tracing, metrics, CI
- Implement domain-specific aggregates and use-cases


