package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"sparticus/controllers/rest"
	"sparticus/database"
	"sparticus/repositories"
	"sparticus/services"
)

func main() {

	repositories.Init(database.Database())
	database.Migrate();
	services.Init()
	router := httprouter.New()

	//Initialize rest controllers
	rest.Init(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("File Services %s", port)
	err := http.ListenAndServe(":"+port, router)
	log.Println(err.Error())
}
