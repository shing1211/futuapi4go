# Makefile for futuapi4go

.PHONY: help build test vet lint fmt clean cover install-tools ci build-examples

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOVET := $(GOCMD) vet
GOFMT := gofmt
GOMOD := $(GOCMD) mod
GOCOVER := $(GOCMD) test -coverprofile=coverage.out

# Default target
help:
	@echo "futuapi4go Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build        Build all packages"
	@echo "  make test         Run all tests"
	@echo "  make test-race    Run tests with race detector"
	@echo "  make vet          Run go vet"
	@echo "  make fmt          Format code"
	@echo "  make lint         Run linters (requires golint, staticcheck)"
	@echo "  make cover        Run tests with coverage report"
	@echo "  make ci           Run full CI pipeline (build, vet, fmt, test)"
	@echo "  make clean        Clean build artifacts"
	@echo "  make build-examples  Build all example programs"

# Build all packages
build:
	$(GOBUILD) ./...

# Build example programs
build-examples:
	cd cmd/examples && $(GOBUILD) ./...

# Run all tests
test:
	$(GOTEST) ./...

# Run tests with race detector
test-race:
	$(GOTEST) -race ./...

# Run go vet
vet:
	$(GOVET) ./...

# Format code
fmt:
	$(GOFMT) -w .
	@echo "Formatting complete. Check for any unformatted files with: gofmt -l ."

# Check formatting (exit non-zero if files need formatting)
fmt-check:
	@files=$$($(GOFMT) -l .); \
	if [ -n "$$files" ]; then \
		echo "Files need formatting:"; \
		echo "$$files"; \
		exit 1; \
	fi
	@echo "All files properly formatted."

# Install linting tools
install-tools:
	$(GOCMD) install golang.org/x/lint/golint@latest
	$(GOCMD) install honnef.co/go/tools/cmd/staticcheck@latest

# Run linters (install tools first with: make install-tools)
lint: vet
	@echo "Running golint..."
	golint -set_exit_status ./...
	@echo "Running staticcheck..."
	staticcheck ./...

# Run tests with coverage report
cover:
	$(GOCOVER) ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo "Coverage summary:"
	$(GOCMD) tool cover -func=coverage.out | tail -1

# Full CI pipeline (mirrors .github/workflows/ci.yml)
ci: fmt-check vet build test

# Clean build artifacts
clean:
	$(GOCMD) clean
	rm -f coverage.out coverage.html
	find . -name "*.test" -delete 2>/dev/null || true
