package interfaces

import (
	"context"

	"go-service/entities"
)

type ConversationStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.ConversationEntity) *entities.ConversationEntity
}
