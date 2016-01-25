package main

import (
	"os"
	"nafue/controllers/rest"
	"nafue/database"
	"nafue/repositories"
	"nafue/services"
)

func main() {

	repositories.Init(database.Database())
		database.Migrate();
	services.Init()

	//Initialize Server
	rest.Init()

	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	// Start Server
	rest.Listen(":" + port)
}
