package interfaces

import (
	"context"

	"go-service/entities"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type DocumentChunkStoreRepositoryInterface interface {
	CreateBatch(ctx context.Context, chunks []*entities.DocumentChunkEntity) []*entities.DocumentChunkEntity
	UpdateEmbedding(ctx context.Context, id uuid.UUID, embedding *pgvector.Vector) *entities.DocumentChunkEntity
	MarkFailed(ctx context.Context, documentId uuid.UUID, errorMessage string)
	WithTransaction(tx *gorm.DB) DocumentChunkStoreRepositoryInterface
}
