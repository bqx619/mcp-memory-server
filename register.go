package main

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMcp(server *server.MCPServer, tools McpToolsImpl) {

	storeMemory := mcp.NewTool(
		"store_memory",
		mcp.WithDescription(`Store new information.

Examples:
{
	"content": "Memory content",
}`),
		mcp.WithString("content", mcp.Required(), mcp.Description("The memory content to store, such as a fact, note, or piece of information.")),
	)

	retrieveMemory := mcp.NewTool(
		"retrieve_memory",
		mcp.WithDescription(`Find relevant memories based on query.

Examples:
{
	"query": "find this memory",
	"n_results": 5
}`),
		mcp.WithString("query", mcp.Required(), mcp.Description("Search query to find relevant memories based on content.")),
		mcp.WithNumber("n_results", mcp.DefaultNumber(5), mcp.Description("Maximum number of results to return.")),
	)

	deleteMemory := mcp.NewTool(
		"delete_memory",
		mcp.WithDescription(`Delete a specific memory by its id.

Examples:
{	
	"memory_id": "a1b2c3d4..."
}`),
		mcp.WithString("memory_id", mcp.Required(), mcp.Description("The ID of the memory to delete.")),
	)

	server.AddTool(storeMemory, tools.StoreMemory)
	server.AddTool(retrieveMemory, tools.RetrieveMemory)
	server.AddTool(deleteMemory, tools.DeleteMemory)

}
