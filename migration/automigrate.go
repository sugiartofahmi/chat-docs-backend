package migration

import (
	"gorm.io/gorm"

	"go-service/entities"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS vector`).Error; err != nil {
		return err
	}

	return db.AutoMigrate(
		&entities.ProjectEntity{},
		&entities.DocumentEntity{},
		&entities.DocumentChunkEntity{},
		&entities.ConversationEntity{},
		&entities.JobEntity{},
	)
}
