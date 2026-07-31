package interfaces

import (
	"context"

	"go-service/entities"

	"github.com/google/uuid"
)

type DocumentChunkQueryRepositoryInterface interface {
	FindByDocumentId(ctx context.Context, documentId uuid.UUID) []*entities.DocumentChunkEntity
}
