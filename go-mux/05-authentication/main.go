// Authentication Examples: Basic Auth, JWT, API Key, OAuth
// Run with: go run main.go

package main

import (
	"context"        // Context for request context - used for context.WithValue() to store user/claims in request
	"encoding/base64"// Base64 encoding/decoding - used for decoding Basic Auth credentials (base64(user:pass))
	"encoding/json"  // JSON encoding/decoding - used in JSONResponse() for json.NewEncoder(w).Encode(data)
	"fmt"           // Formatted I/O - used for fmt.Println(), fmt.Sprintf() to format strings
	"log"           // Logging - used for log.Fatal() to log errors and exit
	"net/http"      // HTTP server - used for http.ResponseWriter, http.Request, http.ListenAndServe, http.StatusUnauthorized
	"os"            // Operating system - used for os.Getenv() to read JWT_SECRET, GITHUB_CLIENT_ID from env
	"strings"        // String manipulation - used for strings.SplitN(), strings.TrimPrefix() to parse auth headers
	"time"          // Time operations - used for time.Now().Add() to set JWT token expiration

	"github.com/golang-jwt/jwt/v5"  // JWT library - used for jwt.NewWithClaims(), jwt.Parse() to create/validate tokens
	"github.com/gorilla/mux"        // Gorilla Mux router - used for mux.NewRouter(), r.HandleFunc(), r.PathPrefix()
	"golang.org/x/crypto/bcrypt"    // Password hashing - used for bcrypt.GenerateFromPassword(), bcrypt.CompareHashAndPassword()
)

// =============================================================
// CONFIGURATION
// =============================================================

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	// In-memory user store (use a database in production!)
	users = map[string]User{
		"admin": {
			ID:       "1",
			Username: "admin",
			Password: hashPassword("admin123"),
			Role:     "admin",
		},
		"user": {
			ID:       "2",
			Username: "user",
			Password: hashPassword("user123"),
			Role:     "user",
		},
	}

	// API Keys (store hashed in production!)
	apiKeys = map[string]string{
		"sk-live-abc123xyz": "user-1",
		"sk-live-def456uvw": "user-2",
	}

	// OAuth state (in production, use Redis/database)
	oauthStates = make(map[string]string)
)

// =============================================================
// MODELS
// =============================================================

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Never expose password
	Role     string `json:"role"`
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// =============================================================
// MAIN
// =============================================================

func main() {
	r := mux.NewRouter()

	// =========================================================
	// PUBLIC ROUTES
	// =========================================================
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")
	r.HandleFunc("/auth/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/refresh", RefreshTokenHandler).Methods("POST")

	// =========================================================
	// BASIC AUTH ROUTES
	// =========================================================
	basicAuth := r.PathPrefix("/basic").Subrouter()
	basicAuth.Use(BasicAuthMiddleware)
	basicAuth.HandleFunc("/protected", BasicProtectedHandler)
	basicAuth.HandleFunc("/admin", AdminOnlyHandler)

	// =========================================================
	// JWT AUTH ROUTES
	// =========================================================
	jwtAuth := r.PathPrefix("/jwt").Subrouter()
	jwtAuth.Use(JWTAuthMiddleware)
	jwtAuth.HandleFunc("/protected", JWTProtectedHandler)
	jwtAuth.HandleFunc("/admin", AdminOnlyHandler)
	jwtAuth.HandleFunc("/profile", ProfileHandler)

	// =========================================================
	// API KEY AUTH ROUTES
	// =========================================================
	apiKeyAuth := r.PathPrefix("/api").Subrouter()
	apiKeyAuth.Use(APIKeyMiddleware)
	apiKeyAuth.HandleFunc("/v1/data", APIProtectedHandler)

	// =========================================================
	// OAUTH ROUTES (GitHub example)
	// =========================================================
	r.HandleFunc("/oauth/github", GitHubOAuthHandler)
	r.HandleFunc("/oauth/github/callback", GitHubOAuthCallbackHandler)

	// =========================================================
	// INFO ROUTES
	// =========================================================
	r.HandleFunc("/auth/info", AuthInfoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	fmt.Println("\n=== Authentication Examples ===")
	fmt.Println("\n1. BASIC AUTH:")
	fmt.Println("   curl -u admin:admin123 http://localhost:8000/basic/protected")
	fmt.Println("   curl -u user:user123 http://localhost:8000/basic/admin")
	fmt.Println("\n2. JWT AUTH:")
	fmt.Println("   # Login first:")
	fmt.Println("   curl -X POST http://localhost:8000/auth/login -d '{\"username\":\"admin\",\"password\":\"admin123\"}'")
	fmt.Println("   # Use token:")
	fmt.Println("   curl -H \"Authorization: Bearer <token>\" http://localhost:8000/jwt/protected")
	fmt.Println("\n3. API KEY:")
	fmt.Println("   curl -H \"X-API-Key: sk-live-abc123xyz\" http://localhost:8000/api/v1/data")
	fmt.Println("\n4. OAUTH (GitHub):")
	fmt.Println("   Visit: http://localhost:8000/oauth/github")
	fmt.Println("\n5. INFO:")
	fmt.Println("   curl http://localhost:8000/auth/info")
	fmt.Println("\nDefault users:")
	fmt.Println("   admin / admin123 (admin role)")
	fmt.Println("   user / user123 (user role)")

	log.Fatal(http.ListenAndServe(":"+port, r))
}

// =============================================================
// PUBLIC HANDLERS
// =============================================================

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]string{
		"message": "Authentication Examples API",
		"docs":    "/auth/info",
	})
}

func AuthInfoHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, map[string]interface{}{
		"authentication_methods": map[string]interface{}{
			"basic_auth": map[string]interface{}{
				"description": "Simple username/password via HTTP Authorization header",
				"usage":       "curl -u username:password /basic/protected",
				"pros":        "Simple, built into HTTP",
				"cons":        "Credentials sent with every request, base64 encoded (not encrypted)",
				"best_for":    "Internal services, simple APIs",
			},
			"jwt": map[string]interface{}{
				"description": "JSON Web Token - stateless, signed tokens",
				"usage":       "POST /auth/login, then use Bearer token in Authorization header",
				"pros":        "Stateless, can include claims, widely supported",
				"cons":        "Cannot revoke without blacklist, token size",
				"best_for":    "REST APIs, microservices, mobile apps",
			},
			"api_key": map[string]interface{}{
				"description": "Long-lived key for API access",
				"usage":       "Include X-API-Key header with requests",
				"pros":        "Simple, easy to rotate, good for server-to-server",
				"cons":        "No user context, long-lived = higher risk",
				"best_for":    "Server-to-server, third-party integrations",
			},
			"oauth": map[string]interface{}{
				"description": "Delegated authorization via third parties",
				"usage":       "Redirect to OAuth provider, callback with code, exchange for token",
				"pros":        "Don't handle passwords, social login, fine-grained scopes",
				"cons":        "Complex, depends on provider",
				"best_for":    "Social login, third-party app access",
			},
		},
		"test_credentials": map[string]string{
			"admin": "admin / admin123 (admin role)",
			"user":  "user / user123 (user role)",
		},
	})
}

// =============================================================
// BASIC AUTHENTICATION
// =============================================================

// BasicAuthMiddleware validates HTTP Basic Auth
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Parse "Basic base64(username:password)"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Decode base64
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "Invalid base64 encoding", http.StatusUnauthorized)
			return
		}

		// Split username:password
		creds := strings.SplitN(string(decoded), ":", 2)
		if len(creds) != 2 {
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		username, password := creds[0], creds[1]

		// Validate against user store
		user, exists := users[username]
		if !exists || !checkPassword(password, user.Password) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Add user to context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func BasicProtectedHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	JSONResponse(w, map[string]interface{}{
		"message": "Access granted!",
		"user":    user,
	})
}

// =============================================================
// JWT AUTHENTICATION
// =============================================================

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, exists := users[req.Username]
	if !exists || !checkPassword(req.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	JSONResponse(w, map[string]interface{}{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   3600, // 1 hour
		"user": map[string]string{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

func generateJWT(user User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "mux-auth-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTAuthMiddleware validates JWT tokens
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Parse "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JWTProtectedHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*Claims)
	JSONResponse(w, map[string]interface{}{
		"message": "Access granted!",
		"claims":  claims,
	})
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// In production, validate refresh token from request
	// For demo, just show the concept
	JSONResponse(w, map[string]string{
		"message": "Send refresh token to get new access token",
	})
}

// =============================================================
// API KEY AUTHENTICATION
// =============================================================

// APIKeyMiddleware validates API keys
func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "Missing X-API-Key header", http.StatusUnauthorized)
			return
		}

		userID, exists := apiKeys[apiKey]
		if !exists {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		// Add user ID to context
		ctx := context.WithValue(r.Context(), "api_user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func APIProtectedHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("api_user_id").(string)
	JSONResponse(w, map[string]interface{}{
		"message": "API access granted!",
		"user_id": userID,
		"data":    []string{"item1", "item2", "item3"},
	})
}

// =============================================================
// OAUTH 2.0 (GitHub Example)
// =============================================================

func GitHubOAuthHandler(w http.ResponseWriter, r *http.Request) {
	// In production, get these from env
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		http.Error(w, "GitHub OAuth not configured", http.StatusInternalServerError)
		return
	}

	// Generate random state for CSRF protection
	state := generateRandomString(32)
	oauthStates[state] = state // Store for validation

	redirectURI := os.Getenv("REDIRECT_URI")
	if redirectURI == "" {
		redirectURI = "http://localhost:8081/oauth/github/callback"
	}
	authURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s",
		clientID,
		redirectURI,
		state,
	)

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func GitHubOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	// Validate state (CSRF protection)
	if _, exists := oauthStates[state]; !exists {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}
	delete(oauthStates, state)

	// Exchange code for token
	// (Implementation would call GitHub API)
	JSONResponse(w, map[string]string{
		"message": "OAuth callback received",
		"code":    code[:10] + "...",
		"note":    "In production, exchange this code for access token",
	})
}

// =============================================================
// HELPER FUNCTIONS
// =============================================================

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, exists := users[req.Username]; exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	users[req.Username] = User{
		ID:       fmt.Sprintf("%d", time.Now().UnixNano()),
		Username: req.Username,
		Password: hashPassword(req.Password),
		Role:     "user",
	}

	JSONResponse(w, map[string]string{
		"message": "User registered successfully",
		"username": req.Username,
	})
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte('a' + (i % 26))
	}
	return string(b)
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// =============================================================
// ROLE-BASED AUTHORIZATION
// =============================================================

// AdminOnlyHandler - requires admin role
func AdminOnlyHandler(w http.ResponseWriter, r *http.Request) {
	// Check if from Basic Auth
	if user, ok := r.Context().Value("user").(User); ok {
		if user.Role != "admin" {
			http.Error(w, "Admin access required", http.StatusForbidden)
			return
		}
		JSONResponse(w, map[string]interface{}{
			"message": "Admin access granted!",
			"user":    user,
		})
		return
	}

	// Check if from JWT Auth
	if claims, ok := r.Context().Value("claims").(*Claims); ok {
		if claims.Role != "admin" {
			http.Error(w, "Admin access required", http.StatusForbidden)
			return
		}
		JSONResponse(w, map[string]interface{}{
			"message": "Admin access granted!",
			"claims":  claims,
		})
		return
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*Claims)
	JSONResponse(w, map[string]interface{}{
		"user_id":    claims.UserID,
		"username":   claims.Username,
		"role":       claims.Role,
		"expires_at": claims.ExpiresAt,
	})
}
