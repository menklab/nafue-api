package main

import (
	 _ "github.com/joho/godotenv/autoload"
	"nafue/controllers/rest"
	"nafue/database"
	"nafue/repositories"
	"nafue/services"

	"nafue/config"
)

func main() {

	repositories.Init(database.Database())
		database.Migrate();
	services.Init()

	//Initialize Server
	rest.Init()

	port := config.Port

	if  port == "" {
		port = "9090"
	}

	// Start Server
	rest.Listen(":" + port)
}
