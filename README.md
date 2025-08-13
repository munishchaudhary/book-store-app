# Book Store APP

A RESTful API for managing CRUD operations on books with Go, using Swagger for API documentation and code generation.

## Features

- **CRUD Operations**: Create, Read, Update, Delete books
- **Database Integration**: MySQL database with proper connection handling
- **API Documentation**: Auto-generated from Swagger specification
- **Docker Support**: Multi-stage Docker build with optimized runtime image
- **Makefile**: Automated build, test, and deployment tasks
- **Health Checks**: Built-in health monitoring endpoint

## Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose
- MySQL (for local development)
- Swagger CLI (optional, for local development)

## Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/munishchaudhary/book-store-app.git
   cd book-store-app
   ```

2. **Start the application with Docker Compose**
   ```bash
   docker-compose up --build
   ```

   This will:
   - Build the application using multi-stage Docker build
   - Start MySQL database with sample data
   - Start the Book Store API on port 8080

3. **Access the API**
   - API Base URL: `http://localhost:8080`
   - Swagger UI: `http://localhost:8080/docs`
   - Health Check: `http://localhost:8080/health`

### Local Development

1. **Install dependencies**
   ```bash
   go mod tidy
   ```

2. **Install Swagger CLI (optional)**
   ```bash
   make install-swagger
   ```

3. **Set up environment variables**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=3306
   export DB_USER=root
   export DB_PASSWORD=password
   export DB_NAME=bookstore
   export SVC_HOSTNAME=localhost
   export SVC_PORT=8080
   export SVC_SCHEMES=http
   ```

4. **Build and run**
   ```bash
   make build
   ./book-store
   ```

## Makefile Commands

The project includes a comprehensive Makefile with the following targets:

```bash
make help          # Show available commands
make clean         # Clean build artifacts
make generate      # Generate code from swagger.yaml
make build         # Build the application locally
make test          # Run tests
make docker-build  # Build Docker image
make docker-run    # Run application in Docker
make all           # Clean, generate, build, and test
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/books` | Get all books |
| GET | `/books/{id}` | Get a specific book |
| POST | `/books` | Create a new book |
| PUT | `/books/{id}` | Update a book |
| DELETE | `/books/{id}` | Delete a book |

## Book Model

```json
{
  "id": 1,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "year": 1925
}
```

## Health Check Response

```json
{
  "status": "healthy",
  "timestamp": "2024-08-14T01:15:26Z",
  "service": "book-store-api",
  "version": "1.0.0"
}
```

## Validation Rules

- **Title**: Required, max 255 characters
- **Author**: Required, max 255 characters
- **Year**: Required, between 1800 and current year
- **ID**: Must be positive integer (for updates)

## Database Schema

The application uses a MySQL database with the following schema:

```sql
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Docker Compose

The `docker-compose.yml` includes:

- **MySQL Database**: Version 8.0 with persistent storage
- **Book Store API**: Built from source with environment configuration
- **Health Checks**: For both services
- **Networking**: Isolated network for services

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_USER` | Database user | `root` |
| `DB_PASSWORD` | Database password | `password` |
| `DB_NAME` | Database name | `bookstore` |
| `SVC_HOSTNAME` | Service hostname | `localhost` |
| `SVC_PORT` | Service port | `8080` |
| `SVC_SCHEMES` | Service schemes | `http` |

## Development Workflow

1. **Modify API Specification**: Edit `swagger.yaml`
2. **Regenerate Code**: `make generate`
3. **Update Handlers**: Modify handlers in `src/handlers/`
4. **Add Validation**: Update validation rules in `src/validation/`
5. **Test**: `make test`
6. **Build**: `make build`
7. **Deploy**: `docker-compose up --build`


