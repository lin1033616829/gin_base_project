package middleware

import (
	"github.com/gin-gonic/gin"

)

// RateMiddleware 限制访问
func RateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
