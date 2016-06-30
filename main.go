package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/menkveldj/nafue-api/config"
	"github.com/menkveldj/nafue-api/database"
	"github.com/menkveldj/nafue-api/repositories"
	"github.com/menkveldj/nafue-api/services"
	"github.com/menkveldj/nafue-api/api"
)

func main() {

	repositories.Init(database.Database())
	database.Migrate()
	services.Init()

	//Initialize Server
	api.Init()

	port := config.Port

	if port == "" {
		port = "8080"
	}

	// Start Server
	api.Listen(":" + port)
}
