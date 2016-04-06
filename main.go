package main

import (
	_ "nafue-api/Godeps/_workspace/src/github.com/joho/godotenv/autoload"
	"nafue-api/config"
	"nafue-api/controllers/rest"
	"nafue-api/database"
	"nafue-api/repositories"
	"nafue-api/services"
)

func main() {

	repositories.Init(database.Database())
	database.Migrate()
	services.Init()

	//Initialize Server
	rest.Init()

	port := config.Port

	if port == "" {
		port = "8080"
	}

	// Start Server
	rest.Listen(":" + port)
}
