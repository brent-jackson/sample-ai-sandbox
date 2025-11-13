#!/bin/bash

# Test client for the Sample MCP Server
# This script demonstrates how to interact with the server programmatically

echo "=== Testing Sample MCP Server ==="
echo

# Ensure the server is built
echo "Building server..."
go build -o bin/sample-mcp-server main.go
if [ $? -ne 0 ]; then
    echo "Failed to build server"
    exit 1
fi

echo "✅ Server built successfully"
echo

# Test 1: Basic startup test
echo "Test 1: Testing server startup..."
timeout 2s ./bin/sample-mcp-server > /dev/null 2>&1 &
SERVER_PID=$!
sleep 1

if kill -0 $SERVER_PID 2>/dev/null; then
    echo "✅ Server starts successfully"
    kill $SERVER_PID 2>/dev/null
else
    echo "❌ Server failed to start"
fi
echo

# Test 2: JSON-RPC initialization
echo "Test 2: Testing JSON-RPC initialization..."
INIT_REQUEST='{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2024-11-05",
    "capabilities": {},
    "clientInfo": {
      "name": "test-client",
      "version": "1.0.0"
    }
  }
}'

echo "$INIT_REQUEST" | timeout 3s ./bin/sample-mcp-server 2>/dev/null | head -1 | grep -q "jsonrpc"
if [ $? -eq 0 ]; then
    echo "✅ Server responds to JSON-RPC initialization"
else
    echo "❌ Server does not respond properly to JSON-RPC"
fi
echo

# Test 3: Run unit tests
echo "Test 3: Running unit tests..."
go test -v
if [ $? -eq 0 ]; then
    echo "✅ All unit tests pass"
else
    echo "❌ Some unit tests failed"
fi
echo

echo "=== Test Summary ==="
echo "Your MCP server is ready to use!"
echo
echo "Next steps:"
echo "1. Configure Claude Desktop (see USAGE.md)"
echo "2. Or install MCP Inspector: npm install -g @modelcontextprotocol/inspector"
echo "3. Or build a custom client using the Go SDK"
echo
echo "Available tools: calculate, text_transform, current_time"
echo "Available resources: sample://config, sample://status, sample://docs/readme"