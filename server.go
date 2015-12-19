package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"os"
	"net/http"
	"sparticus/controllers/rest"
)

func main() {

	router := httprouter.New()

	//Initialize rest controlleØrs
	rest.Init(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("File Services %s", port)
	err := http.ListenAndServe(":"+port, router)
	log.Println(err.Error())
}
