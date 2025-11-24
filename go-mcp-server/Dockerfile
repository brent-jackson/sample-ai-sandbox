# Multi-stage Dockerfile for sample-mcp-server

# Build stage
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /build

# Copy go mod files, and download dependencies
# Doing this before the go build allows for downloaded dependencies to be cached in its own layer
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o sample-mcp-server main.go

RUN go test -v ./...

# Final stage - minimal runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -S mcp && adduser -S mcp -G mcp

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/sample-mcp-server .

# Change ownership
RUN chown -R mcp:mcp /app

# Switch to non-root user
USER mcp

# Expose port (if needed for HTTP endpoints)
# EXPOSE 8080

# Set environment variables
ENV MCP_SERVER_NAME="sample-mcp-server" \
    MCP_SERVER_VERSION="1.0.0"

# Health check (optional - adjust if your server has a health endpoint)
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#   CMD ["/app/sample-mcp-server", "--health"]

# Run the MCP server
ENTRYPOINT ["/app/sample-mcp-server"]
CMD []
