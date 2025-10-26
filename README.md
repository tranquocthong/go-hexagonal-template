## Golang Service Template (DDD + Hexagonal)

Production-ready template for Go backends following DDD and Hexagonal architecture.

### Quickstart
```bash
make tidy
make run
# http://localhost:8080/healthz
```

### Docs
- docs/GETTING_STARTED.md
- docs/ARCHITECTURE.md
- docs/API_REFERENCE.md

### Highlights
- Clean layering (Domain, Application, Ports, Adapters)
- JSON logging (slog), config via env
- JWT auth with protected routes
- Example endpoints illustrating the full flow


