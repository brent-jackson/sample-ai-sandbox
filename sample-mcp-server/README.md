# Sample MCP Server

A sample Model Context Protocol (MCP) server implementation in Go using the official [modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk).

## Features

This sample MCP server demonstrates:

- **Tools**: Interactive functions that can be called by MCP clients
  - `calculate`: Basic mathematical operations
  - `text_transform`: Text manipulation utilities  
  - `current_time`: Date/time formatting in various formats

- **Resources**: Static content that can be read by MCP clients
  - Server configuration in JSON format
  - Real-time server status and statistics
  - Documentation in Markdown format

## Quick Start

1. **Prerequisites**: Go 1.19 or later

2. **Clone and build**:
   ```bash
   git clone <repository-url>
   cd sample-ai-sandbox
   go mod download
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

4. **Connect from an MCP client**: Use the provided `mcp-config.json` configuration

## Tools Reference

### calculate
Performs basic mathematical calculations.

**Parameters**:
- `expression` (string, required): Mathematical expression to evaluate

**Examples**:
- `"2 + 3"` → `5`
- `"sqrt(16)"` → `4` 
- `"sin(pi/2)"` → `1`

### text_transform  
Transforms text in various ways.

**Parameters**:
- `text` (string, required): Text to transform
- `operation` (string, required): One of: `uppercase`, `lowercase`, `reverse`, `word_count`, `char_count`

**Examples**:
- `text: "Hello World", operation: "uppercase"` → `"HELLO WORLD"`
- `text: "Hello World", operation: "word_count"` → `"Word count: 2"`

### current_time
Returns current date/time in various formats.

**Parameters**:
- `format` (string, optional): One of: `iso8601` (default), `unix`, `human_readable`, `custom`
- `custom_format` (string, optional): Go time format string (when format is `custom`)

**Examples**:
- `format: "iso8601"` → `"2025-11-13T10:30:00Z"`
- `format: "human_readable"` → `"Wednesday, November 13, 2025 at 10:30 AM UTC"`

## Resources Reference

### sample://config
Returns server configuration as JSON.

### sample://status  
Returns current server status and statistics as JSON.

### sample://docs/readme
Returns this documentation in Markdown format.

## Configuration

The `mcp-config.json` file provides a sample configuration for connecting to this server:

```json
{
  "mcpServers": {
    "sample-mcp-server": {
      "command": "go",
      "args": ["run", "main.go"],
      "cwd": "."
    }
  }
}
```

## Development

The server is structured with:

- `SampleMCPServer`: Main server struct that implements MCP protocol
- Tool handlers: Functions that implement each tool's functionality
- Resource handlers: Functions that provide static content
- MCP SDK integration: Uses official Go SDK for protocol handling

### Adding New Tools

1. Add tool definition in `handleToolList()`
2. Add case in `handleToolCall()` 
3. Implement handler function following the pattern of existing tools

### Adding New Resources

1. Add resource definition in `handleResourceList()`
2. Add case in `handleResourceRead()`
3. Return appropriate content with correct MIME type

## License

This is a sample implementation for educational purposes.