# tennis-court-api-main

This project is a Go-based API for managing tennis court facilities, users, and roles. 

## API Targets

- **Matches** The API should handle different kind of matches (Americano, Singles, Doubles, Series and more)

## Project Structure:

- **`cmd/api`**: Contains the main application entry point and dependency injection setup.
- **`internal/`**: Core application logic, divided into modules:
    - **`health`**: Health check endpoint.
- **`api`**: Likely contains API definitions (e.g., OpenAPI/Swagger specifications, though no specific files are present in the provided structure).
- **`Dockerfile`, `docker-compose.prod.yml`, `docker-compose.test.yml`**: Docker-related files for building, running, and testing the application in containers.
- **`go.mod`, `go.sum`**: Go module definition and dependency files.
- **`.env.example`**: Example environment variables.
- **`.gitignore`**: Git ignore file.

## Technologies Used:

- **Go**: Primary programming language.
- **Gin Web Framework**: (Inferred from `http_handler.go` and common Go API patterns, might need confirmation)

## Test

- **Repository and Services** Test should be created for every service and repository
- **Simple** Keep test simple and clean

## Setup & Development:

1.  Testing with `main_test.go` and potentially `docker-compose.test.yml`.