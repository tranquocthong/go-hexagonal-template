## API Reference

### Health
GET /healthz
Response 200
```json
{"status":"ok","time":"2025-01-01T00:00:00Z"}
```

### Auth
POST /api/v1/login
Request
```json
{"email":"user@example.com","password":"secret"}
```
Response 200
```json
{"access_token":"<jwt>"}
```

GET /api/v1/me (Bearer token)
Response 200
```json
{"subject":"user@example.com"}
```

### Greetings
GET /api/v1/greetings
Response 200
```json
[{"id":"hello","message":"Hello, World!","created_at":"..."}]
```

GET /api/v1/greetings/{id}
Response 200
```json
{"id":"hello","message":"Hello, World!","created_at":"..."}
```

POST /api/v1/greetings (Bearer token)
Request
```json
{"id":"hi","message":"Hi there"}
```
Response 201
```json
{"id":"hi","message":"Hi there","created_at":"..."}
```

Notes:
- Reads are served via application Queries; writes via Commands.


