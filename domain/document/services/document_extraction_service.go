package services

import (
	"context"
	"fmt"
	"strings"

	documentConstants "go-service/domain/document/constants"
	documentInterfaces "go-service/domain/document/interfaces"
	"go-service/entities"
	"go-service/infrastructure/pdfparser"
	"go-service/infrastructure/textchunk"

	jobInterfaces "go-service/infrastructure/job/interfaces"
	openrouterDtos "go-service/infrastructure/openrouter/dtos"
	openrouterInterfaces "go-service/infrastructure/openrouter/interfaces"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type DocumentExtractionService struct {
	db                           *gorm.DB
	documentStoreRepository      documentInterfaces.DocumentStoreRepositoryInterface
	documentChunkStoreRepository documentInterfaces.DocumentChunkStoreRepositoryInterface
	documentChunkQueryRepository documentInterfaces.DocumentChunkQueryRepositoryInterface
	openRouterService            openrouterInterfaces.OpenRouterServiceInterface
	openRouterEmbeddingModel     string
	jobPool                      jobInterfaces.JobPoolInterface
}

func NewDocumentExtractionService(
	db *gorm.DB,
	documentStoreRepository documentInterfaces.DocumentStoreRepositoryInterface,
	documentChunkStoreRepository documentInterfaces.DocumentChunkStoreRepositoryInterface,
	documentChunkQueryRepository documentInterfaces.DocumentChunkQueryRepositoryInterface,
	openRouterService openrouterInterfaces.OpenRouterServiceInterface,
	openRouterEmbeddingModel string,
	jobPool jobInterfaces.JobPoolInterface,
) documentInterfaces.DocumentExtractionServiceInterface {
	return &DocumentExtractionService{
		db:                           db,
		documentStoreRepository:      documentStoreRepository,
		documentChunkStoreRepository: documentChunkStoreRepository,
		documentChunkQueryRepository: documentChunkQueryRepository,
		openRouterService:            openRouterService,
		openRouterEmbeddingModel:     openRouterEmbeddingModel,
		jobPool:                      jobPool,
	}
}

// Extract parses the PDF, chunks the text, and bulk-creates the chunk rows
// without embeddings yet, then delegates the embedding stage as its own job.
func (s *DocumentExtractionService) Extract(ctx context.Context, documentId uuid.UUID, fileBuffer []byte) {
	text, err := pdfparser.ExtractText(fileBuffer)
	if err != nil {
		s.MarkFailed(ctx, documentId, err.Error())
		return
	}
	if strings.TrimSpace(text) == "" {
		s.MarkFailed(ctx, documentId, documentConstants.DOCUMENT_EMPTY_TEXT)
		return
	}

	// TODO: chunk size/overlap and vector(1536) dimension are both tied to
	// documentConstants + the current OpenRouter embedding model; changing
	// the model later means revisiting both.
	chunks := textchunk.Chunk(text, documentConstants.CHUNK_SIZE, documentConstants.CHUNK_OVERLAP)
	if len(chunks) == 0 {
		s.MarkFailed(ctx, documentId, documentConstants.DOCUMENT_CHUNK_EMPTY)
		return
	}

	chunkEntities := make([]*entities.DocumentChunkEntity, len(chunks))
	for i, chunk := range chunks {
		chunkEntities[i] = &entities.DocumentChunkEntity{
			DocumentId: documentId,
			ChunkIndex: i,
			Content:    chunk,
		}
	}
	s.documentChunkStoreRepository.CreateBatch(ctx, chunkEntities)

	s.jobPool.Delegate(documentConstants.DOCUMENT_EMBEDDING_JOB_NAME, documentInterfaces.DocumentEmbeddingJobPayload{
		DocumentId: documentId,
	})
}

// EmbedAndStore fetches the chunk rows Extract just created, embeds them all
// in a single batched call, fills in their embeddings, and marks the document ready.
func (s *DocumentExtractionService) EmbedAndStore(ctx context.Context, documentId uuid.UUID) {
	chunks := s.documentChunkQueryRepository.FindByDocumentId(ctx, documentId)
	if len(chunks) == 0 {
		s.MarkFailed(ctx, documentId, documentConstants.DOCUMENT_NO_CHUNKS_TO_EMBED)
		return
	}

	contents := make([]string, len(chunks))
	for i, chunk := range chunks {
		contents[i] = chunk.Content
	}

	embeddings, err := s.embedChunks(ctx, contents)
	if err != nil {
		s.MarkFailed(ctx, documentId, err.Error())
		return
	}

	s.db.Transaction(func(tx *gorm.DB) error {
		chunkStoreTx := s.documentChunkStoreRepository.WithTransaction(tx)
		for i, chunk := range chunks {
			vector := pgvector.NewVector(embeddings[i])
			chunk.Embedding = &vector
			chunkStoreTx.UpdateEmbedding(ctx, chunk.Id, chunk.Embedding)
		}
		s.documentStoreRepository.WithTransaction(tx).UpdateStatus(ctx, documentId, documentConstants.DOCUMENT_STATUS_READY)
		return nil
	})
}

func (s *DocumentExtractionService) MarkFailed(ctx context.Context, documentId uuid.UUID, errorMessage string) {
	s.documentStoreRepository.MarkFailed(ctx, documentId, errorMessage)
	s.documentChunkStoreRepository.MarkFailed(ctx, documentId, errorMessage)
}

func (s *DocumentExtractionService) embedChunks(ctx context.Context, chunks []string) (embeddings [][]float32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	response := s.openRouterService.CreateEmbedding(ctx, &openrouterDtos.EmbeddingRequestDto{
		Model: s.openRouterEmbeddingModel,
		Input: chunks,
	})

	embeddings = make([][]float32, len(chunks))
	for _, item := range response.Data {
		if item.Index < 0 || item.Index >= len(embeddings) {
			continue
		}
		vector := make([]float32, len(item.Embedding))
		for i, value := range item.Embedding {
			vector[i] = float32(value)
		}
		embeddings[item.Index] = vector
	}

	return embeddings, nil
}
