package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"nafue-api/config"
)

type CORSMiddleware struct{}

func (self *CORSMiddleware) Init(r *gin.Engine) {
	r.Use(self.corsMiddleware)
}

func (self *CORSMiddleware) corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", config.CorsHost)
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		log.Println("OPTIONS")
		c.AbortWithStatus(204)
	} else {
		c.Next()
	}
}
