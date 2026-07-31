package interfaces

import "github.com/google/uuid"

type DocumentExtractionJobPayload struct {
	DocumentId uuid.UUID `json:"document_id"`
	FileBuffer []byte    `json:"file_buffer"`
}

type DocumentEmbeddingJobPayload struct {
	DocumentId uuid.UUID `json:"document_id"`
}
