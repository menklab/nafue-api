package rest

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	r *gin.Engine
)

func Init() {

	r = gin.Default()

	// CORS Requests
	new(CORSController).Init(r)

	//rest API controllers
	new(HealthyController).Init(r)
	new(LetsEncryptController).Init(r)
	new(FileController).Init(r)
}

func Listen(uri string) {
	err := r.Run(uri)
	log.Println(err.Error())
}
