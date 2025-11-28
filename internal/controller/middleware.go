package controller

import "github.com/gin-gonic/gin"

func SetUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", "foo")
		c.Next()
	}
}
