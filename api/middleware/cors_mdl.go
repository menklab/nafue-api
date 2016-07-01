package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/menkveldj/nafue-api/utility"
	"github.com/menkveldj/nafue-api/config"
)

func CORS() gin.HandlerFunc {
	return corsMiddleware
}

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", config.CorsHost)
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Token, X-Device-Token")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, X-Auth-Token, X-Device-Token")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		utility.Debug("OPTIONS")
		c.AbortWithStatus(204)
	} else {
		c.Next()
	}
}
