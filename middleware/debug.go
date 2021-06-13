package middleware

import (
	"PingLeMe-Backend/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DebugAPI 筛选用于Debug的API
func DebugAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !conf.SystemDebugFlag {
			c.JSON(http.StatusNotFound, nil)
			c.Abort()
		}
		c.Next()
	}
}
