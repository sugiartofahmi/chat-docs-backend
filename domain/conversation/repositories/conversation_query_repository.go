package repositories

import (
	"context"
	"log"

	conversationDtos "go-service/domain/conversation/dtos"
	conversationInterfaces "go-service/domain/conversation/interfaces"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"
	"go-service/infrastructure/exceptions"
	"go-service/infrastructure/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConversationQueryRepository struct {
	db                *gorm.DB
	conversationModel *gorm.DB
}

func NewConversationQueryRepository(db *gorm.DB) conversationInterfaces.ConversationQueryRepositoryInterface {
	return &ConversationQueryRepository{
		db:                db,
		conversationModel: db.Model(&entities.ConversationEntity{}),
	}
}

func (r *ConversationQueryRepository) Pagination(ctx context.Context, projectId uuid.UUID, dto *conversationDtos.ConversationQueryRequestDto) *infradtos.PaginationResultDto[entities.ConversationEntity] {
	var result infradtos.PaginationResultDto[entities.ConversationEntity]

	query := r.conversationModel.WithContext(ctx).Where("project_id = ?", projectId)

	r.querySearch(&query, dto)
	r.querySort(&query, dto)

	if err := query.Count(&result.Count).Error; err != nil {
		log.Println("Error count conversations:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	if err := query.Scopes(utils.Paginate(&dto.PaginationQueryRequestDto)).Find(&result.Data).Error; err != nil {
		log.Println("Error paginate conversations:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (r *ConversationQueryRepository) querySearch(query **gorm.DB, dto *conversationDtos.ConversationQueryRequestDto) {
	if dto.Search != "" {
		*query = (*query).Where("content ILIKE ?", "%"+dto.Search+"%")
	}
}

func (r *ConversationQueryRepository) querySort(query **gorm.DB, dto *conversationDtos.ConversationQueryRequestDto) {
	sortBy := "created_at"
	sortableColumns := []string{"created_at", "role"}

	if dto.SortBy != "" && utils.Contains(sortableColumns, dto.SortBy) {
		sortBy = dto.SortBy
	}

	order := "asc"
	if dto.Order != "" {
		order = string(dto.Order)
	}

	*query = (*query).Order(sortBy + " " + order)
}
