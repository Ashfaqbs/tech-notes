// 01 - Basic Routes Example
// Run with: go run main.go
// Test: curl http://localhost:8100/

package main

import (
	"encoding/json"  // JSON encoding/decoding - used in JSONResponse() for json.NewEncoder(w).Encode(data)
	"fmt"           // Formatted I/O - used for fmt.Println() and fmt.Fprintf() to format strings
	"log"           // Logging - used in log.Fatal() to log errors and exit
	"net/http"      // HTTP server - used for http.ResponseWriter, http.Request, http.ListenAndServe, http.StatusCreated, etc.
	"time"          // Time operations - used for time.Now().Format() to get current timestamp

	"github.com/gorilla/mux"  // Gorilla Mux router - used for mux.NewRouter() and route handling
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// =====================================================
	// 1. Simple route - exact path match
	// =====================================================
	r.HandleFunc("/", HomeHandler)

	// =====================================================
	// 2. Static route - fixed path
	// =====================================================
	r.HandleFunc("/hello", HelloHandler)

	// =====================================================
	// 3. Route with JSON response
	// =====================================================
	r.HandleFunc("/api/time", TimeHandler)

	// =====================================================
	// 4. Route that returns HTML
	// =====================================================
	r.HandleFunc("/html", HTMLHandler)

	// =====================================================
	// 5. Route with custom status code
	// =====================================================
	r.HandleFunc("/created", CreatedHandler)
	r.HandleFunc("/notfound", NotFoundHandler)

	// Start server
	fmt.Println("=== Basic Routes Example ===")
	fmt.Println("Server: http://localhost:8100")
	fmt.Println("\nTry these:")
	fmt.Println("  curl http://localhost:8100/")
	fmt.Println("  curl http://localhost:8100/hello")
	fmt.Println("  curl http://localhost:8100/api/time")
	fmt.Println("  curl http://localhost:8100/html")
	fmt.Println("  curl http://localhost:8100/created")
	fmt.Println("  curl http://localhost:8100/notfound")

	log.Fatal(http.ListenAndServe(":8100", r))
}

// HomeHandler - Welcome page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]string{
		"message":   "Welcome to Mux Basic Routes!",
		"path":      r.URL.Path,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// HelloHandler - Simple text response
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World!\n")
}

// TimeHandler - Returns current time as JSON
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]string{
		"current_time": time.Now().Format(time.RFC3339),
		"unix":         fmt.Sprintf("%d", time.Now().Unix()),
	})
}

// HTMLHandler - Returns HTML response
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
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
`)
}

// CreatedHandler - Returns 201 Created status
func CreatedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "created",
		"message": "Resource created successfully",
		"code":    "201",
	})
}

// NotFoundHandler - Returns 404 Not Found
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "not_found",
		"message": "The requested resource was not found",
		"code":    "404",
	})
}

// JSONResponse - Helper to send JSON responses
func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
