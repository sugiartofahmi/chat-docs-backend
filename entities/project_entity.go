package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectEntity struct {
	Id        uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string           `gorm:"size:255;not null"                                json:"name"`
	Documents []DocumentEntity `gorm:"foreignKey:ProjectId"                             json:"documents,omitempty"`
	CreatedAt *time.Time       `gorm:"autoCreateTime"                                   json:"created_at"`
	UpdatedAt *time.Time       `gorm:"autoUpdateTime"                                   json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index"                                            json:"-"`
}

func (ProjectEntity) TableName() string { return "projects" }
