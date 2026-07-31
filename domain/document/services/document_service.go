package services

import (
	"context"

	documentConstants "go-service/domain/document/constants"
	documentInterfaces "go-service/domain/document/interfaces"
	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"
	jobInterfaces "go-service/infrastructure/job/interfaces"

	"github.com/google/uuid"
)

type DocumentService struct {
	documentStoreRepository documentInterfaces.DocumentStoreRepositoryInterface
	projectQueryRepository  projectInterfaces.ProjectQueryRepositoryInterface
	jobPool                 jobInterfaces.JobPoolInterface
}

func NewDocumentService(
	documentStoreRepository documentInterfaces.DocumentStoreRepositoryInterface,
	projectQueryRepository projectInterfaces.ProjectQueryRepositoryInterface,
	jobPool jobInterfaces.JobPoolInterface,
) documentInterfaces.DocumentServiceInterface {
	return &DocumentService{
		documentStoreRepository: documentStoreRepository,
		projectQueryRepository:  projectQueryRepository,
		jobPool:                 jobPool,
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
	document := s.documentStoreRepository.Create(ctx, entity)

	s.jobPool.Delegate(documentConstants.DOCUMENT_EXTRACTION_JOB_NAME, documentInterfaces.DocumentExtractionJobPayload{
		DocumentId: document.Id,
		FileBuffer: fileBuffer,
	})

	return document
}
