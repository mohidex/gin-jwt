package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/identity-service/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.ValidateJWT(c); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
