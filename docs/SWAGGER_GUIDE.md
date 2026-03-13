# Swagger API Documentation Guide

This project integrates **swaggo/swag v2** to automatically generate API documentation (Swagger UI) from annotations in the Go source code. The generated spec follows the **OpenAPI 3.1** standard.

## How It Works

The entire API documentation is generated based on declarative comments (starting with `// @`) placed above the `main()` function and HTTP handler functions in `internal/adapters/inbound/http/server.go`.

The tool parses these comments and automatically creates the OpenAPI 3.1 structure files (`swagger.json`, `swagger.yaml`) in the `docs/swagger` directory.

## Makefile Commands

To simplify the process, the necessary commands have been integrated into the `Makefile`:

- **Generate Swagger Documentation:**
  ```bash
  make swagger
  ```
  *Note: This runs the command `swag init --v3.1 -g cmd/service/main.go -o ./docs/swagger`*

- **Run the Application (Auto-generates Swagger beforehand):**
  ```bash
  make run
  ```

## Installing the swag CLI

The `swag` CLI tool must be installed separately. Install swag v2:

```bash
go install github.com/swaggo/swag/v2/cmd/swag@latest
```

## Viewing Documentation in Browser

1. Start the application by running `make run`.
2. Open your web browser and navigate to:
   **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

Here, you will see the Swagger UI dashboard displaying all defined API endpoints.

## How to Add/Update Documentation for a New API

When creating a new API endpoint, follow these steps to ensure the documentation is updated:

1. **Define Request/Response Schema (Optional)**
   If your endpoint receives or returns JSON, ensure you have valid structs defined (e.g., `domain.User`, `createGreetingRequest`). You can use the `example:"..."` struct tag to display sample data in Swagger.

2. **Add Annotations to the Handler**
   Write comments immediately above the handler function. For example:
   ```go
   // handleGetSomething godoc
   // @Summary      Short description of the API
   // @Description  More detailed description of the endpoint's functionality
   // @Tags         category_name
   // @Produce      json
   // @Success      200 {object} domain.MyResponseObject "Success"
   // @Router       /api/v1/something [get]
   func (s *Server) handleGetSomething(w http.ResponseWriter, r *http.Request) {
       // logic...
   }
   ```
   *You can reference existing handlers in `server.go` and `server_extras.go` for examples of Request Body parameters (`@Param`), URL Path variables (`@Param id path`), JWT Security (`@Security BearerAuth`), etc.*

3. **Update the Router**
   Register the new handler with the HTTP multiplexer (mux).

4. **Build and Run**
   Execute `make run`. The tool will automatically regenerate the Swagger JSON based on your new code and instantly display the updates on the UI.

## Global Annotations (main.go)

The top-level API metadata is declared above `func main()` in `cmd/service/main.go`:

```go
// @title           Golang Template API
// @version         1.0
// @description     A description of your API
// @servers         http://localhost:8080

// @securityDefinitions.bearer BearerAuth
// @bearerFormat JWT
```

> **Note:** `@servers` replaces `@host` + `@BasePath` from OpenAPI 2.0. `@securityDefinitions.bearer` replaces `@securityDefinitions.apikey` and correctly describes HTTP Bearer token authentication.

## Security Annotation for Protected Endpoints

To mark an endpoint as requiring JWT authentication:

```go
// @Security BearerAuth
```

---
> 🔗 View the full list of supported declarative comments at the official documentation: [Swaggo Declarative Comments Format](https://github.com/swaggo/swag/tree/v2?tab=readme-ov-file#declarative-comments-format)
