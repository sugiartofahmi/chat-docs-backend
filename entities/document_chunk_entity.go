package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
)

type DocumentChunkEntity struct {
	Id         uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	DocumentId uuid.UUID        `gorm:"type:uuid;not null;index"                         json:"document_id"`
	ChunkIndex int              `gorm:"not null"                                         json:"chunk_index"`
	Content    string           `gorm:"type:text;not null"                               json:"content"`
	Embedding  *pgvector.Vector `gorm:"type:vector(1536)"                                json:"-"`
	CreatedAt  *time.Time       `gorm:"autoCreateTime"                                   json:"created_at"`
	UpdatedAt  *time.Time       `gorm:"autoUpdateTime"                                   json:"updated_at"`

	Document DocumentEntity `gorm:"foreignKey:DocumentId;constraint:OnDelete:CASCADE" json:"document"`
}

func (DocumentChunkEntity) TableName() string { return "document_chunks" }
