package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/upstash/vector-go"
)

var _ = godotenv.Load(".env")

func TestNewVectorAction(t *testing.T) {

}

func TestVectorAction_Upsert(t *testing.T) {
	c := &ConfigVector{
		Provider: "upstash",
		URL:      os.Getenv("VECTOR_URL"),
		Token:    os.Getenv("VECTOR_TOKEN"),
	}
	ins := NewVectorAction(c)
	err := ins.Upsert(context.Background(), "1", "test2", map[string]any{"test": "test"})
	if err != nil {
		t.Fatalf("failed to upsert: %v", err)
	}
	results, err := ins.Search(context.Background(), "test", 1)
	if err != nil {
		t.Fatalf("failed to search: %v", err)
	}
	t.Logf("results: %v", results)
	err = ins.Delete(context.Background(), "1")
	if err != nil {
		t.Fatalf("failed to delete: %v", err)
	}
	results, err = ins.Search(context.Background(), "test", 10)
	if err != nil {
		t.Fatalf("failed to search: %v", err)
	}
	t.Logf("results: %v", results)
}

func TestVectorAction_UpsertMany(t *testing.T) {
	c := &ConfigVector{
		Provider: "upstash",
		URL:      os.Getenv("VECTOR_URL"),
		Token:    os.Getenv("VECTOR_TOKEN"),
	}
	ins := NewVectorAction(c)
	// read test_data all file
	files, err := filepath.Glob("test_data/*.json")
	if err != nil {
		t.Fatalf("failed to read test_data: %v", err)
	}
	udata := make([]vector.UpsertData, 0)
	for _, file := range files {
		t.Logf("reading file: %s", file)
		content, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		var data map[string]any
		err = json.Unmarshal(content, &data)
		if err != nil {
			t.Fatalf("failed to unmarshal file: %v", err)
		}
		t.Logf("content: %s", string(content))

		for _, item := range data["data"].([]any) {
			itemMap := item.(map[string]any)
			udata = append(udata, vector.UpsertData{
				Id:   fmt.Sprintf("%d", int64(itemMap["id"].(float64))),
				Data: itemMap["view"].(string),
				Metadata: map[string]any{
					"significance": itemMap["significance"].(string),
				},
			})
		}

		err = ins.UpsertMany(context.Background(), udata)
		if err != nil {
			t.Fatalf("failed to upsert many: %v", err)
		}
	}

}
