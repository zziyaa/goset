build:
	@echo "Building..."
	@go build ./...

test:
	@echo "Running tests..."
	@go test ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

vet:
	@echo "Vetting code..."
	@go vet ./...

check: fmt vet test

clean:
	@echo "Cleaning..."
	@go clean

.PHONY: build test fmt vet check clean