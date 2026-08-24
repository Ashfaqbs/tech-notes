# Complete Mux Tutorial Guide

## Table of Contents

1. [Environment Variables](#1-environment-variables)
2. [Docker Containerization](#2-docker-containerization)
3. [Authentication](#3-authentication)

---

## 1. Environment Variables

### Why Use Environment Variables?

- Different config for dev/staging/prod
- Keep secrets out of code
- Follow 12-factor app methodology

### How to Read Environment Variables

```go
package main

import (
    "os"
)

func main() {
    // Simple - returns empty string if not set
    port := os.Getenv("PORT")
    
    // With default value
    port := getEnv("PORT", "8000")
    
    // As integer
    timeout := getEnvInt("TIMEOUT", 30)
}

// Helper function with default
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// Helper for integers
func getEnvInt(key string, defaultValue int) int {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    i, err := strconv.Atoi(value)
    if err != nil {
        return defaultValue
    }
    return i
}
```

### Using .env Files (Local Development)

Install godotenv:
```bash
go get github.com/joho/godotenv
```

Load in code:
```go
import "github.com/joho/godotenv"

func init() {
    // Load .env file (optional, won't fail if missing)
    _ = godotenv.Load()
}
```

Create `.env` file:
```bash
PORT=8000
ENV=development
DATABASE_URL=postgres://user:pass@localhost:5432/mydb
JWT_SECRET=your-secret-key
```

### Run with Environment Variables

```bash
# Inline
PORT=9000 ENV=production go run main.go

# From .env file (via godotenv)
go run main.go

# Export first
export PORT=9000
export ENV=production
go run main.go
```

### Example Project Structure

```
03-env-config/
├── main.go           # Application code
├── .env.example      # Template (commit this)
├── .env              # Actual values (DO NOT COMMIT)
├── go.mod
└── go.sum
```

---

## 2. Docker Containerization

### Dockerfile Explained

```dockerfile
# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Runtime
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]
```

### Multi-Stage Build Benefits

| Aspect | Single Stage | Multi-Stage |
|--------|--------------|-------------|
| Image Size | ~800MB | ~15MB |
| Build Tools | Included | Not included |
| Security | More attack surface | Minimal |

### Build and Run

```bash
# Build image
docker build -t mux-app .

# Run container
docker run -p 8000:8000 \
  -e PORT=8000 \
  -e ENV=production \
  -e DATABASE_URL="postgres://..." \
  mux-app
```

### Using Docker Compose

```bash
# Start all services
docker-compose up

# Start in background
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop all services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### docker-compose.yml Services

| Service | Image | Purpose |
|---------|-------|---------|
| api | (build) | Go application |
| db | postgres:16-alpine | PostgreSQL database |
| redis | redis:7-alpine | Cache/sessions |

### Environment Variables in Docker

```yaml
services:
  api:
    environment:
      - PORT=8000
      - ENV=production
      - DATABASE_URL=postgres://user:pass@db:5432/mydb
    # OR use .env file
    env_file:
      - .env.production
```

### .dockerignore

Always create `.dockerignore` to exclude:
- `.git/`
- `.env` files (secrets!)
- `vendor/`
- Build artifacts
- IDE files

---

## 3. Authentication

### Authentication Methods Comparison

| Method | Security | Complexity | Best For |
|--------|----------|------------|----------|
| Basic Auth | Low | Very Low | Internal services |
| JWT | High | Medium | REST APIs, mobile |
| API Key | Medium | Low | Server-to-server |
| OAuth 2.0 | High | High | Social login, 3rd party |

---

### 1. Basic Authentication

**How it works:**
- Client sends `Authorization: Basic base64(username:password)`
- Server decodes and validates credentials

**Implementation:**
```go
func BasicAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        
        // Parse "Basic base64"
        parts := strings.SplitN(authHeader, " ", 2)
        decoded, _ := base64.StdEncoding.DecodeString(parts[1])
        creds := strings.SplitN(string(decoded), ":", 2)
        username, password := creds[0], creds[1]
        
        // Validate
        user, exists := users[username]
        if !exists || !checkPassword(password, user.Password) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Add to context and continue
        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Usage
r.Use(BasicAuthMiddleware)
```

**Test:**
```bash
curl -u admin:admin123 http://localhost:8000/basic/protected
```

**Pros:**
- Built into HTTP
- Simple to implement

**Cons:**
- Credentials sent with every request
- Base64 is not encryption (only encoding)
- No token expiration

---

### 2. JWT (JSON Web Tokens)

**How it works:**
1. Client POSTs credentials to `/auth/login`
2. Server validates and returns signed JWT
3. Client includes `Authorization: Bearer <token>` in subsequent requests
4. Server validates token signature

**Token Structure:**
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.    # Header (algorithm)
eyJ1c2VyX2lkIjoiMSIsInJvbGUiOiJhZG1pbiJ9.  # Payload (claims)
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c  # Signature
```

**Implementation:**
```go
import "github.com/golang-jwt/jwt/v5"

// Generate token
func generateJWT(user User) (string, error) {
    claims := jwt.MapClaims{
        "user_id": user.ID,
        "username": user.Username,
        "role": user.Role,
        "exp": time.Now().Add(1 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Validate middleware
func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

**Test:**
```bash
# Login
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'

# Use token
curl -H "Authorization: Bearer <token>" http://localhost:8000/jwt/protected
```

**Pros:**
- Stateless (no database lookup)
- Can include user data in claims
- Expiration built-in

**Cons:**
- Cannot revoke without blacklist
- Larger than session tokens

---

### 3. API Key Authentication

**How it works:**
- Server generates long-lived key
- Client sends `X-API-Key: <key>` header
- Server validates key against store

**Implementation:**
```go
func APIKeyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        
        userID, exists := apiKeys[apiKey]
        if !exists {
            http.Error(w, "Invalid API key", http.StatusUnauthorized)
            return
        }
        
        ctx := context.WithValue(r.Context(), "api_user_id", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

**Test:**
```bash
curl -H "X-API-Key: sk-live-abc123xyz" http://localhost:8000/api/v1/data
```

**Pros:**
- Simple to implement
- Easy to rotate/revoke
- Good for server-to-server

**Cons:**
- No user context (just identifies the key owner)
- Long-lived = higher risk if compromised

**Best Practices:**
- Prefix keys: `sk-live-`, `sk-test-`
- Hash keys before storing (like passwords)
- Allow key rotation

---

### 4. OAuth 2.0

**How it works (Authorization Code Flow):**
1. Redirect user to provider (Google, GitHub, etc.)
2. User authorizes your app
3. Provider redirects back with `code`
4. Exchange `code` for `access_token`
5. Use token to access user info

**Implementation (GitHub):**
```go
func GitHubOAuthHandler(w http.ResponseWriter, r *http.Request) {
    state := generateRandomString(32)
    oauthStates[state] = state  // CSRF protection
    
    authURL := fmt.Sprintf(
        "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s",
        os.Getenv("GITHUB_CLIENT_ID"),
        "http://localhost:8000/oauth/github/callback",
        state,
    )
    http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func GitHubOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    
    // Exchange code for token
    resp, _ := http.Post(
        fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
            os.Getenv("GITHUB_CLIENT_ID"),
            os.Getenv("GITHUB_CLIENT_SECRET"),
            code,
        ),
        "application/json", nil,
    )
    
    // Parse response, get user info
}
```

**Test:**
```bash
# Visit in browser
http://localhost:8000/oauth/github
```

**Pros:**
- Don't handle passwords
- Social login convenience
- Fine-grained scopes

**Cons:**
- Complex implementation
- Depends on provider uptime
- Multiple providers = more complexity

---

### Password Hashing

Always hash passwords before storing:

```go
import "golang.org/x/crypto/bcrypt"

// Hash password
func hashPassword(password string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash)
}

// Verify password
func checkPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

**NEVER:**
- Store plain text passwords
- Use MD5/SHA1 (too fast, vulnerable to rainbow tables)
- Roll your own crypto

**ALWAYS:**
- Use bcrypt, argon2, or scrypt
- Use unique salt per password (bcrypt does this automatically)

---

### Role-Based Access Control (RBAC)

```go
func AdminOnlyHandler(w http.ResponseWriter, r *http.Request) {
    // From JWT claims
    claims := r.Context().Value("claims").(*Claims)
    
    if claims.Role != "admin" {
        http.Error(w, "Admin access required", http.StatusForbidden)
        return
    }
    
    // Admin logic here
}
```

---

## Quick Reference

### Environment Variables
```go
os.Getenv("KEY")
godotenv.Load()
```

### Docker
```bash
docker build -t app .
docker run -p 8000:8000 app
docker-compose up
```

### Basic Auth
```bash
curl -u user:pass /endpoint
```

### JWT
```bash
curl -X POST /auth/login -d '{"username":"u","password":"p"}'
curl -H "Authorization: Bearer <token>" /protected
```

### API Key
```bash
curl -H "X-API-Key: sk-xxx" /api/endpoint
```

---

## Project Structure

```
mux-tutorial/
├── main.go                 # Basic tutorial
├── LEARNING_GUIDE.md       # Mux basics
├── 03-env-config/          # Environment variables
│   ├── main.go
│   ├── .env.example
│   └── go.mod
├── 04-docker/              # Docker setup
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── .dockerignore
└── 05-authentication/      # All auth methods
    ├── main.go
    └── go.mod
```
