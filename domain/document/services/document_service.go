package services

import (
	"context"

	documentConstants "go-service/domain/document/constants"
	documentInterfaces "go-service/domain/document/interfaces"
	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"

	"github.com/google/uuid"
)

type DocumentService struct {
	documentStoreRepository documentInterfaces.DocumentStoreRepositoryInterface
	projectQueryRepository  projectInterfaces.ProjectQueryRepositoryInterface
}

func NewDocumentService(
	documentStoreRepository documentInterfaces.DocumentStoreRepositoryInterface,
	projectQueryRepository projectInterfaces.ProjectQueryRepositoryInterface,
) documentInterfaces.DocumentServiceInterface {
	return &DocumentService{
		documentStoreRepository: documentStoreRepository,
		projectQueryRepository:  projectQueryRepository,
	}
}

func (s *DocumentService) Upload(ctx context.Context, projectId uuid.UUID, filename string, fileBuffer []byte) *entities.DocumentEntity {
	isProjectExist := s.projectQueryRepository.IsExistById(ctx, projectId)
	if !isProjectExist {
		panic(*exceptions.NotFoundException(documentConstants.DOCUMENT_PROJECT_NOT_FOUND))
	}

	entity := &entities.DocumentEntity{
		ProjectId: projectId,
		Filename:  filename,
	}
	return s.documentStoreRepository.Create(ctx, entity)
}
