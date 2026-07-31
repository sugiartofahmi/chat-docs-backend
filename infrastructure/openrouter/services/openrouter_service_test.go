package services

import (
	"context"
	"strings"
	"testing"

	"go-service/infrastructure/config"
	"go-service/infrastructure/integrations"
	"go-service/infrastructure/openrouter/dtos"
)

func TestCreateEmbedding(t *testing.T) {
	if config.OpenRouterAPIKey == "" || strings.HasSuffix(config.OpenRouterAPIKey, "-not-set") {
		t.Skip("OPENROUTER_API_KEY not set")
	}

	service := NewOpenRouterService(integrations.NewHttpClient())
	resp := service.CreateEmbedding(context.Background(), &dtos.EmbeddingRequestDto{
		Model: config.OpenRouterEmbeddingModel,
		Input: []string{"hello world"},
	})

	if len(resp.Data) == 0 || len(resp.Data[0].Embedding) == 0 {
		t.Fatal("expected non-empty embedding")
	}
}
