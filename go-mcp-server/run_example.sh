#!/bin/bash

# Example script showing how to test the MCP server manually
# Note: This requires an MCP client to properly test

echo "Building the MCP server..."
make build 

echo "Sample MCP Server built successfully!"
echo ""
echo "To use this server with an MCP client:"
echo "1. Use the configuration in mcp-config.json"
echo "2. The server provides the following tools:"
echo "   - calculate: Basic math operations"
echo "   - text_transform: Text manipulation"
echo "   - current_time: Time formatting"
echo ""
echo "3. The server provides these resources:"
echo "   - sample://config: Server configuration"
echo "   - sample://status: Server status"
echo "   - sample://docs/readme: Documentation"
echo ""
echo "To run the server directly (for testing):"
echo "./bin/sample-mcp-server"
echo ""
echo "Note: The server uses stdio transport and expects MCP protocol messages via stdin/stdout"