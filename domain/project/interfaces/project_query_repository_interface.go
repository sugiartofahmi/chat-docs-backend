package interfaces

import (
	"context"

	projectDtos "go-service/domain/project/dtos"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"

	"github.com/google/uuid"
)

type ProjectQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *projectDtos.ProjectQueryRequestDto) *infradtos.PaginationResultDto[entities.ProjectEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.ProjectEntity
	IsExistByName(ctx context.Context, name string) bool
	IsExistById(ctx context.Context, id uuid.UUID) bool
}
