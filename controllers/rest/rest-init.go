package rest

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	r *gin.Engine
)

func Init() {

	// start gin
	r = gin.Default()

	// Setup Middleware
	new(CORSMiddleware).Init(r)

	//rest API controllers
	new(HealthyController).Init(r)
	new(LetsEncryptController).Init(r)
	new(FileController).Init(r)
	new(PaymentController).Init(r)
}

func Listen(uri string) {
	err := r.Run(uri)
	log.Println(err.Error())
}
