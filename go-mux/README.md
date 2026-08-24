# Gorilla Mux Complete Tutorial

A comprehensive, hands-on tutorial for learning Gorilla Mux router for Go web development.

## Prerequisites

- Go 1.24+ installed
- Basic Go knowledge
- Terminal/command line familiarity

### Coming from Spring Boot?

If you know Spring Boot, check out **[SPRINGBOOT_VS_MUX.md](SPRINGBOOT_VS_MUX.md)** for a complete comparison showing how to do:

- REST APIs (`@RestController` → `mux.Router`)
- Configuration (`application.yml` → `.env` + config structs)
- Database/JPA (`@Entity`, `JpaRepository` → GORM)
- Validation (`@Valid` → `go-playground/validator`)
- Cron Jobs (`@Scheduled` → `robfig/cron`)
- Caching (`@Cacheable` → Redis with go-redis)
- Security (Spring Security → Custom middleware + JWT)
- Testing (`@SpringBootTest` → `httptest` + testify)
- And more!

## Quick Start

```bash
cd mux-tutorial

# Start with basics
cd 01-basic-routes && go run main.go

# Then path/query params
cd 02-path-and-query-params && go run main.go

# Environment variables
cd 03-env-config && go run main.go

# Docker (from 04-docker folder)
docker-compose up

# Authentication
cd 05-authentication && go run main.go
```

---

## Tutorial Structure

| Folder | Port | Topics |
|--------|------|--------|
| [01-basic-routes](01-basic-routes/) | 8100 | Simple routes, JSON, HTML, status codes |
| [02-path-and-query-params](02-path-and-query-params/) | 8101 | Path variables, query parameters, regex patterns |
| [03-env-config](03-env-config/) | 8000 | Environment variables, .env files, config patterns |
| [04-docker](04-docker/) | 8000 | Dockerfile, docker-compose, multi-stage builds |
| [05-authentication](05-authentication/) | 8081 | Basic Auth, JWT, API Key, OAuth 2.0 |

---

## What You'll Learn

### 01 - Basic Routes
- Creating a mux router
- Registering routes
- Handler functions
- JSON responses
- HTML responses
- Custom status codes

### 02 - Path & Query Parameters
- Path variables `{id}`
- Regex patterns `{id:[0-9]+}`
- Multiple path variables
- Query parameter extraction
- Pagination patterns
- Combined path + query

### 03 - Environment Variables
- Reading env vars with `os.Getenv()`
- Using `.env` files with godotenv
- Config struct pattern
- Graceful shutdown
- Server timeouts

### 04 - Docker
- Multi-stage Docker builds
- docker-compose for orchestration
- PostgreSQL + Redis services
- Health checks
- Best practices

### 05 - Authentication
- Basic Authentication
- JWT tokens
- API Key authentication
- OAuth 2.0 (GitHub)
- Password hashing with bcrypt
- Role-based access control

---

## Running All Examples

| Example | Command | URL |
|---------|---------|-----|
| 01 - Basic Routes | `cd 01-basic-routes && go run main.go` | http://localhost:8100 |
| 02 - Path & Query | `cd 02-path-and-query-params && go run main.go` | http://localhost:8101 |
| 03 - Env Config | `cd 03-env-config && go run main.go` | http://localhost:8000 |
| 04 - Docker | `cd 04-docker && docker-compose up` | http://localhost:8000 |
| 05 - Auth | `cd 05-authentication && go run main.go` | http://localhost:8081 |

---

## Testing APIs

### Using curl

```bash
# Basic route
curl http://localhost:8100/

# Path variable
curl http://localhost:8101/users/john

# Query params
curl "http://localhost:8101/search?q=golang&limit=20"

# Basic auth
curl -u admin:admin123 http://localhost:8081/basic/protected

# JWT
curl -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

### Using Browser

Just visit the URLs in your browser for GET endpoints!

---

## File Structure

```
mux-tutorial/
├── README.md                     # This file
├── GUIDE.md                      # Comprehensive guide
├── LEARNING_GUIDE.md             # Mux basics reference
├── main.go                       # Original tutorial
├── 01-basic-routes/
│   ├── README.md                 # Documentation for this folder
│   ├── main.go                   # Basic routes example
│   ├── go.mod
│   └── go.sum
├── 02-path-and-query-params/
│   ├── README.md
│   ├── main.go                   # Path/query example
│   ├── go.mod
│   └── go.sum
├── 03-env-config/
│   ├── README.md
│   ├── main.go                   # Env vars example
│   ├── .env.example
│   ├── go.mod
│   └── go.sum
├── 04-docker/
│   ├── README.md
│   ├── Dockerfile                # Multi-stage build
│   ├── docker-compose.yml        # App + DB + Redis
│   ├── .dockerignore
│   └── main.go
└── 05-authentication/
    ├── README.md
    ├── main.go                   # All auth methods
    ├── go.mod
    └── go.sum
```

---

## Each Folder Contains

- `README.md` - Complete documentation with:
  - File descriptions
  - How to run
  - All API endpoints with request/response examples
  - Code snippets
  - Best practices

- `main.go` - Working example code
- `go.mod` / `go.sum` - Dependencies

---

## Additional Resources

- [Gorilla Mux GitHub](https://github.com/gorilla/mux)
- [Go http package docs](https://pkg.go.dev/net/http)
- [Gorilla Handlers](https://github.com/gorilla/handlers) - Common middleware

---

## Auth Test Credentials

| Username | Password | Role |
|----------|----------|------|
| admin | admin123 | admin |
| user | user123 | user |

---

## Key Concepts Summary

### Routes
```go
r := mux.NewRouter()
r.HandleFunc("/path", Handler)
```

### Path Variables
```go
r.HandleFunc("/users/{id}", Handler)
vars := mux.Vars(r)
id := vars["id"]
```

### Query Parameters
```go
query := r.URL.Query()
value := query.Get("key")
```

### Middleware
```go
r.Use(MyMiddleware)
```

### Environment Variables
```go
port := os.Getenv("PORT")
```

---

## Next Steps

After completing this tutorial:

1. Build your own API combining concepts
2. Add a database (PostgreSQL with pgx, MongoDB with mgo)
3. Add Redis for caching
4. Implement WebSocket support
5. Add rate limiting
6. Set up CI/CD pipeline
7. Deploy to cloud (AWS, GCP, Azure)

---

## Troubleshooting

### Port already in use
```bash
# Find process using port
netstat -ano | findstr :8100

# Kill process
taskkill /PID <PID> /F
```

### Module not found
```bash
go mod tidy
go get github.com/gorilla/mux
```

### Docker issues
```bash
docker-compose down -v
docker-compose up --build
```

---

Happy coding! 🚀
