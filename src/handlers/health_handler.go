package handlers

import (
    "log"
    "time"

    "github.com/go-openapi/runtime/middleware"
    "github.com/munishchaudhary/book-store-app/src/generated/restapi/operations"
)

// HealthCheckHandler handles health check requests
func HealthCheckHandler(params operations.HealthCheckParams) middleware.Responder {
    start := time.Now()
    log.Printf("[HealthCheck] request received")
	healthStatus := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "book-store-api",
		"version":   "1.0.0",
	}

    log.Printf("[HealthCheck] success in %s", time.Since(start))
    return StatusOK(healthStatus)
}
