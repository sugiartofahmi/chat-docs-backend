package services

import (
	"context"
	"strings"

	conversationConstants "go-service/domain/conversation/constants"
	conversationDtos "go-service/domain/conversation/dtos"
	conversationInterfaces "go-service/domain/conversation/interfaces"
	documentInterfaces "go-service/domain/document/interfaces"
	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"
	"go-service/infrastructure/exceptions"
	openrouterDtos "go-service/infrastructure/openrouter/dtos"
	openrouterInterfaces "go-service/infrastructure/openrouter/interfaces"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
)

type ConversationService struct {
	conversationStoreRepository  conversationInterfaces.ConversationStoreRepositoryInterface
	conversationQueryRepository  conversationInterfaces.ConversationQueryRepositoryInterface
	projectQueryRepository       projectInterfaces.ProjectQueryRepositoryInterface
	documentChunkQueryRepository documentInterfaces.DocumentChunkQueryRepositoryInterface
	openRouterService            openrouterInterfaces.OpenRouterServiceInterface
	openRouterEmbeddingModel     string
	openRouterChatModel          string
}

func NewConversationService(
	conversationStoreRepository conversationInterfaces.ConversationStoreRepositoryInterface,
	conversationQueryRepository conversationInterfaces.ConversationQueryRepositoryInterface,
	projectQueryRepository projectInterfaces.ProjectQueryRepositoryInterface,
	documentChunkQueryRepository documentInterfaces.DocumentChunkQueryRepositoryInterface,
	openRouterService openrouterInterfaces.OpenRouterServiceInterface,
	openRouterEmbeddingModel string,
	openRouterChatModel string,
) conversationInterfaces.ConversationServiceInterface {
	return &ConversationService{
		conversationStoreRepository:  conversationStoreRepository,
		conversationQueryRepository:  conversationQueryRepository,
		projectQueryRepository:       projectQueryRepository,
		documentChunkQueryRepository: documentChunkQueryRepository,
		openRouterService:            openRouterService,
		openRouterEmbeddingModel:     openRouterEmbeddingModel,
		openRouterChatModel:          openRouterChatModel,
	}
}

func (s *ConversationService) Pagination(ctx context.Context, projectId uuid.UUID, dto *conversationDtos.ConversationQueryRequestDto) *infradtos.PaginationResultDto[entities.ConversationEntity] {
	if !s.projectQueryRepository.IsExistById(ctx, projectId) {
		panic(*exceptions.NotFoundException(conversationConstants.CONVERSATION_PROJECT_NOT_FOUND))
	}

	return s.conversationQueryRepository.Pagination(ctx, projectId, dto)
}

func (s *ConversationService) Ask(ctx context.Context, projectId uuid.UUID, content string) *entities.ConversationEntity {
	if !s.projectQueryRepository.IsExistById(ctx, projectId) {
		panic(*exceptions.NotFoundException(conversationConstants.CONVERSATION_PROJECT_NOT_FOUND))
	}

	s.conversationStoreRepository.Create(ctx, &entities.ConversationEntity{
		ProjectId: projectId,
		Role:      conversationConstants.CONVERSATION_ROLE_USER,
		Content:   content,
	})

	queryEmbeddingResponse := s.openRouterService.CreateEmbedding(ctx, &openrouterDtos.EmbeddingRequestDto{
		Model: s.openRouterEmbeddingModel,
		Input: []string{content},
	})

	var queryEmbedding pgvector.Vector
	if len(queryEmbeddingResponse.Data) > 0 {
		vector := make([]float32, len(queryEmbeddingResponse.Data[0].Embedding))
		for i, value := range queryEmbeddingResponse.Data[0].Embedding {
			vector[i] = float32(value)
		}
		queryEmbedding = pgvector.NewVector(vector)
	}

	chunks := s.documentChunkQueryRepository.FindNearestByProjectId(ctx, projectId, queryEmbedding, conversationConstants.RETRIEVAL_TOP_K)

	var replyContent string
	if len(chunks) == 0 {
		replyContent = conversationConstants.CONVERSATION_NO_DOCUMENTS
	} else {
		contexts := make([]string, len(chunks))
		for i, chunk := range chunks {
			contexts[i] = chunk.Content
		}

		chatResponse := s.openRouterService.CreateChatCompletion(ctx, &openrouterDtos.ChatCompletionRequestDto{
			Model: s.openRouterChatModel,
			Messages: []openrouterDtos.ChatMessageDto{
				{
					Role:    "system",
					Content: "Answer the user's question using only the context below. If the context doesn't contain the answer, say so.\n\nContext:\n" + strings.Join(contexts, "\n---\n"),
				},
				{
					Role:    "user",
					Content: content,
				},
			},
		})

		if len(chatResponse.Choices) > 0 {
			replyContent = chatResponse.Choices[0].Message.Content
		}
	}

	return s.conversationStoreRepository.Create(ctx, &entities.ConversationEntity{
		ProjectId: projectId,
		Role:      conversationConstants.CONVERSATION_ROLE_ASSISTANT,
		Content:   replyContent,
	})
}
