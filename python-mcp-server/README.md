# Python FastMCP Server

A simple MCP server built with FastMCP that demonstrates basic tool functionality.

## Setup and Installation

### Prerequisites
- Python 3.10 or later (FastMCP requirement)

### Installation

1. **Create and activate virtual environment**:
   ```bash
   cd /Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/python-mcp-server
   python3.11 -m venv venv
   source venv/bin/activate
   ```

2. **Install dependencies**:
   ```bash
   pip install --index-url https://pypi.org/simple/ fastmcp
   ```

## Running the Server

⚠️ **Important**: The MCP server runs as a stdio process and will appear to "hang" when run directly because it's waiting for JSON-RPC messages on stdin. This is normal behavior!

### Method 1: With Claude Desktop (Recommended)

1. **Configure Claude Desktop**:
   ```bash
   mkdir -p ~/Library/Application\ Support/Claude/
   cat > ~/Library/Application\ Support/Claude/claude_desktop_config.json << 'EOF'
   {
     "mcpServers": {
       "python-demo-server": {
         "command": "/Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/python-mcp-server/venv/bin/python",
         "args": ["/Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/python-mcp-server/server.py"],
         "cwd": "/Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/python-mcp-server"
       }
     }
   }
   EOF
   ```

2. **Restart Claude Desktop** and ask: "Add 5 and 3 using the add tool"

### Method 2: With MCP Inspector

```bash
# Install MCP Inspector (if not already installed)
npm install -g @modelcontextprotocol/inspector

# Activate Python environment and run inspector
cd /Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/python-mcp-server
source venv/bin/activate
mcp-inspector python server.py
```

### Method 3: Test with curl (if FastMCP supports HTTP)

Some versions of FastMCP may support HTTP endpoints. Check if the server exposes HTTP:

```bash
cd /Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/python-mcp-server
source venv/bin/activate
python server.py --help  # Check available options
```

## Available Tools

### `add(a: int, b: int) -> int`
Adds two integers together.

**Example usage in Claude**:
- "Add 5 and 3"
- "What's 10 plus 7?"
- "Use the add tool to calculate 15 + 25"

## Extending the Server

You can easily add more tools to the server by adding functions with the `@mcp.tool` decorator:

```python
@mcp.tool
def multiply(a: int, b: int) -> int:
    """Multiply two numbers"""
    return a * b

@mcp.tool
def greet(name: str) -> str:
    """Greet someone by name"""
    return f"Hello, {name}!"
```

## Troubleshooting

### Server appears to hang
This is normal! MCP servers run as stdio processes waiting for JSON-RPC messages. Use one of the client methods above to interact with it.

### Python version error
FastMCP requires Python 3.10+. Make sure you're using the correct Python version:
```bash
python3.11 --version  # Should show 3.11.x
```

### Import errors
Make sure you've activated the virtual environment:
```bash
source venv/bin/activate
pip list | grep fastmcp  # Should show fastmcp installation
```

## Development

To modify the server:
1. Edit `server.py`
2. Save the file
3. Restart Claude Desktop or the MCP Inspector to pick up changes

The FastMCP framework automatically handles:
- JSON schema generation from Python type hints
- Input validation
- MCP protocol implementation
- Error handling