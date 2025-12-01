.PHONY: help test build clean fmt lint run-example

help:
	@echo "Available targets:"
	@echo "  test         - Run all tests"
	@echo "  build        - Build the SDK"
	@echo "  clean        - Clean build artifacts"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  run-example  - Run the simple example"

test:
	go test -v ./...

build:
	go build ./...

clean:
	go clean ./...
	rm -rf bin/

fmt:
	go fmt ./...

lint:
	@which golangci-lint > /dev/null || (echo "Please install golangci-lint: https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

run-example:
	@if [ -z "$$BROWSER_USE_API_KEY" ]; then \
		echo "Error: BROWSER_USE_API_KEY environment variable is not set"; \
		exit 1; \
	fi
	cd examples/simple && go run main.go
