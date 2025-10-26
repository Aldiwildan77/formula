# Makefile for formula

# Variables
PACKAGE := formula
FUNCTIONS_PKG := $(PACKAGE)/functions
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Default target
.PHONY: all
all: test

# Build the project
.PHONY: build
build:
	@echo "Building formula..."
	@go build ./...

# Run all tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v

# Run tests for functions package only
.PHONY: test-functions
test-functions:
	@echo "Running tests for functions package..."
	@go test $(FUNCTIONS_PKG)/... -v -coverprofile=${COVERAGE_FILE}

# Generate HTML coverage report (excluding test files)
.PHONY: coverage-html
test-functions-coverage: test-functions
	@echo "Generating HTML coverage report (functions) ..."
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated at $(COVERAGE_HTML)"

# Clean up generated files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f $(COVERAGE_FILE) $(CUSTOM_COVERAGE) $(COVERAGE_HTML)
	@go clean

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all             : Default target, runs tests"
	@echo "  build           : Build the project"
	@echo "  test            : Run all tests"
	@echo "  test-functions  : Run tests for functions package only"
	@echo "  test-coverage   : Run tests with coverage"
	@echo "  coverage-html   : Generate HTML coverage report (excluding test files)"
	@echo "  clean           : Clean up generated files"
	@echo "  help            : Show this help message"