# Validation Service Makefile

# Variables
BINARY_NAME=validation-service
DOCKER_IMAGE=validation-service
VERSION=1.0.0-SNAPSHOT

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build targets
.PHONY: all build clean test coverage run docker-build docker-run help

all: test build

## Build the application
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/main.go

## Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)

## Run tests
test:
	$(GOTEST) -v ./...

## Run tests with coverage
coverage:
	$(GOTEST) -cover ./...
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

## Run the application
run:
	$(GOCMD) run ./cmd/main.go

## Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

## Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

## Run Docker container
docker-run:
	docker run -p 8081:8081 --name $(BINARY_NAME) $(DOCKER_IMAGE):latest

## Stop and remove Docker container
docker-stop:
	docker stop $(BINARY_NAME) || true
	docker rm $(BINARY_NAME) || true

## Format code
fmt:
	$(GOCMD) fmt ./...

## Lint code
lint:
	golangci-lint run

## Security scan
security:
	gosec ./...

## Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/main.go
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/main.go
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/main.go

## Development with hot reload (requires air)
dev:
	air

## Show help
help:
	@echo "Available commands:"
	@echo "  build        Build the application"
	@echo "  clean        Clean build artifacts"
	@echo "  test         Run tests"
	@echo "  coverage     Run tests with coverage report"
	@echo "  run          Run the application"
	@echo "  deps         Install dependencies"
	@echo "  docker-build Build Docker image"
	@echo "  docker-run   Run Docker container"
	@echo "  docker-stop  Stop Docker container"
	@echo "  fmt          Format code"
	@echo "  lint         Lint code"
	@echo "  security     Run security scan"
	@echo "  build-all    Build for multiple platforms"
	@echo "  dev          Run with hot reload"
	@echo "  help         Show this help message"