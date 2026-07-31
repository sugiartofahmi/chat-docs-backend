package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "go-service/infrastructure/dtos"
)

type ProjectQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignProjectQueryRequestDto(c *gin.Context) *ProjectQueryRequestDto {
	q := &ProjectQueryRequestDto{}
	if err := c.ShouldBindQuery(q); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return q
}
