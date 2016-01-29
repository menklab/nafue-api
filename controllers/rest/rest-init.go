package rest

import (
	"github.com/gin-gonic/gin"
	yaag_gin "github.com/betacraft/yaag/gin"
	"github.com/betacraft/yaag/yaag"
	"log"
)

var (
	r *gin.Engine
)

func Init() {

	// start gin
	r = gin.Default()

	// generate api docs
	yaag.Init(&yaag.Config{On: true, DocTitle: "Nafue", DocPath: "docs/index.html", BaseUrls: map[string]string{"Production": "https://api.nafue.com", "Local Dev": "http://localhost:9090"}})
	r.Use(yaag_gin.Document())

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
