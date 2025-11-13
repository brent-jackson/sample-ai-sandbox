# How to Interact with the Sample MCP Server

There are several ways to interact with your MCP server. Here are the main methods:

## 1. Using Claude Desktop (Recommended)

### Setup
1. **Install Claude Desktop** from [https://claude.ai/download](https://claude.ai/download)

2. **Configure Claude Desktop** to use your MCP server:
   ```bash
   # On macOS
   mkdir -p ~/Library/Application\ Support/Claude/
   cat > ~/Library/Application\ Support/Claude/claude_desktop_config.json << 'EOF'
   {
     "mcpServers": {
       "sample-mcp-server": {
         "command": "/usr/local/go/bin/go",
         "args": ["run", "main.go"],
         "cwd": "/Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/sample-mcp-server"
       }
     }
   }
   EOF
   ```

   Or use the pre-built binary:
   ```bash
   cat > ~/Library/Application\ Support/Claude/claude_desktop_config.json << 'EOF'
   {
     "mcpServers": {
       "sample-mcp-server": {
         "command": "/Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/sample-mcp-server/bin/sample-mcp-server"
       }
     }
   }
   EOF
   ```

3. **Restart Claude Desktop**

### Usage
Once configured, you can ask Claude to:
- "Calculate 2 + 3 using the calculator tool"
- "Transform 'hello world' to uppercase" 
- "Get the current time in human readable format"
- "Show me the server configuration"
- "What's the server status?"

## 2. Using MCP Inspector (Development/Testing)

The MCP Inspector is a web-based tool for testing MCP servers:

```bash
# Install the MCP Inspector
npm install -g @modelcontextprotocol/inspector

# Run it with your server
mcp-inspector /Users/brentjackson/git/github/brent-jackson/sample-ai-sandbox/sample-mcp-server/bin/sample-mcp-server
```

## 3. Using a Custom MCP Client

### Simple Test Client
Here's a basic example of how to connect to your server programmatically:

```go
// client_example.go
package main

import (
    "context"
    "log"
    "os/exec"
    
    "github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
    // Create a client
    client := mcp.NewClient(&mcp.Implementation{
        Name:    "test-client",
        Version: "1.0.0",
    }, nil)
    
    // Connect using command transport
    cmd := exec.Command("./bin/sample-mcp-server")
    transport := &mcp.CommandTransport{Command: cmd}
    
    session, err := client.Connect(context.Background(), transport, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()
    
    // List available tools
    tools, err := session.ListTools(context.Background(), &mcp.ListToolsParams{})
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Available tools: %v", tools)
    
    // Call the calculator tool
    result, err := session.CallTool(context.Background(), &mcp.CallToolParams{
        Name: "calculate",
        Arguments: map[string]any{
            "expression": "2 + 3",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Calculator result: %v", result)
}
```

## 4. Direct JSON-RPC Testing

You can also test the server directly using JSON-RPC messages via stdin/stdout:

### Test the Calculator Tool
```bash
# Start the server in background
./bin/sample-mcp-server &
SERVER_PID=$!

# Send a JSON-RPC message
echo '{
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
}' | ./bin/sample-mcp-server

# Clean up
kill $SERVER_PID
```

## Available Tools and Resources

### Tools
Your server provides these interactive tools:

1. **calculate**
   - Purpose: Basic mathematical calculations
   - Parameters: `expression` (string)
   - Examples: "2 + 3", "sqrt(16)", "sin(pi/2)"

2. **text_transform** 
   - Purpose: Text manipulation
   - Parameters: `text` (string), `operation` (string)
   - Operations: uppercase, lowercase, reverse, word_count, char_count

3. **current_time**
   - Purpose: Get current time in various formats
   - Parameters: `format` (optional), `custom_format` (optional)
   - Formats: iso8601, unix, human_readable, custom

### Resources
Your server provides these static resources:

1. **sample://config** - Server configuration (JSON)
2. **sample://status** - Current server status (JSON)  
3. **sample://docs/readme** - Documentation (Markdown)

## Example Interactions

### Via Claude Desktop
Once configured, you can naturally ask Claude:

- "Use the calculator to compute sqrt(16)"
- "Transform 'Hello World' to lowercase using the text tool"
- "What time is it? Use the unix timestamp format"
- "Show me the server configuration"
- "What's in the documentation resource?"

### Via MCP Inspector
1. Open the inspector web interface
2. Click on "Tools" to see available tools
3. Select a tool and fill in parameters
4. Click "Resources" to browse static content

## Troubleshooting

### Common Issues
1. **Server won't start**: Check that Go is installed and the binary is built
2. **Claude Desktop not connecting**: Verify the config path and restart Claude
3. **Permission errors**: Ensure the binary is executable (`chmod +x bin/sample-mcp-server`)

### Debugging
```bash
# Test if the server builds correctly
go build -o bin/sample-mcp-server main.go

# Test basic functionality
go test -v

# Check server startup
echo '{}' | ./bin/sample-mcp-server 2>&1 | head -10
```

The server implements the full MCP protocol and is ready to be used with any compatible MCP client!