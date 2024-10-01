# Go commands
GOCMD=go
GOFMT=gofmt

# Directories definitions
OUTPUT_DIR=./out
BUILD_OUTPUT=$(OUTPUT_DIR)/bin/
COVERAGE_OUTPUT=$(OUTPUT_DIR)/coverage.txt

.PHONY: all run build clean format test


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
