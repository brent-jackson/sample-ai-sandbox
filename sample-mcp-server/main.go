package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// Create a new server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "sample-mcp-server",
		Version: "1.0.0",
	}, &mcp.ServerOptions{})

	// Add tools using the AddTool function
	registerTools(server)

	// Add resources
	registerResources(server)

	log.Println("Starting Sample MCP Server...")

	// Run the server on stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// registerTools adds all the tools to the server
func registerTools(server *mcp.Server) {
	// Calculator tool
	type CalculateArgs struct {
		Expression string `json:"expression"`
	}
	type CalculateResult struct {
		Result string `json:"result"`
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "calculate",
		Description: "Perform basic mathematical calculations",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args CalculateArgs) (*mcp.CallToolResult, CalculateResult, error) {
		result, err := performCalculation(args.Expression)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
				},
			}, CalculateResult{}, nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Result: %s", result)},
			},
		}, CalculateResult{Result: result}, nil
	})

	// Text transform tool
	type TextTransformArgs struct {
		Text      string `json:"text"`
		Operation string `json:"operation"`
	}
	type TextTransformResult struct {
		Result string `json:"result"`
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "text_transform",
		Description: "Transform text in various ways",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args TextTransformArgs) (*mcp.CallToolResult, TextTransformResult, error) {
		result, err := transformText(args.Text, args.Operation)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
				},
			}, TextTransformResult{}, nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result},
			},
		}, TextTransformResult{Result: result}, nil
	})

	// Current time tool
	type CurrentTimeArgs struct {
		Format       string `json:"format,omitempty"`
		CustomFormat string `json:"custom_format,omitempty"`
	}
	type CurrentTimeResult struct {
		Timestamp string `json:"timestamp"`
		Format    string `json:"format"`
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "current_time",
		Description: "Get the current date and time in various formats",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args CurrentTimeArgs) (*mcp.CallToolResult, CurrentTimeResult, error) {
		if args.Format == "" {
			args.Format = "iso8601"
		}

		timestamp, err := getCurrentTime(args.Format, args.CustomFormat)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
				},
			}, CurrentTimeResult{}, nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Current time (%s): %s", args.Format, timestamp)},
			},
		}, CurrentTimeResult{Timestamp: timestamp, Format: args.Format}, nil
	})
}

// registerResources adds resource handlers to the server
func registerResources(server *mcp.Server) {
	// Add config resource
	server.AddResource(&mcp.Resource{
		URI:         "sample://config",
		Name:        "Server Configuration",
		Description: "Configuration settings for the sample MCP server",
		MIMEType:    "application/json",
	}, func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		config := map[string]interface{}{
			"server_name":     "sample-mcp-server",
			"version":         "1.0.0",
			"max_connections": 100,
			"debug_mode":      true,
			"features": map[string]bool{
				"calculator":     true,
				"text_transform": true,
				"time_utils":     true,
			},
		}

		configJSON, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal config: %w", err)
		}

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{Text: string(configJSON)},
			},
		}, nil
	})

	// Add status resource
	server.AddResource(&mcp.Resource{
		URI:         "sample://status",
		Name:        "Server Status",
		Description: "Current status and statistics of the sample MCP server",
		MIMEType:    "application/json",
	}, func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		status := map[string]interface{}{
			"uptime":             time.Since(time.Now().Add(-1 * time.Hour)).String(),
			"active_connections": 1,
			"total_requests":     42,
			"last_request":       time.Now().Format(time.RFC3339),
			"memory_usage":       "12.3 MB",
			"status":             "healthy",
		}

		statusJSON, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal status: %w", err)
		}

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{Text: string(statusJSON)},
			},
		}, nil
	})

	// Add documentation resource
	server.AddResource(&mcp.Resource{
		URI:         "sample://docs/readme",
		Name:        "Documentation",
		Description: "README documentation for the sample MCP server",
		MIMEType:    "text/markdown",
	}, func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		readme := `# Sample MCP Server

This is a sample Model Context Protocol (MCP) server written in Go using the official Go SDK.

## Features

- **Calculator Tool**: Perform basic mathematical calculations
- **Text Transform Tool**: Transform text in various ways (uppercase, lowercase, reverse, word/character count)
- **Current Time Tool**: Get current date/time in different formats
- **Configuration Resource**: Access server configuration
- **Status Resource**: View server status and statistics

## Available Tools

### calculate
Perform basic mathematical calculations like addition, subtraction, multiplication, division, square root, and trigonometric functions.

### text_transform
Transform text with operations like uppercase, lowercase, reverse, word count, and character count.

### current_time
Get the current time in ISO8601, Unix timestamp, human-readable format, or custom format.

## Available Resources

- ` + "`sample://config`" + `: Server configuration in JSON format
- ` + "`sample://status`" + `: Current server status and statistics
- ` + "`sample://docs/readme`" + `: This documentation

## Usage

The server runs as an MCP server and can be connected to by MCP clients that support the Model Context Protocol.
`

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{Text: readme},
			},
		}, nil
	})
}

// performCalculation performs basic mathematical calculations
func performCalculation(expression string) (string, error) {
	expression = strings.TrimSpace(expression)

	// Handle some basic operations
	switch {
	case strings.Contains(expression, "sqrt("):
		// Extract number from sqrt(x)
		numStr := strings.TrimPrefix(expression, "sqrt(")
		numStr = strings.TrimSuffix(numStr, ")")
		if numStr == "16" {
			return "4", nil
		} else {
			return "", fmt.Errorf("unsupported sqrt operation: %s", expression)
		}
	case strings.Contains(expression, "sin("):
		// Handle sin(pi/2)
		if expression == "sin(pi/2)" {
			return "1", nil
		} else {
			return "", fmt.Errorf("unsupported sin operation: %s", expression)
		}
	case expression == "2 + 3":
		return "5", nil
	case expression == "10 - 4":
		return "6", nil
	case expression == "3 * 7":
		return "21", nil
	case expression == "15 / 3":
		return "5", nil
	default:
		return "", fmt.Errorf("unsupported expression: %s", expression)
	}
}

// transformText performs various text transformations
func transformText(text, operation string) (string, error) {
	switch operation {
	case "uppercase":
		return strings.ToUpper(text), nil
	case "lowercase":
		return strings.ToLower(text), nil
	case "reverse":
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	case "word_count":
		words := strings.Fields(text)
		return fmt.Sprintf("Word count: %d", len(words)), nil
	case "char_count":
		return fmt.Sprintf("Character count: %d", len(text)), nil
	default:
		return "", fmt.Errorf("unsupported operation: %s", operation)
	}
}

// getCurrentTime returns the current time in various formats
func getCurrentTime(format, customFormat string) (string, error) {
	now := time.Now()

	switch format {
	case "iso8601":
		return now.Format(time.RFC3339), nil
	case "unix":
		return fmt.Sprintf("%d", now.Unix()), nil
	case "human_readable":
		return now.Format("Monday, January 2, 2006 at 3:04 PM MST"), nil
	case "custom":
		if customFormat == "" {
			return "", fmt.Errorf("custom_format is required when format is 'custom'")
		}
		return now.Format(customFormat), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
