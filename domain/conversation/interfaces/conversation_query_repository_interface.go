package interfaces

import (
	"context"

	conversationDtos "go-service/domain/conversation/dtos"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"

	"github.com/google/uuid"
)

type ConversationQueryRepositoryInterface interface {
	Pagination(ctx context.Context, projectId uuid.UUID, dto *conversationDtos.ConversationQueryRequestDto) *infradtos.PaginationResultDto[entities.ConversationEntity]
}
