package repositories

import (
	"context"
	"log"

	documentInterfaces "go-service/domain/document/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type DocumentChunkStoreRepository struct {
	db                 *gorm.DB
	documentChunkModel *gorm.DB
}

func NewDocumentChunkStoreRepository(db *gorm.DB) documentInterfaces.DocumentChunkStoreRepositoryInterface {
	return &DocumentChunkStoreRepository{
		db:                 db,
		documentChunkModel: db.Model(&entities.DocumentChunkEntity{}),
	}
}

func (repo *DocumentChunkStoreRepository) WithTransaction(tx *gorm.DB) documentInterfaces.DocumentChunkStoreRepositoryInterface {
	return &DocumentChunkStoreRepository{
		db:                 tx,
		documentChunkModel: tx.Model(&entities.DocumentChunkEntity{}),
	}
}

func (repo *DocumentChunkStoreRepository) CreateBatch(ctx context.Context, chunks []*entities.DocumentChunkEntity) []*entities.DocumentChunkEntity {
	query := repo.documentChunkModel.WithContext(ctx)

	if err := query.Create(&chunks).Error; err != nil {
		log.Println("Error create document chunks:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return chunks
}

func (repo *DocumentChunkStoreRepository) UpdateEmbedding(ctx context.Context, id uuid.UUID, embedding *pgvector.Vector) *entities.DocumentChunkEntity {
	query := repo.documentChunkModel.WithContext(ctx)

	entity := &entities.DocumentChunkEntity{Id: id, Embedding: embedding}
	if err := query.Where("id = ?", id).Update("embedding", embedding).Error; err != nil {
		log.Println("Error update document chunk embedding:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}
