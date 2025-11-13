#!/bin/bash

# Setup script for Python FastMCP Server

echo "🐍 Setting up Python FastMCP Server..."

# Check Python version
echo "Checking Python version..."
python3.11 --version || {
    echo "❌ Error: Python 3.11 not found. FastMCP requires Python 3.10+"
    echo "Please install Python 3.11 or later"
    exit 1
}

# Create virtual environment
echo "Creating virtual environment..."
python3.11 -m venv venv

# Activate and install dependencies
echo "Installing FastMCP..."
source venv/bin/activate
pip install --upgrade pip
pip install --index-url https://pypi.org/simple/ fastmcp

echo "✅ Setup complete!"
echo ""
echo "🚀 Usage:"
echo "1. Configure Claude Desktop:"
echo '   Edit ~/Library/Application\ Support/Claude/claude_desktop_config.json'
echo ""
echo "2. Or use MCP Inspector:"
echo "   npm install -g @modelcontextprotocol/inspector"
echo "   source venv/bin/activate && mcp-inspector python server.py"
echo ""
echo "3. Available tool: add(a, b) - adds two numbers"
echo ""
echo "⚠️  Note: The server will appear to 'hang' when run directly."
echo "   This is normal - it's waiting for MCP protocol messages!"