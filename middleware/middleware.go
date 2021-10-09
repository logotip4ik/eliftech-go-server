package middleware

import (
	"home/work/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := token.IsTokenValid(c); err != nil {
			c.String(401, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}