package interfaces

import (
	"context"

	"go-service/entities"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
)

type DocumentChunkQueryRepositoryInterface interface {
	FindByDocumentId(ctx context.Context, documentId uuid.UUID) []*entities.DocumentChunkEntity
	FindNearestByProjectId(ctx context.Context, projectId uuid.UUID, queryEmbedding pgvector.Vector, limit int) []*entities.DocumentChunkEntity
}
