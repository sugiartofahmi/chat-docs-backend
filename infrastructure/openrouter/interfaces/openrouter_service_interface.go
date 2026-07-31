package interfaces

import (
	"context"

	"go-service/infrastructure/openrouter/dtos"
)

type OpenRouterServiceInterface interface {
	CreateEmbedding(ctx context.Context, dto *dtos.EmbeddingRequestDto) *dtos.EmbeddingResponseDto
	CreateChatCompletion(ctx context.Context, dto *dtos.ChatCompletionRequestDto) *dtos.ChatCompletionResponseDto
}
