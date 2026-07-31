package interfaces

import (
	"context"

	"go-service/entities"

	"github.com/google/uuid"
)

type DocumentServiceInterface interface {
	Upload(ctx context.Context, projectId uuid.UUID, filename string, fileBuffer []byte) *entities.DocumentEntity
}
