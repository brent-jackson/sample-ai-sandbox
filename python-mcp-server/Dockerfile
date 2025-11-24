# Multi-stage Dockerfile for python-mcp-server

# Build stage
FROM python:3.11-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev linux-headers

# Set working directory
WORKDIR /build

# Copy requirements and install dependencies
COPY requirements.txt ./
RUN pip install --no-cache-dir --user -r requirements.txt

# Final stage - minimal runtime image
FROM python:3.11-alpine

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -S mcp && adduser -S mcp -G mcp

WORKDIR /app

# Copy Python dependencies from builder
COPY --from=builder /root/.local /home/mcp/.local

# Copy application code
COPY server.py enhanced_server.py ./

# Change ownership
RUN chown -R mcp:mcp /app

# Switch to non-root user
USER mcp

# Make sure scripts in .local are in PATH
ENV PATH=/home/mcp/.local/bin:$PATH

# Set environment variables
ENV MCP_SERVER_NAME="python-mcp-server" \
    MCP_SERVER_VERSION="1.0.0" \
    PYTHONUNBUFFERED=1

# Health check (optional - adjust if your server has a health endpoint)
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#   CMD ["python", "-c", "import sys; sys.exit(0)"]

# Run the MCP server (default to enhanced_server.py)
ENTRYPOINT ["python"]
CMD ["enhanced_server.py"]
