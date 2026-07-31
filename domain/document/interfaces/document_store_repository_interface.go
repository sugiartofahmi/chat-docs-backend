package interfaces

import (
	"context"

	"go-service/entities"

	"gorm.io/gorm"
)

type DocumentStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.DocumentEntity) *entities.DocumentEntity
	WithTransaction(tx *gorm.DB) DocumentStoreRepositoryInterface
}
