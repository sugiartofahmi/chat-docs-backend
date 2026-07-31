package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "go-service/infrastructure/dtos"
)

type ConversationQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignConversationQueryRequestDto(c *gin.Context) *ConversationQueryRequestDto {
	q := &ConversationQueryRequestDto{}
	if err := c.ShouldBindQuery(q); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return q
}
