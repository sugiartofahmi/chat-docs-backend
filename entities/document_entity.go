package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentEntity struct {
	Id           uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProjectId    uuid.UUID             `gorm:"type:uuid;not null;index"                         json:"project_id"`
	Filename     string                `gorm:"size:255;not null"                                json:"filename"`
	Status       string                `gorm:"size:50;not null;default:'processing'"            json:"status"`
	ErrorMessage *string               `gorm:"type:text"                                        json:"error_message,omitempty"`
	Chunks       []DocumentChunkEntity `gorm:"foreignKey:DocumentId"                            json:"chunks,omitempty"`
	CreatedAt    *time.Time            `gorm:"autoCreateTime"                                   json:"created_at"`
	UpdatedAt    *time.Time            `gorm:"autoUpdateTime"                                   json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `gorm:"index"                                            json:"-"`

	Project ProjectEntity `gorm:"foreignKey:ProjectId;constraint:OnDelete:CASCADE" json:"project"`
}

func (DocumentEntity) TableName() string { return "documents" }
