package config

import "github.com/gin-gonic/gin"

type Routes struct {
	//PreAuth *gin.RouterGroup
	//Auth    *gin.RouterGroup
	Public  *gin.RouterGroup
	//Admin   *gin.RouterGroup
	//Media   *gin.RouterGroup
}
