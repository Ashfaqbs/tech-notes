// Environment Variables & Configuration Example
// Run with:
//   PORT=9000 ENV=development go run main.go
//   OR
//   go run main.go (uses defaults from .env file via godotenv)

package main

import (
	"context"       // Context for graceful shutdown - used for context.WithTimeout() to set shutdown deadline
	"encoding/json" // JSON encoding/decoding - used in JSONResponse() for json.NewEncoder(w).Encode(data)
	"fmt"          // Formatted I/O - used for fmt.Println(), fmt.Sprintf() to format strings
	"log"          // Logging - used for log.Printf(), log.Fatalf() to log server events
	"net/http"     // HTTP server - used for http.ResponseWriter, http.Request, http.Server, http.ListenAndServe
	"os"           // Operating system functions - used for os.Getenv() to read environment variables
	"os/signal"    // Signal handling - used for signal.Notify() to catch Ctrl+C (SIGINT, SIGTERM)
	"strconv"      // String conversion - used for strconv.Atoi() to parse timeout values from env
	"strings"      // String manipulation - used for strings.SplitN(), strings.Contains() to mask passwords
	"syscall"      // System calls - used for syscall.SIGTERM to handle graceful shutdown signals
	"time"         // Time operations - used for time.Second, time.Now(), timeouts in server config

	"github.com/gorilla/mux"    // Gorilla Mux router - used for mux.NewRouter(), r.HandleFunc()
	"github.com/joho/godotenv"  // godotenv - loads .env files into environment variables
)

// Config holds all application configuration
type Config struct {
	Port         string
	Env          string // development, staging, production
	DatabaseURL  string
	JWTSecret    string
	AllowedOrigins []string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Global config instance
var config Config

func main() {
	// Load .env file (optional, won't fail if not present)
	_ = godotenv.Load()

	// Initialize config from environment
	config = Config{
		Port:         getEnv("PORT", "8000"),
		Env:          getEnv("ENV", "development"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://localhost:5432/mydb"),
		JWTSecret:    getEnv("JWT_SECRET", "change-me-in-production"),
		AllowedOrigins: []string{
			getEnv("ALLOWED_ORIGIN", "http://localhost:3000"),
		},
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Create router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/config", ConfigHandler)
	r.HandleFunc("/health", HealthHandler)

	// Create server with timeouts
	srv := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      r,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	// Start server in goroutine
	go func() {
		log.Printf("[%s] Server starting on port %s", config.Env, config.Port)
		log.Printf("Database URL: %s", config.DatabaseURL)
		log.Printf("Allowed Origins: %v", config.AllowedOrigins)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}

// HomeHandler - uses config
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message":   "Welcome!",
		"env":       config.Env,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	JSONResponse(w, response)
}

// ConfigHandler - shows current config (don't expose secrets in production!)
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"port":           config.Port,
		"env":            config.Env,
		"database_url":   maskPassword(config.DatabaseURL),
		"allowed_origins": config.AllowedOrigins,
		"read_timeout":   config.ReadTimeout.String(),
		"write_timeout":  config.WriteTimeout.String(),
		// NOTE: Never expose JWT_SECRET in production!
		// "jwt_secret": config.JWTSecret,
	}
	JSONResponse(w, response)
}

// HealthHandler - health check with env info
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":      "healthy",
		"environment": config.Env,
		"uptime":      time.Now().Format(time.RFC3339),
	}
	JSONResponse(w, response)
}

// JSONResponse sends JSON response
func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// getEnv reads environment variable with fallback default
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvInt reads environment variable as int with fallback
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

// maskPassword masks password in connection strings for logging
func maskPassword(url string) string {
	// Simple masking - show URL but hide password
	if strings.Contains(url, "@") {
		parts := strings.SplitN(url, "@", 2)
		if strings.Contains(parts[0], ":") {
			creds := strings.SplitN(parts[0], ":", 2)
			return creds[0] + ":***@" + parts[1]
		}
	}
	return url
}
