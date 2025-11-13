package main

import (
	"context"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Example client that demonstrates how to connect to the sample MCP server
func main() {
	// Create a client
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "sample-client",
		Version: "1.0.0",
	}, nil)

	// Get the path to the server binary (assuming we're running from examples/ directory)
	serverPath := filepath.Join("..", "bin", "sample-mcp-server")

	// Connect using command transport to our server
	cmd := exec.Command(serverPath)
	transport := &mcp.CommandTransport{Command: cmd}

	session, err := client.Connect(context.Background(), transport, nil)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer session.Close()

	// List available tools
	log.Println("=== Available Tools ===")
	tools, err := session.ListTools(context.Background(), &mcp.ListToolsParams{})
	if err != nil {
		log.Fatalf("Failed to list tools: %v", err)
	}

	for _, tool := range tools.Tools {
		log.Printf("Tool: %s - %s", tool.Name, tool.Description)
	}

	// Test the calculator tool
	log.Println("\n=== Testing Calculator Tool ===")
	calcResult, err := session.CallTool(context.Background(), &mcp.CallToolParams{
		Name: "calculate",
		Arguments: map[string]any{
			"expression": "2 + 3",
		},
	})
	if err != nil {
		log.Fatalf("Calculator tool failed: %v", err)
	}
	log.Printf("Calculator result: %v", calcResult.Content)

	// Test the text transform tool
	log.Println("\n=== Testing Text Transform Tool ===")
	textResult, err := session.CallTool(context.Background(), &mcp.CallToolParams{
		Name: "text_transform",
		Arguments: map[string]any{
			"text":      "Hello World",
			"operation": "uppercase",
		},
	})
	if err != nil {
		log.Fatalf("Text transform tool failed: %v", err)
	}
	log.Printf("Text transform result: %v", textResult.Content)

	// Test the time tool
	log.Println("\n=== Testing Time Tool ===")
	timeResult, err := session.CallTool(context.Background(), &mcp.CallToolParams{
		Name: "current_time",
		Arguments: map[string]any{
			"format": "human_readable",
		},
	})
	if err != nil {
		log.Fatalf("Time tool failed: %v", err)
	}
	log.Printf("Time result: %v", timeResult.Content)

	// List available resources
	log.Println("\n=== Available Resources ===")
	resources, err := session.ListResources(context.Background(), &mcp.ListResourcesParams{})
	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	for _, resource := range resources.Resources {
		log.Printf("Resource: %s (%s) - %s", resource.URI, resource.MIMEType, resource.Description)
	}

	// Read a resource
	log.Println("\n=== Reading Config Resource ===")
	configResource, err := session.ReadResource(context.Background(), &mcp.ReadResourceParams{
		URI: "sample://config",
	})
	if err != nil {
		log.Fatalf("Failed to read config resource: %v", err)
	}

	if len(configResource.Contents) > 0 {
		log.Printf("Config content (first 200 chars): %.200s...", configResource.Contents[0].Text)
	}

	log.Println("\n=== Client Test Complete ===")
}
