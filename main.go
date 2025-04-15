package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/http"
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

	transport := http.NewHTTPTransport("/")
	transport.WithAddr(fmt.Sprintf(":%d", cfg.HTTP.Port))

	server := mcp_golang.NewServer(transport)
	RegisterMcp(server, tools)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
