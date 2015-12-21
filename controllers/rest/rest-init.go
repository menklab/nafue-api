package rest

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	r *gin.Engine
)

func Init() {

	//	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	//rest API controllers
	new(HealthyController).Init(r)
	new(FileController).Init(r)
}

func Listen(uri string) {
	err := r.Run(uri)
	log.Println(err.Error())
}
