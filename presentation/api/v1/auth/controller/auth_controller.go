package controller

import (
	"net/http"

	"go-service/domain/auth/constants"
	"go-service/domain/auth/dtos"
	authInterfaces "go-service/domain/auth/interfaces"
	"go-service/infrastructure/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService authInterfaces.AuthServiceInterface
}

func NewAuthController(router *gin.Engine, authService authInterfaces.AuthServiceInterface) {
	controller := &AuthController{authService: authService}

	authRoute := router.Group("/api/v1/auth")
	authRoute.POST("/login", controller.Login())
	authRoute.POST("/register", controller.Register())
}

func (c *AuthController) Login() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := dtos.AssignAuthLoginRequestDto(httpContext)
		result := c.authService.Login(ctx, dto)
		response := utils.SuccessResponse(http.StatusOK, constants.AUTH_LOGIN_SUCCESS, result)
		httpContext.JSON(http.StatusOK, response)
	}
}

func (c *AuthController) Register() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		ctx := httpContext.Request.Context()
		dto := dtos.AssignAuthRegisterRequestDto(httpContext)
		result := c.authService.Register(ctx, dto)
		response := utils.SuccessResponse(http.StatusCreated, constants.AUTH_REGISTER_SUCCESS, result)
		httpContext.JSON(http.StatusCreated, response)
	}
}
