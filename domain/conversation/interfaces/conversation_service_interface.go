package interfaces

import (
	"context"

	conversationDtos "go-service/domain/conversation/dtos"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"

	"github.com/google/uuid"
)

type ConversationServiceInterface interface {
	Ask(ctx context.Context, projectId uuid.UUID, content string) *entities.ConversationEntity
	Pagination(ctx context.Context, projectId uuid.UUID, dto *conversationDtos.ConversationQueryRequestDto) *infradtos.PaginationResultDto[entities.ConversationEntity]
}
