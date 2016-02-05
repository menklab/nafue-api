package main

import (
	_ "github.com/joho/godotenv/autoload"
	"nafue/config"
	"nafue/controllers/rest"
	"nafue/database"
	"nafue/repositories"
	"nafue/services"
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