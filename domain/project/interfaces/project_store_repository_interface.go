package interfaces

import (
	"context"

	"go-service/entities"

	"gorm.io/gorm"
)

type ProjectStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.ProjectEntity) *entities.ProjectEntity
	WithTransaction(tx *gorm.DB) ProjectStoreRepositoryInterface
}
