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