# futuapi4go Makefile

.PHONY: build test vet lint clean install bench release docs-check help

# Build all packages
build:
	go build ./...

# Run all tests (with race detection)
test:
	go test -race ./...

# Run tests with coverage
test-cover:
	go test -cover ./...

# Run go vet
vet:
	go vet ./...

# Run fmt check
fmt:
	gofmt -l .
	gofmt -d .

# Format code
fmt-fix:
	gofmt -w .

# Run static analysis
staticcheck:
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...

# Run benchmarks
bench:
	go test -bench=. -benchmem -count=3 ./...

# Clean build artifacts
clean:
	go clean -cache
	rm -f *.out

# Install dependencies
install:
	go mod download
	go mod tidy

# Quick check (fmt + vet + build)
check: fmt-fix vet build

# Run specific package tests
test-pkg:
	go test -race ./pkg/constant/...
	go test -race ./client/...
	go test -race ./internal/client/...

# Run integration tests (requires a running OpenD)
test-integration:
	FUTU_INTEGRATION_TESTS=1 go test -race ./test/integration/...

# Create a GitHub release for the tag at HEAD (manual fallback; normally
# handled by .github/workflows/release.yml when a v* tag is pushed).
release:
	@tag=$$(git describe --tags --exact-match 2>/dev/null) || { \
		echo "HEAD is not tagged. Create and push a tag first:"; \
		echo "  git tag -a vX.Y.Z -m 'vX.Y.Z' && git push origin vX.Y.Z"; \
		exit 1; \
	}; \
	gh release create "$$tag" --title "$$tag" --generate-notes

# Check README translations are consistent with the English canonical
docs-check:
	python3 scripts/check_i18n.py

# Show help
help:
	@echo "futuapi4go Makefile targets:"
	@echo "  make build         - Build all packages"
	@echo "  make test         - Run all tests with race detection"
	@echo "  make test-cover  - Run tests with coverage"
	@echo "  make vet         - Run go vet"
	@echo "  make fmt         - Check code formatting"
	@echo "  make fmt-fix     - Fix code formatting"
	@echo "  make staticcheck - Run static analysis"
	@echo "  make bench       - Run benchmarks"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make install     - Install dependencies"
	@echo "  make check       - Quick check (fmt + vet + build)"
	@echo "  make test-pkg    - Run specific package tests"
	@echo "  make release     - Create GitHub release (requires git tag)"
	@echo "  make docs-check  - Check README translations"
	@echo "  make help        - Show this help"