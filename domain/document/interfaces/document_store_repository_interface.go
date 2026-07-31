package interfaces

import (
	"context"

	"go-service/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.DocumentEntity) *entities.DocumentEntity
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) *entities.DocumentEntity
	WithTransaction(tx *gorm.DB) DocumentStoreRepositoryInterface
}
