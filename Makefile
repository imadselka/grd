.PHONY: test build lint coverage clean

all: lint test build

build:
	@echo "Building..."
	@go build -v ./...

test:
	@echo "Running tests..."
	@go test -v ./...

test-race:
	@echo "Running tests with race detection..."
	@go test -race -v ./...

coverage:
	@echo "Running tests with coverage..."
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -html=coverage.txt -o coverage.html

lint:
	@echo "Running linters..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi

bench:
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

clean:
	@echo "Cleaning..."
	@rm -f coverage.txt coverage.html

help:
	@echo "Available targets:"
	@echo "  all        - Run lint, test, and build"
	@echo "  build      - Build the package"
	@echo "  test       - Run tests"
	@echo "  test-race  - Run tests with race detection"
	@echo "  coverage   - Generate test coverage report"
	@echo "  lint       - Run linters"
	@echo "  bench      - Run benchmarks"
	@echo "  clean      - Remove generated files"
	@echo "  help       - Show this help message"
