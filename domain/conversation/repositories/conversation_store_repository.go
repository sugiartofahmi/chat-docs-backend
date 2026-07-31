package repositories

import (
	"context"
	"log"

	conversationInterfaces "go-service/domain/conversation/interfaces"
	"go-service/entities"
	"go-service/infrastructure/exceptions"

	"gorm.io/gorm"
)

type ConversationStoreRepository struct {
	db                *gorm.DB
	conversationModel *gorm.DB
}

func NewConversationStoreRepository(db *gorm.DB) conversationInterfaces.ConversationStoreRepositoryInterface {
	return &ConversationStoreRepository{
		db:                db,
		conversationModel: db.Model(&entities.ConversationEntity{}),
	}
}

func (repo *ConversationStoreRepository) Create(ctx context.Context, entity *entities.ConversationEntity) *entities.ConversationEntity {
	query := repo.conversationModel.WithContext(ctx)

	if err := query.Create(entity).Error; err != nil {
		log.Println("Error create conversation:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}
