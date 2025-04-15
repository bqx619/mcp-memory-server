package main

import (
	"context"

	"github.com/upstash/vector-go"
)

type VectorAction struct {
	Config *ConfigVector

	client *vector.Index
}

func NewVectorAction(config *ConfigVector) *VectorAction {
	if config.Provider != "upstash" {
		panic("unsupported vector provider: " + config.Provider)
	}
	return &VectorAction{
		Config: config,
		client: vector.NewIndex(config.URL, config.Token),
	}
}

func (v *VectorAction) Upsert(ctx context.Context, id, data string, metadata map[string]any) error {
	return v.client.UpsertData(vector.UpsertData{
		Id:       id,
		Data:     data,
		Metadata: metadata,
	})
}

func (v *VectorAction) UpsertMany(ctx context.Context, data []vector.UpsertData) error {
	return v.client.UpsertDataMany(data)
}

func (v *VectorAction) Search(ctx context.Context, data string, topk int) ([]vector.VectorScore, error) {
	results, err := v.client.QueryData(vector.QueryData{
		Data:        data,
		TopK:        topk,
		IncludeData: true,
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (v *VectorAction) Delete(ctx context.Context, id string) error {
	_, err := v.client.Delete(id)
	return err
}

func (v *VectorAction) DeleteMany(ctx context.Context, ids []string) error {
	_, err := v.client.DeleteMany(ids)
	return err
}
