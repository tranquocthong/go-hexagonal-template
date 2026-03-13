## Golang Service Template (DDD + Hexagonal)

Production-ready template for Go backends following DDD and Hexagonal architecture.

### Customizing for Your Project
Before using this template for a real project, you should replace the generic module name (`example.com/yourorg/yourservice`) with your own (e.g., `github.com/myusername/myproject`).

1. Rename the module in `go.mod`.
2. Find and replace all occurrences of `example.com/yourorg/yourservice` in all `.go` files with your new module name.

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
- docs/SWAGGER_GUIDE.md

### Highlights
- Clean layering (Domain, Application, Ports, Adapters)
- JSON logging (slog), config via env
- JWT auth with protected routes
- Example endpoints illustrating the full flow


