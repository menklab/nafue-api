
package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/yvasiyarov/swagger/generator"
	"log"
)

var (
	r *gin.Engine
)

func Init() {
	params := generator.Params{
		ApiPackage:      "nafue",
		MainApiFile:     "nafue/controllers/rest/rest-init.go",
		OutputFormat:    "swagger",
		OutputSpec:      "docs",
		//ControllerClass: "(_ctrl)$",
		Ignore:          "",
	}

	// generate api docs
	err := generator.Run(params)
	if err != nil {
		log.Fatal(err.Error())
	}

	// start gin
	r = gin.Default()

	// Setup Middleware
	new(CORSMiddleware).Init(r)

	// Server docs
	r.Static("/docs", "./docs")


	//rest API controllers
	new(HealthyController).Init(r)
	new(LetsEncryptController).Init(r)
	new(FileController).Init(r)
}

func Listen(uri string) {
	err := r.Run(uri)
	log.Println(err.Error())
}
