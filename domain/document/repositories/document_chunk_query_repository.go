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
	"gorm.io/gorm/clause"
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

func (repo *DocumentChunkQueryRepository) FindNearestByProjectId(ctx context.Context, projectId uuid.UUID, queryEmbedding pgvector.Vector, limit int) []*entities.DocumentChunkEntity {
	var chunks []*entities.DocumentChunkEntity

	query := repo.documentChunkModel.WithContext(ctx).
		Joins("JOIN documents ON documents.id = document_chunks.document_id").
		Where("documents.project_id = ? AND document_chunks.embedding IS NOT NULL", projectId).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "document_chunks.embedding <=> ?",
				Vars:               []any{queryEmbedding},
				WithoutParentheses: true,
			},
		}).
		Limit(limit)

	if err := query.Find(&chunks).Error; err != nil {
		log.Println("Error find nearest document chunks:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return chunks
}
