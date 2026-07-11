package dtos

import (
	"github.com/gin-gonic/gin"
)

type AuthRegisterRequestDto struct {
	Name     string `json:"name" binding:"required,min=3,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func AssignAuthRegisterRequestDto(c *gin.Context) *AuthRegisterRequestDto {
	dto := &AuthRegisterRequestDto{}
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return dto
}
