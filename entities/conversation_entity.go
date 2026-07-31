package entities

import (
	"time"

	"github.com/google/uuid"
)

type ConversationEntity struct {
	Id        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProjectId uuid.UUID  `gorm:"type:uuid;not null;index"                         json:"project_id"`
	Role      string     `gorm:"size:20;not null"                                 json:"role"`
	Content   string     `gorm:"type:text;not null"                               json:"content"`
	CreatedAt *time.Time `gorm:"autoCreateTime"                                   json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"                                   json:"updated_at"`

	Project ProjectEntity `gorm:"foreignKey:ProjectId;constraint:OnDelete:CASCADE" json:"project"`
}

func (ConversationEntity) TableName() string { return "conversations" }
