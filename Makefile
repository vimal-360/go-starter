.PHONY: build run test clean deps fmt vet lint help

# Variables
BINARY_NAME=server
BINARY_PATH=./bin/$(BINARY_NAME)
CMD_PATH=./cmd/server
GO_FILES=$(shell find . -name "*.go" -type f -not -path "./vendor/*")

# Default target
all: fmt vet build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Download and install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	go build -o $(BINARY_PATH) $(CMD_PATH)

run: build ## Build and run the application
	@echo "Starting $(BINARY_NAME)..."
	$(BINARY_PATH)

dev: ## Run the application in development mode with auto-reload
	go run $(CMD_PATH)

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

fmt: ## Format Go source code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run golangci-lint (requires golangci-lint to be installed)
	golangci-lint run

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -rf data/

install-tools: ## Install development tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

docker-build: ## Build Docker image
	docker build -t $(BINARY_NAME) .

docker-run: docker-build ## Build and run Docker container
	docker run -p 8080:8080 $(BINARY_NAME)

migrate: ## Run database migrations
	go run $(CMD_PATH) migrate

seed: ## Seed the database with sample data
	go run $(CMD_PATH) seed