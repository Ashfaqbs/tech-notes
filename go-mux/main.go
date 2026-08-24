// Gorilla Mux Tutorial - Complete Example
// Run with: go run main.go
// Then visit: http://localhost:8000

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// =====================================================
	// 1. BASIC ROUTES - Simple path matching
	// =====================================================
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/hello", HelloHandler)

	// =====================================================
	// 2. PATH VARIABLES - Dynamic segments in URL
	// =====================================================
	// {name} - matches anything until next slash
	r.HandleFunc("/users/{name}", GetUserHandler)

	// {id:[0-9]+} - matches only numbers (regex pattern)
	r.HandleFunc("/users/{id:[0-9]+}", GetUserByIDHandler)

	// Multiple path variables
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", GetArticleHandler)

	// =====================================================
	// 3. QUERY PARAMETERS - URL params after ?
	// =====================================================
	r.HandleFunc("/search", SearchHandler)
	r.HandleFunc("/products", GetProductsHandler)

	// =====================================================
	// 4. HTTP METHODS - Different handlers for GET/POST/PUT/DELETE
	// =====================================================
	// Method-specific routes
	r.HandleFunc("/api/items", GetItemsHandler).Methods("GET")
	r.HandleFunc("/api/items", CreateItemHandler).Methods("POST")
	r.HandleFunc("/api/items/{id}", UpdateItemHandler).Methods("PUT")
	r.HandleFunc("/api/items/{id}", DeleteItemHandler).Methods("DELETE")

	// =====================================================
	// 5. PATH PREFIX - Group routes under a prefix
	// =====================================================
	// All /api/v1/* routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, map[string]string{"status": "ok", "version": "1"})
	})
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, map[string]string{"health": "healthy"})
	})

	// =====================================================
	// 6. HOST MATCHING - Different handlers for different domains
	// =====================================================
	admin := r.Host("admin.localhost").Subrouter()
	admin.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, map[string]string{"area": "admin"})
	})

	// =====================================================
	// 7. MIDDLEWARE - Apply to all routes
	// =====================================================
	r.Use(LoggingMiddleware)
	r.Use(RecoveryMiddleware)

	// Start server
	fmt.Println("Server starting on http://localhost:8000")
	fmt.Println("\nTry these URLs:")
	fmt.Println("  GET  http://localhost:8000/")
	fmt.Println("  GET  http://localhost:8000/hello")
	fmt.Println("  GET  http://localhost:8000/users/john")
	fmt.Println("  GET  http://localhost:8000/users/123")
	fmt.Println("  GET  http://localhost:8000/articles/tech/42")
	fmt.Println("  GET  http://localhost:8000/search?q=golang&limit=10")
	fmt.Println("  GET  http://localhost:8000/api/items")
	fmt.Println("  POST http://localhost:8000/api/items")
	fmt.Println("  GET  http://localhost:8000/api/v1/status")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// =====================================================
// HANDLER FUNCTIONS
// =====================================================

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]string{
		"message": "Welcome to Gorilla Mux Tutorial!",
		"path":    r.URL.Path,
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]string{
		"greeting": "Hello, World!",
	})
}

// GetUserHandler - demonstrates path variable extraction
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path variables using mux.Vars()
	vars := mux.Vars(r)
	name := vars["name"]

	JSONResponse(w, map[string]string{
		"handler": "GetUserHandler",
		"name":    name,
	})
}

// GetUserByIDHandler - demonstrates regex path variable
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // This will only match numeric IDs

	JSONResponse(w, map[string]string{
		"handler": "GetUserByIDHandler",
		"user_id": id,
	})
}

// GetArticleHandler - demonstrates multiple path variables
func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	id := vars["id"]

	JSONResponse(w, map[string]string{
		"handler":  "GetArticleHandler",
		"category": category,
		"article_id": id,
	})
}

// SearchHandler - demonstrates query parameter extraction
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters using r.URL.Query()
	query := r.URL.Query()
	searchTerm := query.Get("q")     // Get single value
	limitStr := query.Get("limit")   // Get single value
	tags := query["tags"]            // Get multiple values (e.g., ?tags=go&tags=rust)

	// Parse limit with default value
	limit := 10
	if limitStr != "" {
		fmt.Sscanf(limitStr, "%d", &limit)
	}

	JSONResponse(w, map[string]interface{}{
		"handler":    "SearchHandler",
		"query":      searchTerm,
		"limit":      limit,
		"tags":       tags,
		"all_params": query,
	})
}

// GetProductsHandler - demonstrates filtering with query params
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	category := query.Get("category")
	minPrice := query.Get("min_price")
	maxPrice := query.Get("max_price")
	sort := query.Get("sort")

	// Build response based on filters
	response := map[string]interface{}{
		"handler": "GetProductsHandler",
		"filters": map[string]string{
			"category": category,
			"min_price": minPrice,
			"max_price": maxPrice,
			"sort":      sort,
		},
	}

	JSONResponse(w, response)
}

// GetItemsHandler - GET endpoint
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]interface{}{
		"action": "list_items",
		"items":  []string{"item1", "item2", "item3"},
	})
}

// CreateItemHandler - POST endpoint
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var item map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		// If no body provided, create a default item
		item = map[string]interface{}{"name": "default_item"}
	}

	JSONResponseWithStatus(w, map[string]interface{}{
		"action": "create_item",
		"item":   item,
		"status": "created",
	}, http.StatusCreated)
}

// UpdateItemHandler - PUT endpoint
func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item map[string]interface{}
	json.NewDecoder(r.Body).Decode(&item)

	JSONResponse(w, map[string]interface{}{
		"action": "update_item",
		"id":     id,
		"item":   item,
	})
}

// DeleteItemHandler - DELETE endpoint
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	JSONResponse(w, map[string]interface{}{
		"action": "delete_item",
		"id":     id,
		"status": "deleted",
	})
}

// =====================================================
// MIDDLEWARE FUNCTIONS
// =====================================================

// LoggingMiddleware - logs every request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log before handling
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// RecoveryMiddleware - recovers from panics
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// =====================================================
// HELPER FUNCTIONS
// =====================================================

// JSONResponse sends a JSON response with status 200
func JSONResponse(w http.ResponseWriter, data interface{}) {
	JSONResponseWithStatus(w, data, http.StatusOK)
}

// JSONResponseWithStatus sends a JSON response with custom status
func JSONResponseWithStatus(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
