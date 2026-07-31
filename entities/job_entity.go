package entities

import (
	"time"

	"github.com/google/uuid"
)

type JobEntity struct {
	Id         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	JobName    string     `gorm:"size:255;not null;index"                          json:"job_name"`
	Source     string     `gorm:"size:20;not null"                                 json:"source"`
	Payload    string     `gorm:"type:text;not null"                               json:"payload"`
	Status     string     `gorm:"size:20;not null;default:'pending'"               json:"status"`
	Error      *string    `gorm:"type:text"                                        json:"error,omitempty"`
	StartedAt  *time.Time `                                                        json:"started_at,omitempty"`
	FinishedAt *time.Time `                                                        json:"finished_at,omitempty"`
	CreatedAt  *time.Time `gorm:"autoCreateTime"                                   json:"created_at"`
}

func (JobEntity) TableName() string { return "jobs" }
