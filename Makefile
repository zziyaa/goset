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

race:
	@echo "Running race detector..."
	@go test -race ./...

check: fmt vet test race

clean:
	@echo "Cleaning..."
	@go clean

.PHONY: build test fmt vet race check clean