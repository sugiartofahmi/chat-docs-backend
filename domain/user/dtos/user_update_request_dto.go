package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserUpdateRequestDto struct {
	Name     *string    `json:"name"`
	Password *string    `json:"password"`
	Status   *int       `json:"status"`
	RoleId   *uuid.UUID `json:"role_id"`
}

func AssignUserUpdateRequestDto(c *gin.Context) *UserUpdateRequestDto {
	dto := &UserUpdateRequestDto{}
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return dto
}
