# =============================================================
# Dockerfile for Go + Mux Application
# =============================================================

# -------------------------------------------------------------
# Stage 1: Build the application
# -------------------------------------------------------------
FROM golang:1.24-alpine AS builder

# Install ca-certificates for HTTPS and git for go get
RUN apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# -ldflags "-s -w" strips debug symbols for smaller binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o main .

# -------------------------------------------------------------
# Stage 2: Runtime image (multi-stage build for minimal size)
# -------------------------------------------------------------
FROM alpine:3.19

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -D appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8000/health || exit 1

# Run the application
CMD ["./main"]
