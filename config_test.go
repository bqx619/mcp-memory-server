package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

func TestLoadConfig(t *testing.T) {
	t.Logf("VECTOR_PROVIDER: %s", os.Getenv("VECTOR_PROVIDER"))
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	t.Logf("Config: %+v", cfg)
}
