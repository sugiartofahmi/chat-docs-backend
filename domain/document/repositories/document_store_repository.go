package repositories

import (
	"context"
	"log"

	documentInterfaces "go-service/domain/document/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"

	"gorm.io/gorm"
)

type DocumentStoreRepository struct {
	db            *gorm.DB
	documentModel *gorm.DB
}

func NewDocumentStoreRepository(db *gorm.DB) documentInterfaces.DocumentStoreRepositoryInterface {
	return &DocumentStoreRepository{
		db:            db,
		documentModel: db.Model(&entities.DocumentEntity{}),
	}
}

func (repo *DocumentStoreRepository) WithTransaction(tx *gorm.DB) documentInterfaces.DocumentStoreRepositoryInterface {
	return &DocumentStoreRepository{
		db:            tx,
		documentModel: tx.Model(&entities.DocumentEntity{}),
	}
}

func (repo *DocumentStoreRepository) Create(ctx context.Context, entity *entities.DocumentEntity) *entities.DocumentEntity {
	query := repo.documentModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create document:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}
