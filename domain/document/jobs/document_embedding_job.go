package jobs

import (
	"context"
	"encoding/json"
	"fmt"

	documentConstants "go-service/domain/document/constants"
	documentInterfaces "go-service/domain/document/interfaces"
	jobInterfaces "go-service/infrastructure/job/interfaces"
)

type DocumentEmbeddingJob struct {
	documentExtractionService documentInterfaces.DocumentExtractionServiceInterface
}

func NewDocumentEmbeddingJob(documentExtractionService documentInterfaces.DocumentExtractionServiceInterface) jobInterfaces.JobInterface {
	return &DocumentEmbeddingJob{documentExtractionService: documentExtractionService}
}

func (j *DocumentEmbeddingJob) Name() string { return documentConstants.DOCUMENT_EMBEDDING_JOB_NAME }

func (j *DocumentEmbeddingJob) Handle(payload []byte) (err error) {
	var data documentInterfaces.DocumentEmbeddingJobPayload
	if unmarshalErr := json.Unmarshal(payload, &data); unmarshalErr != nil {
		return unmarshalErr
	}

	defer func() {
		if r := recover(); r != nil {
			j.documentExtractionService.MarkFailed(context.Background(), data.DocumentId)
			err = fmt.Errorf("%v", r)
		}
	}()

	j.documentExtractionService.EmbedAndStore(context.Background(), data.DocumentId)
	return nil
}
