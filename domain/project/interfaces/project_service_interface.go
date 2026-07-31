package interfaces

import (
	"context"

	projectDtos "go-service/domain/project/dtos"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"

	"github.com/google/uuid"
)

type ProjectServiceInterface interface {
	Pagination(ctx context.Context, dto *projectDtos.ProjectQueryRequestDto) *infradtos.PaginationResultDto[entities.ProjectEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.ProjectEntity
	Create(ctx context.Context, dto *projectDtos.ProjectCreateRequestDto) *entities.ProjectEntity
}
