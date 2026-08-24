# Gorilla Mux - Complete Learning Guide

## Table of Contents
1. [What is Gorilla Mux?](#what-is-gorilla-mux)
2. [Installation](#installation)
3. [Basic Usage](#basic-usage)
4. [Path Variables](#path-variables)
5. [Query Parameters](#query-parameters)
6. [HTTP Methods](#http-methods)
7. [Subrouters](#subrouters)
8. [Middleware](#middleware)
9. [Complete Examples](#complete-examples)

---

## What is Gorilla Mux?

Gorilla Mux is a powerful HTTP request router (multiplexer) for Go. It matches incoming requests against registered routes and calls the appropriate handler.

**Key Features:**
- URL pattern matching with variables (e.g., `/users/{id}`)
- Regular expression support in patterns
- HTTP method restriction
- Header, query, and host matching
- Subrouting for modular applications
- Middleware support

---

## Installation

```bash
go get -u github.com/gorilla/mux
```

---

## Basic Usage

### Creating a Router

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Register routes
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/hello", HelloHandler)

    // Start server
    http.ListenAndServe(":8000", r)
}

// Handler functions
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome Home!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

### Handler Signature

All handlers follow this signature:
```go
func HandlerName(w http.ResponseWriter, r *http.Request) {
    // w - write response
    // r - read request
}
```

---

## Path Variables

Path variables let you capture dynamic parts of the URL.

### Basic Path Variable

```go
// Matches: /users/john, /users/alice, /users/any-value
r.HandleFunc("/users/{name}", GetUserHandler)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // Extract variables
    vars := mux.Vars(r)
    name := vars["name"]
    
    fmt.Fprintf(w, "User: %s", name)
}
```

### Path Variable with Regex

```go
// Only matches numeric IDs: /users/123 (NOT /users/john)
r.HandleFunc("/users/{id:[0-9]+}", GetUserByIDHandler)

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]  // "123"
    
    // Convert to int if needed
    userID, _ := strconv.Atoi(id)
    
    fmt.Fprintf(w, "User ID: %d", userID)
}
```

### Multiple Path Variables

```go
// Matches: /articles/tech/42, /articles/news/100
r.HandleFunc("/articles/{category}/{id:[0-9]+}", GetArticleHandler)

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]  // "tech"
    id := vars["id"]              // "42"
    
    fmt.Fprintf(w, "Category: %s, Article ID: %s", category, id)
}
```

### Common Regex Patterns

| Pattern | Matches | Example |
|---------|---------|---------|
| `{id}` | Anything | `/users/john` |
| `{id:[0-9]+}` | Numbers only | `/users/123` |
| `{slug:[a-z-]+}` | Lowercase slugs | `/blog/my-post` |
| `{path:.*}` | Everything (greedy) | `/files/a/b/c` |

---

## Query Parameters

Query parameters come after `?` in the URL: `/search?q=golang&limit=10`

```go
r.HandleFunc("/search", SearchHandler)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    // Get query parameters
    query := r.URL.Query()
    
    // Get single value
    searchTerm := query.Get("q")      // "golang"
    limitStr := query.Get("limit")    // "10"
    
    // Get multiple values (e.g., ?tags=go&tags=rust)
    tags := query["tags"]             // ["go", "rust"]
    
    // Parse with default
    limit := 10
    if limitStr != "" {
        limit, _ = strconv.Atoi(limitStr)
    }
    
    fmt.Fprintf(w, "Search: %s, Limit: %d, Tags: %v", searchTerm, limit, tags)
}
```

### Query Parameter Examples

| URL | `query.Get("q")` | `query.Get("limit")` | `query["tags"]` |
|-----|------------------|----------------------|-----------------|
| `/search?q=go` | `"go"` | `""` | `[]` |
| `/search?q=go&limit=20` | `"go"` | `"20"` | `[]` |
| `/search?tags=go&tags=rust` | `""` | `""` | `["go", "rust"]` |

---

## HTTP Methods

Restrict routes to specific HTTP methods.

```go
// GET /api/items
r.HandleFunc("/api/items", GetItemsHandler).Methods("GET")

// POST /api/items
r.HandleFunc("/api/items", CreateItemHandler).Methods("POST")

// PUT /api/items/{id}
r.HandleFunc("/api/items/{id}", UpdateItemHandler).Methods("PUT")

// DELETE /api/items/{id}
r.HandleFunc("/api/items/{id}", DeleteItemHandler).Methods("DELETE")
```

### Handler Examples

```go
// GET - List items
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
    items := []string{"item1", "item2", "item3"}
    JSONResponse(w, map[string]interface{}{"items": items})
}

// POST - Create item
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
    var item map[string]interface{}
    json.NewDecoder(r.Body).Decode(&item)
    
    // Save item...
    
    w.WriteHeader(http.StatusCreated)
    JSONResponse(w, map[string]interface{}{
        "status": "created",
        "item": item,
    })
}

// PUT - Update item
func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    var item map[string]interface{}
    json.NewDecoder(r.Body).Decode(&item)
    
    // Update item...
    
    JSONResponse(w, map[string]interface{}{
        "id": id,
        "item": item,
    })
}

// DELETE - Delete item
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    // Delete item...
    
    JSONResponse(w, map[string]interface{}{
        "id": id,
        "status": "deleted",
    })
}
```

---

## Subrouters

Subrouters let you group routes with common conditions.

### Path Prefix Subrouter

```go
// All routes under /api/v1
api := r.PathPrefix("/api/v1").Subrouter()

api.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
    JSONResponse(w, map[string]string{"status": "ok", "version": "1"})
})

api.HandleFunc("/users", GetUsersHandler)
api.HandleFunc("/users/{id}", GetUserHandler)

// Full paths:
// GET /api/v1/status
// GET /api/v1/users
// GET /api/v1/users/123
```

### Host Subrouter

```go
// Admin routes only match admin.localhost
admin := r.Host("admin.localhost").Subrouter()

admin.HandleFunc("/", AdminHomeHandler)
admin.HandleFunc("/dashboard", AdminDashboardHandler)

// Matches:
// http://admin.localhost/
// http://admin.localhost/dashboard
```

### Combined Subrouter

```go
// API routes with both path prefix AND host
api := r.Host("api.example.com").PathPrefix("/v1").Subrouter()

api.HandleFunc("/users", GetUsersHandler).Methods("GET")
api.HandleFunc("/users", CreateUserHandler).Methods("POST")
```

---

## Middleware

Middleware wraps handlers to add cross-cutting concerns like logging, auth, etc.

### Applying Middleware

```go
// Apply to all routes
r.Use(LoggingMiddleware)
r.Use(AuthMiddleware)

// Apply to specific subrouter
api := r.PathPrefix("/api").Subrouter()
api.Use(AuthMiddleware)  // Only for /api/* routes
api.HandleFunc("/data", DataHandler)
```

### Middleware Pattern

```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // BEFORE handler
        fmt.Println("Before:", r.URL.Path)
        
        // Call next handler
        next.ServeHTTP(w, r)
        
        // AFTER handler
        fmt.Println("After:", r.URL.Path)
    })
}
```

### Logging Middleware

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Call handler
        next.ServeHTTP(w, r)
        
        // Log after
        log.Printf("[%s] %s %s (%v)", 
            r.Method, 
            r.URL.Path, 
            r.RemoteAddr,
            time.Since(start),
        )
    })
}
```

### Authentication Middleware

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Validate token...
        if !isValidToken(token) {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        // Token valid, continue
        next.ServeHTTP(w, r)
    })
}

func isValidToken(token string) bool {
    // Your validation logic
    return token == "secret-token"
}
```

### Recovery Middleware (Panic Handler)

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

---

## Complete Examples

### JSON Response Helper

```go
func JSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### Full CRUD API Example

```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

// In-memory store (use a database in production!)
var items = make(map[string]string)

func main() {
    r := mux.NewRouter()
    
    // Routes
    r.HandleFunc("/items", GetItems).Methods("GET")
    r.HandleFunc("/items/{id}", GetItem).Methods("GET")
    r.HandleFunc("/items", CreateItem).Methods("POST")
    r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
    r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
    
    http.ListenAndServe(":8000", r)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
    JSONResponse(w, items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    item, exists := items[id]
    if !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }
    
    JSONResponse(w, map[string]string{"id": id, "value": item})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
    var data map[string]string
    json.NewDecoder(r.Body).Decode(&data)
    
    id := data["id"]
    value := data["value"]
    
    items[id] = value
    
    w.WriteHeader(http.StatusCreated)
    JSONResponse(w, map[string]string{"id": id, "value": value})
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    var data map[string]string
    json.NewDecoder(r.Body).Decode(&data)
    
    items[id] = data["value"]
    
    JSONResponse(w, map[string]string{"id": id, "value": data["value"]})
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    delete(items, id)
    
    JSONResponse(w, map[string]string{"status": "deleted", "id": id})
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### Testing with curl

```bash
# Get all items
curl http://localhost:8000/items

# Get single item
curl http://localhost:8000/items/123

# Create item
curl -X POST http://localhost:8000/items \
  -H "Content-Type: application/json" \
  -d '{"id":"123","value":"test"}'

# Update item
curl -X PUT http://localhost:8000/items/123 \
  -H "Content-Type: application/json" \
  -d '{"value":"updated"}'

# Delete item
curl -X DELETE http://localhost:8000/items/123
```

---

## Quick Reference

### Route Matchers

```go
// Path
r.Path("/exact/path")
r.PathPrefix("/prefix/")

// Host
r.Host("example.com")
r.Host("{subdomain}.example.com")

// Methods
r.Methods("GET", "POST")

// Headers
r.Headers("X-Requested-With", "XMLHttpRequest")
r.HeadersRegexp("Content-Type", "application/json")

// Queries
r.Queries("key", "value")
r.Queries("key", "{value:[0-9]+}")

// Schemes
r.Schemes("https")

// Custom matcher
r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
    return r.ProtoMajor == 0
})

// Combine matchers
r.HandleFunc("/api", Handler).
    Host("api.example.com").
    Methods("GET").
    Schemes("https")
```

### Getting Route Info

```go
// Get path variables
vars := mux.Vars(r)
id := vars["id"]

// Get query params
query := r.URL.Query()
q := query.Get("q")

// Get headers
auth := r.Header.Get("Authorization")

// Get request body
var data MyStruct
json.NewDecoder(r.Body).Decode(&data)

// Get context value (set by middleware)
user := r.Context().Value("user")
```

---

## Next Steps

1. **Try the tutorial**: Run `go run main.go` in this directory
2. **Modify handlers**: Change responses, add new routes
3. **Add a database**: Replace in-memory storage with PostgreSQL/MongoDB
4. **Add authentication**: Implement JWT or session-based auth
5. **Add validation**: Validate request inputs before processing

---

## Resources

- Official Docs: https://github.com/gorilla/mux
- Go HTTP Package: https://pkg.go.dev/net/http
- Gorilla Middleware: https://github.com/gorilla/handlers
