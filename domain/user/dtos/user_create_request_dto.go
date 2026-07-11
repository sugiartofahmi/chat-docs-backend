package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserCreateRequestDto struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	RoleId   uuid.UUID `json:"role_id" binding:"required"`
}

func AssignUserCreateRequestDto(c *gin.Context) *UserCreateRequestDto {
	dto := &UserCreateRequestDto{}
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return dto
}
