# Go commands
GOCMD=go
GOFMT=gofmt

# Directories definitions
OUTPUT_DIR=./out
BUILD_OUTPUT=$(OUTPUT_DIR)/bin/
COVERAGE_OUTPUT=$(OUTPUT_DIR)/coverage.txt
MIGRATIONS_DIR=./internal/db/migrations

# Migration definitions
DATABASE_URL=postgres://postgres:password@localhost:5432/wishlist?sslmode=disable

.PHONY: all run build clean format test migrate migrate-down migrate-goto create-migration


# Application

all: build

run:
	$(GOCMD) run ./cmd/wishlistbackend

build:
	$(GOCMD) build -o $(BUILD_OUTPUT) ./cmd/wishlistbackend

clean:
	rm -rf $(OUTPUT_DIR)

# Format

format:
	$(GOFMT) -s -w .

# Test

test:
	$(GOCMD) test -v -race ./...

# Migrations

migrate:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) up $(n)

migrate-down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down $(n)

migrate-goto:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) goto $(version)

create-migration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) --digits=4 -seq $(name)
