package dtos

import (
	"github.com/gin-gonic/gin"
)

type AuthLoginRequestDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func AssignAuthLoginRequestDto(c *gin.Context) *AuthLoginRequestDto {
	dto := &AuthLoginRequestDto{}
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(gin.Error{Err: err, Type: gin.ErrorTypeBind})
	}
	return dto
}
