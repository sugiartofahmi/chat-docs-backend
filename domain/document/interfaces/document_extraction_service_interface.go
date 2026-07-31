package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type DocumentExtractionServiceInterface interface {
	Extract(ctx context.Context, documentId uuid.UUID, fileBuffer []byte)
	EmbedAndStore(ctx context.Context, documentId uuid.UUID)
	MarkFailed(ctx context.Context, documentId uuid.UUID)
}
