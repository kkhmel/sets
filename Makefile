.PHONY: help test bench lint all coverage

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Available targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Visit github.com/golangci/golangci-lint to install."; \
	fi

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

.PHONY: cover
cover:
	@echo "Generating HTML coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: bench
bench:
	@echo "Running benchmarks..."
	@cd benchmark && go test -bench=. -benchmem

.PHONY: all
all: lint test
