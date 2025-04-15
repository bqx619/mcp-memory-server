package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
)

type McpToolsImpl interface {
	StoreMemory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	RetrieveMemory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	DeleteMemory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

type McpTools struct {
	vector *VectorAction
	ctx    context.Context
}

func NewMcpTools(vector *VectorAction, ctx context.Context) McpToolsImpl {
	return &McpTools{
		vector: vector,
		ctx:    ctx,
	}
}

func (t *McpTools) StoreMemory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	content := request.Params.Arguments["content"].(string)
	// generate a random short-id
	_id := uuid.New().String()[:8]

	err := t.vector.Upsert(t.ctx, _id, content, nil)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(fmt.Sprintf("Successfully stored memory with ID: %s", _id)), nil
}

func (t *McpTools) RetrieveMemory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	n_results := 5
	query := request.Params.Arguments["query"].(string)
	if v, ok := request.Params.Arguments["n_results"]; ok {
		n_results = int(v.(float64))
	}

	results, err := t.vector.Search(t.ctx, query, n_results)
	if err != nil {
		return nil, err
	}

	content := make([]string, len(results))
	for i, result := range results {
		content[i] = fmt.Sprintf("Memory %d:\nContent: %s\nMemory Id: %s\nRelevance Score: %.2f", i+1, result.Data, result.Id, result.Score)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Found the following memories:\n%s", strings.Join(content, "\n"))), nil
}

func (t *McpTools) DeleteMemory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	memory_id := request.Params.Arguments["memory_id"].(string)
	err := t.vector.Delete(t.ctx, memory_id)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(fmt.Sprintf("Successfully deleted memory with ID: %s", memory_id)), nil
}
