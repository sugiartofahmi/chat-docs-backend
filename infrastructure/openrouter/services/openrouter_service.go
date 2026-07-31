package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"go-service/infrastructure/config"
	"go-service/infrastructure/exceptions"
	"go-service/infrastructure/integrations"
	"go-service/infrastructure/openrouter/constants"
	"go-service/infrastructure/openrouter/dtos"
	"go-service/infrastructure/openrouter/interfaces"
)

type OpenRouterService struct {
	httpClient *integrations.HttpClient
}

func NewOpenRouterService(httpClient *integrations.HttpClient) interfaces.OpenRouterServiceInterface {
	return &OpenRouterService{httpClient: httpClient}
}

func (s *OpenRouterService) CreateEmbedding(_ context.Context, dto *dtos.EmbeddingRequestDto) *dtos.EmbeddingResponseDto {
	s.httpClient.SetHeaders(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", config.OpenRouterAPIKey)})
	if config.OpenRouterAppURL != "" {
		s.httpClient.SetHeaders(map[string]string{"HTTP-Referer": config.OpenRouterAppURL})
	}
	if config.OpenRouterAppTitle != "" {
		s.httpClient.SetHeaders(map[string]string{"X-Title": config.OpenRouterAppTitle})
	}

	body, err := json.Marshal(dto)
	if err != nil {
		panic(*exceptions.ServerErrorException(err))
	}

	respBody, err := s.httpClient.Post(config.OpenRouterBaseURL+constants.OpenRouterEmbeddingEndpoint, bytes.NewReader(body))
	if err != nil {
		panic(*exceptions.ServerErrorException(fmt.Errorf("%s: %w", constants.OpenRouterEmbeddingAPIError, err)))
	}

	var response dtos.EmbeddingResponseDto
	if err := json.Unmarshal(respBody, &response); err != nil {
		panic(*exceptions.ServerErrorException(err))
	}
	return &response
}
