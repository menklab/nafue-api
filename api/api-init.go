package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/menkveldj/nafue-api/api/middleware"
	"github.com/menkveldj/nafue-api/config"
	"github.com/menkveldj/nafue-api/api/controllers"
)

var (
	r *gin.Engine
)

func Init() {

	// start gin
	r = gin.Default()

	// Setup Middleware
	r.Use(middleware.CORS())

	// setup routes
	routes := &config.Routes{
		Public: r.Group("/api"),
	}


	//rest API controllers
	new(controllers.HealthyController).Init(routes)
	new(controllers.FileController).Init(routes)
	new(controllers.PaymentController).Init(routes)
	new(controllers.BasicAnalyticsController).Init(routes)
}

func Listen(uri string) {
	err := r.Run(uri)
	log.Println(err.Error())
}
