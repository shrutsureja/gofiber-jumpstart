# Load environment variables from .env file
include .env

MIGRATIONS_DIR="./src/database/migrations"

# Default task
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make migrate-up          - Apply all up migrations"
	@echo "  make migrate-down        - Roll back the last migration"
	@echo "  make create name=<name>  - Create a new migration"
	@echo "  make status              - Check the current migration status"

# Apply all up migrations
.PHONY: migrate-up
migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URI) GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR) goose up

# Roll back the last migration
.PHONY: migrate-down
migrate-down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URI) GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR) goose down

# Create a new migration
.PHONY: create
create:
	@if [ -z "$(name)" ]; then \
		echo "Error: name=<migration_name> is required"; \
		exit 1; \
	fi
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URI) GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR) goose create $(name) sql

# Check migration status
.PHONY: status
status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URI) GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR) goose status

# Roll back all migrations
.PHONY: migrate-reset
migrate-reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URI) GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR) goose reset

# Migrate to a specific version
.PHONY: migrate-to
migrate-to:
	@if [ -z "$(v)" ]; then \
		echo "Error: v=<version_number> is required"; \
		exit 1; \
	fi
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URI) GOOSE_MIGRATION_DIR=$(MIGRATIONS_DIR) goose up-to $(v)

# Build the application
.PHONY: build
build:
	@go build -o bin/app src/main.go

# Build the application for multiple OS and architectures in parallel
.PHONY: build-all
build-all:
	@$(MAKE) -j 3 build-linux build-darwin build-windows

# Build for Linux
.PHONY: build-linux
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64 src/main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/app-linux-arm64 src/main.go

# Build for macOS
.PHONY: build-darwin
build-darwin:
	@GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 src/main.go
	@GOOS=darwin GOARCH=arm64 go build -o bin/app-darwin-arm64 src/main.go

# Build for Windows
.PHONY: build-windows
build-windows:
	@GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64.exe src/main.go
	@GOOS=windows GOARCH=arm64 go build -o bin/app-windows-arm64.exe src/main.go

# Clean the build directory
.PHONY: clean
clean:
	@rm -rf bin

# Rebuild the application
.PHONY: rebuild
rebuild: clean build
