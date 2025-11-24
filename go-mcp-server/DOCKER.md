# Docker Deployment Guide for Sample MCP Server

This guide covers how to build and run the Sample MCP Server using Docker.

## Quick Start

### Build the Docker Image
```bash
# Using Make
make docker-build

# Or directly with Docker
docker build -t sample-mcp-server:latest .
```

### Run the Container
```bash
# Using Make
make docker-run

# Or directly with Docker
docker run --rm -i sample-mcp-server:latest
```

### Using Docker Compose
```bash
# Start the service
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the service
docker-compose down
```

## Dockerfile Overview

The Dockerfile uses a **multi-stage build** for optimal image size and security:

### Stage 1: Builder
- Based on `golang:1.25-alpine`
- Installs build dependencies
- Downloads Go modules
- Compiles the application with optimizations
- Runs tests to ensure build quality

### Stage 2: Runtime
- Based on `alpine:latest` (minimal ~5MB base)
- Runs as non-root user for security
- Only contains the compiled binary and CA certificates
- Final image size: ~15-20MB

## Configuration

### Environment Variables
```bash
MCP_SERVER_NAME=sample-mcp-server
MCP_SERVER_VERSION=1.0.0
```

### Custom Build Arguments
```bash
# Build with custom image name/tag
make docker-build IMAGE_NAME=my-mcp-server IMAGE_TAG=v1.0.0

# Or with Docker
docker build -t my-mcp-server:v1.0.0 .
```

## Running in Production

### With Docker Run
```bash
docker run -d \
  --name mcp-server \
  --restart unless-stopped \
  -e MCP_SERVER_NAME=production-mcp \
  sample-mcp-server:latest
```

### With Docker Compose
```bash
# Production deployment
docker-compose up -d

# Scale if needed (future enhancement)
docker-compose up -d --scale sample-mcp-server=3
```

## Development Workflow

### Build and Test
```bash
# Build image
make docker-build

# Run tests inside container
docker run --rm sample-mcp-server:latest go test -v ./...

# Interactive shell for debugging
docker run --rm -it --entrypoint /bin/sh sample-mcp-server:latest
```

### Hot Reload Development
For development, mount the source code:
```bash
docker run --rm -it \
  -v $(pwd):/app \
  -w /app \
  golang:1.25-alpine \
  go run main.go
```

## Image Management

### View Image Details
```bash
# Show image size
make docker-size

# Inspect image
docker inspect sample-mcp-server:latest

# View layers
docker history sample-mcp-server:latest
```

### Clean Up
```bash
# Remove image
make docker-clean

# Remove all unused images
docker image prune -a

# Complete cleanup (compose)
docker-compose down --rmi all --volumes
```

## Registry Deployment

### Push to Docker Hub
```bash
# Login
docker login

# Tag for Docker Hub
docker tag sample-mcp-server:latest username/sample-mcp-server:latest

# Push
docker push username/sample-mcp-server:latest
```

### Push to Private Registry
```bash
# Tag for private registry
make docker-tag REGISTRY=registry.example.com

# Push
make docker-push REGISTRY=registry.example.com
```

## Kubernetes Deployment (Optional)

Basic Kubernetes deployment example:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sample-mcp-server
spec:
  replicas: 2
  selector:
    matchLabels:
      app: sample-mcp-server
  template:
    metadata:
      labels:
        app: sample-mcp-server
    spec:
      containers:
      - name: mcp-server
        image: sample-mcp-server:latest
        env:
        - name: MCP_SERVER_NAME
          value: "sample-mcp-server"
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
```

## Troubleshooting

### Container Won't Start
```bash
# Check logs
docker logs sample-mcp-server

# Run in foreground
docker run --rm -i sample-mcp-server:latest
```

### Build Failures
```bash
# Clean build cache
docker builder prune

# Build without cache
docker build --no-cache -t sample-mcp-server:latest .
```

### Debugging Inside Container
```bash
# Get a shell
docker run --rm -it --entrypoint /bin/sh sample-mcp-server:latest

# Check binary
docker run --rm sample-mcp-server:latest /app/sample-mcp-server --version
```

## Security Best Practices

✅ **Implemented:**
- Multi-stage build (smaller attack surface)
- Non-root user execution
- Minimal base image (Alpine)
- No unnecessary packages
- Static binary compilation

✅ **Recommended:**
- Scan images: `docker scan sample-mcp-server:latest`
- Use specific version tags, not `latest` in production
- Enable Docker Content Trust: `export DOCKER_CONTENT_TRUST=1`
- Regularly update base images

## Performance Optimization

### Image Size Optimization
- Current size: ~15-20MB
- Builder stage excluded from final image
- Static compilation with stripped symbols

### Runtime Optimization
```bash
# Set memory limits
docker run --memory=128m --memory-swap=128m sample-mcp-server:latest

# Set CPU limits
docker run --cpus=0.5 sample-mcp-server:latest
```

## Integration with MCP Clients

### Using with Docker-based MCP Clients
```bash
# Run server accessible to host
docker run --rm -i --name mcp-server sample-mcp-server:latest

# Connect from another container
docker run --link mcp-server:mcp-server client-image
```

## Monitoring

### Basic Monitoring
```bash
# Watch logs
docker logs -f sample-mcp-server

# Monitor resources
docker stats sample-mcp-server
```

### Advanced Monitoring (with Prometheus)
Add metrics endpoint and Prometheus configuration as needed for your use case.

---

For more information, see:
- [Main README](README.md)
- [Usage Guide](USAGE.md)
- [Docker Documentation](https://docs.docker.com/)
