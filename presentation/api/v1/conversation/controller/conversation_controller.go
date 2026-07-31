package controller

import (
	"net/http"

	"go-service/domain/conversation/constants"
	"go-service/domain/conversation/dtos"
	conversationInterfaces "go-service/domain/conversation/interfaces"
	"go-service/infrastructure/middlewares"
	"go-service/infrastructure/utils"
	"go-service/infrastructure/validators"

	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	conversationService conversationInterfaces.ConversationServiceInterface
}

func NewConversationController(router *gin.Engine, conversationService conversationInterfaces.ConversationServiceInterface) {
	controller := &ConversationController{conversationService: conversationService}

	router.GET("/api/v1/projects/:id/conversations", controller.List())
	router.POST(
		"/api/v1/projects/:id/conversations",
		middlewares.ValidateRequestJson[dtos.ConversationCreateRequestDto](),
		controller.Ask(),
	)
}

func (c *ConversationController) List() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		projectId := validators.ValidateUUID(httpContext.Param("id"))
		dto := dtos.AssignConversationQueryRequestDto(httpContext)

		result := c.conversationService.Pagination(ctx, projectId, dto)
		meta := utils.PaginationMetaBuilder(dto.Page, dto.PerPage, int(result.Count))
		items := dtos.ConversationResponseDtoFromEntities(result.Data)
		response := utils.SuccessResponsePagination(http.StatusOK, constants.CONVERSATION_LIST_SUCCESS, items, *meta)
		httpContext.JSON(http.StatusOK, response)
	}
}

func (c *ConversationController) Ask() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		projectId := validators.ValidateUUID(httpContext.Param("id"))
		dto := httpContext.MustGet(middlewares.RequestBodyJsonKey).(*dtos.ConversationCreateRequestDto)

		reply := c.conversationService.Ask(ctx, projectId, dto.Content)
		result := dtos.ConversationResponseDtoFromEntity(reply)
		response := utils.SuccessResponse(http.StatusCreated, constants.CONVERSATION_SUCCESS, result)
		httpContext.JSON(http.StatusCreated, response)
	}
}
