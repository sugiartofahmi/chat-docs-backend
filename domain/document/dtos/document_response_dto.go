package dtos

import (
	"time"

	"go-service/entities"

	"github.com/google/uuid"
)

type DocumentResponseDto struct {
	Id        uuid.UUID  `json:"id"`
	ProjectId uuid.UUID  `json:"project_id"`
	Filename  string     `json:"filename"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
}

func DocumentResponseDtoFromEntity(document *entities.DocumentEntity) *DocumentResponseDto {
	return &DocumentResponseDto{
		Id:        document.Id,
		ProjectId: document.ProjectId,
		Filename:  document.Filename,
		Status:    document.Status,
		CreatedAt: document.CreatedAt,
	}
}
