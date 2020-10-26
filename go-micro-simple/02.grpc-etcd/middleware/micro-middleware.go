package middleware

import "github.com/gin-gonic/gin"

func MicroMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})

		c.Next()
	}
}
