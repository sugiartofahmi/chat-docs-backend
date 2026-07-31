package services

import (
	"context"

	projectConstants "go-service/domain/project/constants"
	projectDtos "go-service/domain/project/dtos"
	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"
	"go-service/infrastructure/exceptions"

	"github.com/google/uuid"
)

type ProjectService struct {
	projectQueryRepository projectInterfaces.ProjectQueryRepositoryInterface
	projectStoreRepository projectInterfaces.ProjectStoreRepositoryInterface
}

func NewProjectService(
	projectQueryRepository projectInterfaces.ProjectQueryRepositoryInterface,
	projectStoreRepository projectInterfaces.ProjectStoreRepositoryInterface,
) projectInterfaces.ProjectServiceInterface {
	return &ProjectService{
		projectQueryRepository: projectQueryRepository,
		projectStoreRepository: projectStoreRepository,
	}
}

func (s *ProjectService) Pagination(ctx context.Context, dto *projectDtos.ProjectQueryRequestDto) *infradtos.PaginationResultDto[entities.ProjectEntity] {
	return s.projectQueryRepository.Pagination(ctx, dto)
}

func (s *ProjectService) Detail(ctx context.Context, id uuid.UUID) *entities.ProjectEntity {
	project := s.projectQueryRepository.FindOneById(ctx, id)
	isProjectNotFound := project == nil
	if isProjectNotFound {
		panic(*exceptions.NotFoundException(projectConstants.PROJECT_NOT_FOUND))
	}
	return project
}

func (s *ProjectService) Create(ctx context.Context, dto *projectDtos.ProjectCreateRequestDto) *entities.ProjectEntity {
	isNameAlreadyUsed := s.projectQueryRepository.IsExistByName(ctx, dto.Name)
	if isNameAlreadyUsed {
		panic(*exceptions.ConflictException(projectConstants.PROJECT_NAME_ALREADY_EXISTS))
	}

	entity := &entities.ProjectEntity{
		Name: dto.Name,
	}
	return s.projectStoreRepository.Create(ctx, entity)
}
