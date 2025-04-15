package main

import (
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func RegisterMcp(server *mcp_golang.Server, tools *McpTools) {

	// store_memory
	server.RegisterTool(
		"store_memory",
		`Store new information.

		Examples:
		{
			"content": "Memory content",
		}`,
		tools.StoreMemory)

	server.RegisterTool(
		"retrieve_memory",
		`Find relevant memories based on query.

		Example:
		{
			"query": "find this memory",
			"n_results": 5
		}`,
		tools.RetrieveMemory)

	server.RegisterTool(
		"delete_memory",
		`Delete a specific memory by its hash.

		Example:
		{
			"memory_id": "a1b2c3d4..."
		}`,
		tools.DeleteMemory)
}
