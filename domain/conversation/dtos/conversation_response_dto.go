package dtos

import (
	"time"

	"go-service/entities"

	"github.com/google/uuid"
)

type ConversationResponseDto struct {
	Id        uuid.UUID  `json:"id"`
	ProjectId uuid.UUID  `json:"project_id"`
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
}

func ConversationResponseDtoFromEntity(conversation *entities.ConversationEntity) *ConversationResponseDto {
	return &ConversationResponseDto{
		Id:        conversation.Id,
		ProjectId: conversation.ProjectId,
		Role:      conversation.Role,
		Content:   conversation.Content,
		CreatedAt: conversation.CreatedAt,
	}
}

func ConversationResponseDtoFromEntities(conversations []*entities.ConversationEntity) []ConversationResponseDto {
	responses := make([]ConversationResponseDto, len(conversations))
	for i, conversation := range conversations {
		responses[i] = *ConversationResponseDtoFromEntity(conversation)
	}
	return responses
}
