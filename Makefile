# Makefile for Book Store Application

# Variables
APP_NAME := book-store-app
DOCKER_IMAGE := book-store-app
DOCKER_TAG := latest
SWAGGER_FILE := swagger.yaml
GENERATED_DIR := src/generated

# Go build variables
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BUILD_FLAGS := -ldflags="-s -w"

# Docker variables
DOCKER_REGISTRY ?= 
FULL_IMAGE_NAME := $(DOCKER_REGISTRY)$(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: help clean build test generate docker-build docker-run all

# Default target
help:
	@echo "Available targets:"
	@echo "  help         - Show this help message"
	@echo "  clean        - Clean build artifacts"
	@echo "  generate     - Generate code from swagger.yaml"
	@echo "  build        - Build the application locally"
	@echo "  test         - Run tests"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run the application in Docker"
	@echo "  all          - Clean, generate, build, and test"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(APP_NAME)
	rm -rf $(GENERATED_DIR)
	@echo "Clean complete"

# Generate code from swagger.yaml
generate:
	@echo "Generating code from $(SWAGGER_FILE)..."
	@if [ ! -f $(SWAGGER_FILE) ]; then \
		echo "Error: $(SWAGGER_FILE) not found"; \
		exit 1; \
	fi
	mkdir -p $(GENERATED_DIR)
	swagger generate server \
		--target $(GENERATED_DIR) \
		--name BookStore \
		--spec $(SWAGGER_FILE) \
		--principal interface{}
	@echo "Code generation complete"

# Build the application locally
build: generate
	@echo "Building $(APP_NAME) for $(GOOS)/$(GOARCH)..."
	go build $(BUILD_FLAGS) -o $(APP_NAME) ./src/main.go
	@echo "Build complete: $(APP_NAME)"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./src/...
	@echo "Tests complete"

# Build Docker image
docker-build:
	@echo "Building Docker image: $(FULL_IMAGE_NAME)"
	docker build -t $(FULL_IMAGE_NAME) .
	@echo "Docker build complete"

# Run the application in Docker
docker-run:
	@echo "Running $(APP_NAME) in Docker..."
	docker run --rm -p 8080:8080 \
		-e DB_HOST=host.docker.internal \
		-e DB_PORT=3306 \
		-e DB_USER=root \
		-e DB_PASSWORD=password \
		-e DB_NAME=bookstore \
		$(FULL_IMAGE_NAME)


# Run all steps: clean, generate, build, test
all: clean generate build test
	@echo "All steps completed successfully"

# Development target - build and run locally
dev: build
	@echo "Starting $(APP_NAME) in development mode..."
	./$(APP_NAME)