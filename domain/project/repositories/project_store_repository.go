package repositories

import (
	"context"
	"log"

	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"

	"gorm.io/gorm"
)

type ProjectStoreRepository struct {
	db           *gorm.DB
	projectModel *gorm.DB
}

func NewProjectStoreRepository(db *gorm.DB) projectInterfaces.ProjectStoreRepositoryInterface {
	return &ProjectStoreRepository{
		db:           db,
		projectModel: db.Model(&entities.ProjectEntity{}),
	}
}

func (repo *ProjectStoreRepository) WithTransaction(tx *gorm.DB) projectInterfaces.ProjectStoreRepositoryInterface {
	return &ProjectStoreRepository{
		db:           tx,
		projectModel: tx.Model(&entities.ProjectEntity{}),
	}
}

func (repo *ProjectStoreRepository) Create(ctx context.Context, entity *entities.ProjectEntity) *entities.ProjectEntity {
	query := repo.projectModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create project:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}
