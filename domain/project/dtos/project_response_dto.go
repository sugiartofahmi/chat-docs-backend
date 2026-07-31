package dtos

import (
	"time"

	"go-service/entities"

	"github.com/google/uuid"
)

type ProjectResponseDto struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ProjectResponseDtoFromEntity(project *entities.ProjectEntity) *ProjectResponseDto {
	return &ProjectResponseDto{
		Id:        project.Id,
		Name:      project.Name,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}
}

func ProjectResponseDtoFromEntities(projects []*entities.ProjectEntity) []ProjectResponseDto {
	responses := make([]ProjectResponseDto, len(projects))
	for i, project := range projects {
		responses[i] = *ProjectResponseDtoFromEntity(project)
	}
	return responses
}
