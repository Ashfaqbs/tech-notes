// 02 - Path Variables and Query Parameters
// Run with: go run main.go
// Test: See commands below

package main

import (
	"encoding/json"  // JSON encoding/decoding - used in JSONResponse() for json.NewEncoder(w).Encode(data)
	"fmt"           // Formatted I/O - used for fmt.Println(), fmt.Fprintf(), fmt.Sprintf() to format strings
	"log"           // Logging - used in log.Fatal() to log errors and exit
	"net/http"      // HTTP server - used for http.ResponseWriter, http.Request, http.ListenAndServe
	"strconv"       // String conversion - used for strconv.Atoi() to convert string to int (limit, page, etc.)
	"strings"       // String manipulation - used for strings.Join() to combine filter strings

	"github.com/gorilla/mux"  // Gorilla Mux router - used for mux.NewRouter(), mux.Vars(r) to extract path variables
)

func main() {
	r := mux.NewRouter()

	// =====================================================
	// PATH VARIABLES
	// =====================================================

	// 1. Single path variable - {name}
	r.HandleFunc("/users/{name}", GetUserByName)

	// 2. Path variable with regex - {id:[0-9]+} (only numbers)
	r.HandleFunc("/users/{id:[0-9]+}", GetUserByID)

	// 3. Multiple path variables
	r.HandleFunc("/articles/{category}/{id}", GetArticle)

	// 4. Path variable with custom regex (slug format)
	r.HandleFunc("/blog/{slug:[a-z0-9-]+}", GetBlogPost)

	// 5. Catch-all path variable
	r.HandleFunc("/files/{path:.*}", GetFile)

	// =====================================================
	// QUERY PARAMETERS
	// =====================================================

	// 6. Simple search with query params
	r.HandleFunc("/search", SearchHandler)

	// 7. Pagination example
	r.HandleFunc("/products", GetProducts)

	// 8. Filtering with multiple query params
	r.HandleFunc("/movies", GetMovies)

	// =====================================================
	// COMBINED: Path + Query
	// =====================================================

	// 9. Path variable + query params together
	r.HandleFunc("/users/{id}/posts", GetUserPosts)

	fmt.Println("=== Path Variables & Query Params Example ===")
	fmt.Println("Server: http://localhost:8101")
	fmt.Println("\n--- Path Variables ---")
	fmt.Println("  curl http://localhost:8101/users/john")
	fmt.Println("  curl http://localhost:8101/users/123")
	fmt.Println("  curl http://localhost:8101/articles/tech/42")
	fmt.Println("  curl http://localhost:8101/blog/my-first-post")
	fmt.Println("  curl http://localhost:8101/files/folder/subfolder/file.txt")
	fmt.Println("\n--- Query Parameters ---")
	fmt.Println("  curl 'http://localhost:8101/search?q=golang&limit=10'")
	fmt.Println("  curl 'http://localhost:8101/products?page=1&per_page=20&sort=price'")
	fmt.Println("  curl 'http://localhost:8101/movies?genre=action&year=2024&rating=8'")
	fmt.Println("\n--- Combined ---")
	fmt.Println("  curl 'http://localhost:8101/users/123/posts?limit=5'")

	log.Fatal(http.ListenAndServe(":8101", r))
}

// =====================================================
// PATH VARIABLE HANDLERS
// =====================================================

// GetUserByName - Extract single path variable
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	JSONResponse(w, map[string]interface{}{
		"handler": "GetUserByName",
		"name":    name,
		"message": fmt.Sprintf("Looking up user: %s", name),
	})
}

// GetUserByID - Path variable with regex (only numeric IDs)
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	JSONResponse(w, map[string]interface{}{
		"handler": "GetUserByID",
		"id":      id,
		"message": fmt.Sprintf("Looking up user with ID: %s", id),
	})
}

// GetArticle - Multiple path variables
func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	id := vars["id"]

	JSONResponse(w, map[string]interface{}{
		"handler":  "GetArticle",
		"category": category,
		"article_id": id,
		"message":  fmt.Sprintf("Article %s in category: %s", id, category),
	})
}

// GetBlogPost - Custom regex (slug format: lowercase, numbers, hyphens)
func GetBlogPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	JSONResponse(w, map[string]interface{}{
		"handler": "GetBlogPost",
		"slug":    slug,
		"message": fmt.Sprintf("Blog post: %s", slug),
	})
}

// GetFile - Catch-all path variable
func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]

	JSONResponse(w, map[string]interface{}{
		"handler": "GetFile",
		"path":    path,
		"message": fmt.Sprintf("File path: /%s", path),
	})
}

// =====================================================
// QUERY PARAMETER HANDLERS
// =====================================================

// SearchHandler - Basic query parameter extraction
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Get single value
	q := query.Get("q")
	limitStr := query.Get("limit")

	// Parse with default
	limit := 10
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	JSONResponse(w, map[string]interface{}{
		"handler": "SearchHandler",
		"query":   q,
		"limit":   limit,
		"all_params": query,
	})
}

// GetProducts - Pagination with query params
func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Parse pagination
	page, _ := strconv.Atoi(query.Get("page"))
	if page == 0 {
		page = 1
	}

	perPage, _ := strconv.Atoi(query.Get("per_page"))
	if perPage == 0 {
		perPage = 10
	}

	sort := query.Get("sort")
	if sort == "" {
		sort = "name"
	}

	JSONResponse(w, map[string]interface{}{
		"handler": "GetProducts",
		"pagination": map[string]int{
			"page":     page,
			"per_page": perPage,
		},
		"sort": sort,
		"products": []map[string]string{
			{"name": "Product A", "price": "$10"},
			{"name": "Product B", "price": "$20"},
			{"name": "Product C", "price": "$30"},
		},
	})
}

// GetMovies - Filtering with multiple query params
func GetMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	genre := query.Get("genre")
	year := query.Get("year")
	rating := query.Get("rating")

	// Build filter description
	filters := []string{}
	if genre != "" {
		filters = append(filters, fmt.Sprintf("genre=%s", genre))
	}
	if year != "" {
		filters = append(filters, fmt.Sprintf("year=%s", year))
	}
	if rating != "" {
		filters = append(filters, fmt.Sprintf("rating>=%s", rating))
	}

	JSONResponse(w, map[string]interface{}{
		"handler": "GetMovies",
		"filters": map[string]string{
			"genre":  genre,
			"year":   year,
			"rating": rating,
		},
		"filter_string": strings.Join(filters, ", "),
		"movies": []map[string]string{
			{"title": "Movie 1", "genre": "action", "year": "2024"},
			{"title": "Movie 2", "genre": "comedy", "year": "2023"},
			{"title": "Movie 3", "genre": "drama", "year": "2024"},
		},
	})
}

// =====================================================
// COMBINED HANDLER
// =====================================================

// GetUserPosts - Path variable + query params
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// Get path variable
	vars := mux.Vars(r)
	userID := vars["id"]

	// Get query params
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(query.Get("offset"))

	JSONResponse(w, map[string]interface{}{
		"handler": "GetUserPosts",
		"user_id": userID,
		"pagination": map[string]int{
			"limit":  limit,
			"offset": offset,
		},
		"posts": []map[string]string{
			{"title": "Post 1", "author_id": userID},
			{"title": "Post 2", "author_id": userID},
			{"title": "Post 3", "author_id": userID},
		},
	})
}

// =====================================================
// HELPER
// =====================================================

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
