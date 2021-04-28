package main

import (
	endpoint "cloudproject/endpoints"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {

	log.Println("Listening on port: " + getPort())
	handlers()
}

func handlers() {
	http.HandleFunc("/charge/", endpoint.EVStations)
	log.Println(http.ListenAndServe(getPort(), nil))
}
