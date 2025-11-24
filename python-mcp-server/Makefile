.PHONY: run build test clean install

# Run the MCP server (enhanced version)
run:
	python enhanced_server.py

# Run the basic MCP server
run-basic:
	python server.py

# Install dependencies
install:
	pip install -r requirements.txt

# Install dependencies in development mode
install-dev:
	pip install -r requirements.txt
	pip install pytest pytest-asyncio

# Run tests (if you have them)
test:
	pytest -v

# Clean Python cache
clean:
	find . -type d -name "__pycache__" -exec rm -r {} +
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete

# Format Python code
fmt:
	black *.py

# Lint Python code
lint:
	pylint *.py

# Run all checks
check: lint test

# Docker targets
IMAGE_NAME ?= sample-ai-sandbox
IMAGE_TAG ?= python-latest
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
		-e MCP_SERVER_NAME=python-mcp-server \
		-e MCP_SERVER_VERSION=1.0.0 \
		$(IMAGE_NAME):$(IMAGE_TAG)

# Build and run Docker container
docker-up: docker-build docker-run

# Start with docker-compose
compose-up:
	docker-compose up -d

# Stop docker-compose services
compose-down:
	docker-compose down

# View docker-compose logs
compose-logs:
	docker-compose logs -f

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
