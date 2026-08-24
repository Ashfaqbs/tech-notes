# 01 - Basic Routes

Simple examples showing how to create basic routes with Gorilla Mux.

## Files in This Folder

| File | Purpose | Edit? |
|------|---------|-------|
| `main.go` | Main application with all route examples | ✅ You edit |
| `go.mod` | Go module definition (dependencies) | ✅ Via `go get` |
| `go.sum` | Dependency checksums (auto-generated) | ❌ Auto-generated |

### Understanding go.mod vs go.sum

**go.mod** - Declares what your project needs:
```go
module basic-routes

go 1.24

require github.com/gorilla/mux v1.8.1  // Your dependency
```

**go.sum** - Contains cryptographic hashes for verification:
```
github.com/gorilla/mux v1.8.1 h1:abcdef123...
github.com/gorilla/mux v1.8.1/go.mod h1:xyz789...
```

**Why two files?**
- `go.mod` = "I need this library"
- `go.sum` = "This is what the library should look like (hash verification)"

**How go.sum is referenced automatically:**
When you run `go run main.go`, Go:
1. Reads `go.mod` for dependencies
2. Verifies each package against `go.sum` hashes
3. If hash mismatch → ERROR (possible tampering)
4. If hash matches → proceeds with build

**You never manually reference go.sum** - Go tooling handles it automatically!

## How to Run

```bash
cd 01-basic-routes
go run main.go
```

Server starts on **http://localhost:8100**

---

## API Endpoints

### 1. `GET /` - Home Page

**Description:** Welcome page with timestamp

**Request:**
```bash
curl http://localhost:8100/
```

**Response:**
```json
{
  "message": "Welcome to Mux Basic Routes!",
  "path": "/",
  "timestamp": "2026-04-16T21:23:48+05:30"
}
```

---

### 2. `GET /hello` - Hello World

**Description:** Simple text response

**Request:**
```bash
curl http://localhost:8100/hello
```

**Response:**
```
Hello, World!
```

---

### 3. `GET /api/time` - Current Time

**Description:** Returns current server time as JSON

**Request:**
```bash
curl http://localhost:8100/api/time
```

**Response:**
```json
{
  "current_time": "2026-04-16T21:23:48Z",
  "unix": "1744840428"
}
```

---

### 4. `GET /html` - HTML Response

**Description:** Returns HTML content

**Request:**
```bash
curl http://localhost:8100/html
```

**Response:**
```html
<!DOCTYPE html>
<html>
<head><title>Mux Example</title></head>
<body>
	<h1>Hello from Gorilla Mux!</h1>
	<p>This is an HTML response.</p>
	<ul>
		<li><a href="/">Home</a></li>
		<li><a href="/hello">Hello</a></li>
		<li><a href="/api/time">Time API</a></li>
	</ul>
</body>
</html>
```

---

### 5. `GET /created` - 201 Created

**Description:** Demonstrates custom HTTP status code (201)

**Request:**
```bash
curl http://localhost:8100/created
```

**Response:**
```json
{
  "status": "created",
  "message": "Resource created successfully",
  "code": "201"
}
```

**HTTP Status:** `201 Created`

---

### 6. `GET /notfound` - 404 Not Found

**Description:** Demonstrates custom 404 response

**Request:**
```bash
curl http://localhost:8100/notfound
```

**Response:**
```json
{
  "status": "not_found",
  "message": "The requested resource was not found",
  "code": "404"
}
```

**HTTP Status:** `404 Not Found`

---

## Code Explanation

### Import Statements Explained

```go
import (
    "encoding/json"  // JSON encoding/decoding
                    // Used: json.NewEncoder(w).Encode(data) in JSONResponse()
    
    "fmt"           // Formatted I/O operations
                    // Used: fmt.Println(), fmt.Fprintf() for printing and formatting strings
    
    "log"           // Logging utilities
                    // Used: log.Fatal() to log errors and exit the program
    
    "net/http"      // HTTP server functionality
                    // Used: http.ResponseWriter, http.Request, http.ListenAndServe()
                    //       http.StatusCreated, http.StatusNotFound, http.StatusOK
    
    "time"          // Time operations
                    // Used: time.Now().Format() to get current timestamp
    
    "github.com/gorilla/mux"  // Gorilla Mux router (external package)
                              // Used: mux.NewRouter(), r.HandleFunc(), mux.Vars()
)
```

### Creating a Router

```go
import "github.com/gorilla/mux"

func main() {
    r := mux.NewRouter()
    
    // Register routes
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/hello", HelloHandler)
    
    // Start server
    http.ListenAndServe(":8100", r)
}
```

### Handler Function Pattern

```go
func HandlerName(w http.ResponseWriter, r *http.Request) {
    // w - write response
    // r - read request
    
    // Example: send JSON
    JSONResponse(w, map[string]string{
        "message": "Hello",
    })
}
```

### JSON Response Helper

```go
func JSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### Setting Custom Status Code

```go
func CreatedHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)  // 201
    json.NewEncoder(w).Encode(map[string]string{
        "status": "created",
    })
}
```

---

## Key Takeaways

1. **Routes are registered** with `r.HandleFunc(path, handler)`
2. **Handlers** follow signature `func(w http.ResponseWriter, r *http.Request)`
3. **Response type** is set via `w.Header().Set("Content-Type", "...")`
4. **Status codes** are set with `w.WriteHeader(http.Status...)`
5. **Router** implements `http.Handler`, so it works with standard `http.ListenAndServe`
