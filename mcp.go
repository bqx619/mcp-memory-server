package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

type McpToolsImpl interface {
	StoreMemory(arguments StoreMemoryArguments) (*mcp_golang.ToolResponse, error)
}

type McpTools struct {
	vector *VectorAction
	ctx    context.Context
}

func NewMcpTools(vector *VectorAction, ctx context.Context) *McpTools {
	return &McpTools{
		vector: vector,
		ctx:    ctx,
	}
}

type StoreMemoryArguments struct {
	Content string `json:"content" jsonschema:"required,description=The memory content to store, such as a fact, note, or piece of information."`
}

func (t *McpTools) StoreMemory(arguments StoreMemoryArguments) (*mcp_golang.ToolResponse, error) {

	// generate a random short-id
	_id := uuid.New().String()[:8]

	err := t.vector.Upsert(t.ctx, _id, arguments.Content, nil)
	if err != nil {
		return nil, err
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Successfully stored memory with ID: %s", _id))), nil
}

type RetrieveMemoryArguments struct {
	Query     string `json:"query" jsonschema:"required,description=Search query to find relevant memories based on content."`
	N_Results *int   `json:"n_results" jsonschema:"description=Maximum number of results to return, default is 5."`
}

func (t *McpTools) RetrieveMemory(arguments RetrieveMemoryArguments) (*mcp_golang.ToolResponse, error) {
	n_results := 5

	if arguments.N_Results != nil {
		n_results = *arguments.N_Results
	}

	results, err := t.vector.Search(t.ctx, arguments.Query, n_results)
	if err != nil {
		return nil, err
	}

	content := make([]string, len(results))
	for i, result := range results {
		content[i] = fmt.Sprintf("Memory %d:\nContent: %s\nId: %s\nRelevance Score: %.2f", i+1, result.Data, result.Id, result.Score)
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Found the following memories:\n%s", strings.Join(content, "\n")))), nil
}

type DeleteMemoryArguments struct {
	MemoryId string `json:"memory_id" jsonschema:"required,description=The ID of the memory to delete."`
}

func (t *McpTools) DeleteMemory(arguments DeleteMemoryArguments) (*mcp_golang.ToolResponse, error) {
	err := t.vector.Delete(t.ctx, arguments.MemoryId)
	if err != nil {
		return nil, err
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Successfully deleted memory with ID: %s", arguments.MemoryId))), nil
}
