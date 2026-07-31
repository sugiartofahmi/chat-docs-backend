package controller

import (
	"net/http"

	"go-service/domain/project/constants"
	"go-service/domain/project/dtos"
	projectInterfaces "go-service/domain/project/interfaces"
	"go-service/infrastructure/middlewares"
	"go-service/infrastructure/utils"
	"go-service/infrastructure/validators"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectService projectInterfaces.ProjectServiceInterface
}

func NewProjectController(router *gin.Engine, projectService projectInterfaces.ProjectServiceInterface) {
	controller := &ProjectController{projectService: projectService}

	projectRoute := router.Group("/api/v1/projects")
	projectRoute.GET("", controller.Pagination())
	projectRoute.GET("/:id", controller.Detail())
	projectRoute.POST("", middlewares.ValidateRequestJson[dtos.ProjectCreateRequestDto](), controller.Create())
}

func (c *ProjectController) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := dtos.AssignProjectQueryRequestDto(httpContext)
		result := c.projectService.Pagination(ctx, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := dtos.ProjectResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, constants.PROJECT_PAGINATION_SUCCESS, items, *meta)
		httpContext.JSON(http.StatusOK, response)
	}
}

func (c *ProjectController) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		id := validators.ValidateUUID(httpContext.Param("id"))
		result := dtos.ProjectResponseDtoFromEntity(c.projectService.Detail(ctx, id))
		response := utils.SuccessResponse(http.StatusOK, constants.PROJECT_DETAIL_SUCCESS, result)
		httpContext.JSON(http.StatusOK, response)
	}
}

func (c *ProjectController) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*dtos.ProjectCreateRequestDto)
		result := dtos.ProjectResponseDtoFromEntity(c.projectService.Create(ctx, dto))
		response := utils.SuccessResponse(http.StatusCreated, constants.PROJECT_CREATE_SUCCESS, result)
		httpContext.JSON(http.StatusCreated, response)
	}
}
