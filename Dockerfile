# Multi-stage Dockerfile for Book Store Application

# Stage 1: Build stage with Go and Swagger
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy swagger specification
COPY swagger.yaml ./

# Install swagger CLI
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# Create src directory and generate code from swagger
RUN mkdir -p src/generated && \
    swagger generate server \
    --target src/generated \
    --name BookStore \
    --spec swagger.yaml \
    --principal interface{}

# Copy source code
COPY src/ ./src/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o book-store-app \
    ./src/main.go

# Stage 2: Runtime stage
FROM alpine:latest

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/book-store-app .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./book-store-app"]