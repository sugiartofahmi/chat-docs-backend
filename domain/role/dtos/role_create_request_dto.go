package dtos

import (
	"github.com/gin-gonic/gin"
)

type RoleCreateRequestDto struct {
	Name string `json:"name" binding:"required"`
}

func AssignRoleCreateRequestDto(c *gin.Context) *RoleCreateRequestDto {
	dto := &RoleCreateRequestDto{}
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return dto
}
