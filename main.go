package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	var _ = godotenv.Load(".env")
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded")
	ctx := context.Background()

	vector := NewVectorAction(cfg.Vector)
	tools := NewMcpTools(vector, ctx)

	s := server.NewMCPServer(
		"Memory MCP",
		"0.0.1",
		server.WithLogging(),
		server.WithRecovery(),
	)
	sse := server.NewSSEServer(s)
	RegisterMcp(s, tools)
	if err := sse.Start(fmt.Sprintf(":%d", cfg.HTTP.Port)); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
