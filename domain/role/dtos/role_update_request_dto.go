package dtos

import (
	"github.com/gin-gonic/gin"
)

type RoleUpdateRequestDto struct {
	Name *string `json:"name"`
}

func AssignRoleUpdateRequestDto(c *gin.Context) *RoleUpdateRequestDto {
	dto := &RoleUpdateRequestDto{}
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return dto
}
