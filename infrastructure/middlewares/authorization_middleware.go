package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"go-service/infrastructure/exceptions"
	"go-service/infrastructure/utils"
)

const AuthUserKey = "user"

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			panic(*exceptions.UnauthenticatedException("missing or invalid authorization header"))
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims := utils.ValidateToken(token)
		c.Set(AuthUserKey, claims.User)
		c.Next()
	}
}
