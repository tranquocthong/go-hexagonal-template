## Getting Started

This template is a production-ready Golang backend following DDD and Hexagonal Architecture (Ports & Adapters).

### Prerequisites
- Go 1.23+ (see `go.mod` for the exact version required)
- Make (optional)
- Docker (optional)
- [`swag`](https://github.com/swaggo/swag) CLI for regenerating API docs:
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

### Run locally
```bash
cd golang-template-prj
make tidy
make run
# open http://localhost:8080/healthz
# open http://localhost:8080/swagger/index.html
```

### Endpoints
- `GET /healthz` — service status
- `GET /api/v1/greetings` — list greetings
- `GET /api/v1/greetings/{id}` — get by id
- `POST /api/v1/greetings` (Bearer token required) — create greeting
- `POST /api/v1/login` — **demo only**: returns a JWT for any non-empty email/password
- `GET /api/v1/me` (Bearer token required) — show current subject

### Config
Environment variables (see `env.example`):

| Variable | Default | Description |
|---|---|---|
| `APP_NAME` | `yourservice` | Service name |
| `APP_ENV` | `development` | Environment |
| `LOG_LEVEL` | `info` | Log level (`debug`, `info`, `warn`, `error`) |
| `HTTP_ADDRESS` | `:8080` | Listen address |
| `JWT_SECRET` | `dev_secret_change_me` | **Change in production** (`openssl rand -hex 32`) |
| `JWT_TTL` | `1h` | Token lifetime |

> The service logs a warning at startup when `JWT_SECRET` is using the default value.

### Run tests
```bash
make test
```

### Build Docker image
```bash
make docker
make docker-run
```

### Next Steps
1. Replace `example.com/yourorg/yourservice` module path with your own
2. Replace in-memory repository with a real adapter (PostgreSQL, MongoDB, etc.)
3. Replace the demo `POST /api/v1/login` with real credential validation
4. Add migrations, tracing, metrics as needed
