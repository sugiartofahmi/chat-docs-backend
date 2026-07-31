package repositories

import (
	"context"
	"log"

	projectDtos "go-service/domain/project/dtos"
	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/entities"
	infradtos "go-service/infrastructure/dtos"
	"go-service/infrastructure/exceptions"
	"go-service/infrastructure/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectQueryRepository struct {
	db           *gorm.DB
	projectModel *gorm.DB
}

func NewProjectQueryRepository(db *gorm.DB) projectInterfaces.ProjectQueryRepositoryInterface {
	return &ProjectQueryRepository{
		db:           db,
		projectModel: db.Model(&entities.ProjectEntity{}),
	}
}

func (r *ProjectQueryRepository) Pagination(ctx context.Context, dto *projectDtos.ProjectQueryRequestDto) *infradtos.PaginationResultDto[entities.ProjectEntity] {
	var result infradtos.PaginationResultDto[entities.ProjectEntity]

	query := r.projectModel.WithContext(ctx)

	r.querySearch(&query, dto)
	r.querySort(&query, dto)

	err := query.Count(&result.Count).Error
	if err != nil {
		log.Println("Error count projects:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	err = query.Scopes(utils.Paginate(&dto.PaginationQueryRequestDto)).Find(&result.Data).Error
	if err != nil {
		log.Println("Error paginate projects:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (r *ProjectQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.ProjectEntity {
	var result entities.ProjectEntity
	err := r.projectModel.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find project by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}
	return &result
}

func (r *ProjectQueryRepository) IsExistByName(ctx context.Context, name string) bool {
	var count int64
	err := r.projectModel.WithContext(ctx).Where("name = ?", name).Count(&count).Error
	if err != nil {
		log.Println("Error count project by name:", err)
		panic(*exceptions.ServerErrorException(err))
	}
	return count > 0
}

func (r *ProjectQueryRepository) IsExistById(ctx context.Context, id uuid.UUID) bool {
	var count int64
	err := r.projectModel.WithContext(ctx).Where("id = ?", id).Count(&count).Error
	if err != nil {
		log.Println("Error count project by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}
	return count > 0
}

func (r *ProjectQueryRepository) querySearch(query **gorm.DB, dto *projectDtos.ProjectQueryRequestDto) {
	if dto.Search != "" {
		*query = (*query).Where("name ILIKE ?", "%"+dto.Search+"%")
	}
}

func (r *ProjectQueryRepository) querySort(query **gorm.DB, dto *projectDtos.ProjectQueryRequestDto) {
	sortBy := "created_at"
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if dto.SortBy != "" {
		isColumnAllowed := utils.Contains(sortableColumns, dto.SortBy)
		if isColumnAllowed {
			sortBy = dto.SortBy
		}
	}

	order := "desc"
	if dto.Order != "" {
		order = string(dto.Order)
	}

	*query = (*query).Order(sortBy + " " + order)
}
