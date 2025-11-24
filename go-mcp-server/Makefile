.PHONY: run build test clean fmt vet

# Run the MCP server
run:
	go run main.go

# Build the MCP server binary
build:
	mkdir -p bin
	go build -o bin/sample-mcp-server main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f bin/sample-mcp-server

# Format Go code
fmt:
	go fmt ./...

# Run Go vet
vet:
	go vet ./...

# Run all checks
check: fmt vet test

# Install dependencies
deps:
	go mod download
	go mod tidy

# Build for multiple platforms
build-all:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/sample-mcp-server-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/sample-mcp-server-darwin-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/sample-mcp-server-windows-amd64.exe main.go

# Docker targets
IMAGE_NAME ?= sample-ai-sandbox
IMAGE_TAG ?= golang-latest
REGISTRY ?= jacksonwbrent

# Build Docker image
docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Run Docker container
docker-run:
	docker run --rm -i $(IMAGE_NAME):$(IMAGE_TAG)

# Run Docker container with environment variables
docker-run-env:
	docker run --rm -i \
		-e MCP_SERVER_NAME=sample-mcp-server \
		-e MCP_SERVER_VERSION=1.0.0 \
		$(IMAGE_NAME):$(IMAGE_TAG)

# Build and run Docker container
docker-up: docker-build docker-run

# Push Docker image to registry (requires login: docker login)
docker-push:
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# Tag Docker image for registry
docker-tag:
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# Clean Docker artifacts
docker-clean:
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG) || true

# Show Docker image size
docker-size:
	docker images $(IMAGE_NAME):$(IMAGE_TAG) --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}"