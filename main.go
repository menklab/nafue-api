package main

import (
	"os"
	"sparticus/controllers/rest"
	"sparticus/database"
	"sparticus/repositories"
	"sparticus/services"
)

func main() {

	repositories.Init(database.Database())
	//	database.Migrate();
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
