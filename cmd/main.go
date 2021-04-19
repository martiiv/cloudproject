package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("Listening on port: " + getPort())

	handlers() //HTTP handlers

}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func handlers() {
	//Todo Add endpoints
	http.HandleFunc("/diag", Diag)
	log.Println(http.ListenAndServe(getPort(), nil))
}
