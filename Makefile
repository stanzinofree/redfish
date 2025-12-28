.PHONY: help build test lint fmt vet clean install run

help:
	@echo "Available targets:"
	@echo "  build      - Build the redfish binary"
	@echo "  test       - Run tests"
	@echo "  lint       - Run golangci-lint"
	@echo "  fmt        - Run gofmt"
	@echo "  vet        - Run go vet"
	@echo "  clean      - Remove build artifacts"
	@echo "  install    - Install redfish to GOPATH/bin"
	@echo "  run        - Run redfish with example query"
	@echo "  check      - Run all checks (fmt, vet, lint, test)"

build:
	@echo "Building redfish..."
	@go build -o redfish ./cmd/redfish

test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --timeout=5m

fmt:
	@echo "Running gofmt..."
	@gofmt -w -s .

vet:
	@echo "Running go vet..."
	@go vet ./...

clean:
	@echo "Cleaning..."
	@rm -f redfish coverage.out
	@go clean

install:
	@echo "Installing redfish..."
	@go install ./cmd/redfish

run: build
	@echo "Running redfish git..."
	@./redfish git

check: fmt vet lint test
	@echo "All checks passed!"

.DEFAULT_GOAL := help
