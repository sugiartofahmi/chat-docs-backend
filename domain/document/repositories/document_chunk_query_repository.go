package repositories

import (
	"context"
	"log"

	documentInterfaces "go-service/domain/document/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentChunkQueryRepository struct {
	db                 *gorm.DB
	documentChunkModel *gorm.DB
}

func NewDocumentChunkQueryRepository(db *gorm.DB) documentInterfaces.DocumentChunkQueryRepositoryInterface {
	return &DocumentChunkQueryRepository{
		db:                 db,
		documentChunkModel: db.Model(&entities.DocumentChunkEntity{}),
	}
}

func (repo *DocumentChunkQueryRepository) FindByDocumentId(ctx context.Context, documentId uuid.UUID) []*entities.DocumentChunkEntity {
	var chunks []*entities.DocumentChunkEntity

	query := repo.documentChunkModel.WithContext(ctx)
	if err := query.Where("document_id = ?", documentId).Order("chunk_index").Find(&chunks).Error; err != nil {
		log.Println("Error find document chunks:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return chunks
}
