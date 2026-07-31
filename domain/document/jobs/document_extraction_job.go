package jobs

import (
	"context"
	"encoding/json"
	"fmt"

	documentConstants "go-service/domain/document/constants"
	documentInterfaces "go-service/domain/document/interfaces"
	jobInterfaces "go-service/infrastructure/job/interfaces"
)

type DocumentExtractionJob struct {
	documentExtractionService documentInterfaces.DocumentExtractionServiceInterface
}

func NewDocumentExtractionJob(documentExtractionService documentInterfaces.DocumentExtractionServiceInterface) jobInterfaces.JobInterface {
	return &DocumentExtractionJob{documentExtractionService: documentExtractionService}
}

func (j *DocumentExtractionJob) Name() string { return documentConstants.DOCUMENT_EXTRACTION_JOB_NAME }

func (j *DocumentExtractionJob) Handle(payload []byte) (err error) {
	var data documentInterfaces.DocumentExtractionJobPayload
	if unmarshalErr := json.Unmarshal(payload, &data); unmarshalErr != nil {
		return unmarshalErr
	}

	defer func() {
		if r := recover(); r != nil {
			j.documentExtractionService.MarkFailed(context.Background(), data.DocumentId)
			err = fmt.Errorf("%v", r)
		}
	}()

	j.documentExtractionService.Extract(context.Background(), data.DocumentId, data.FileBuffer)
	return nil
}
