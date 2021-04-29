package main

import (
	endpoint "cloudproject/endpoints"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	fmt.Println(url.QueryEscape("10.430053,60.79012,10.699832,61.116501"))

	log.Println("Listening on port: " + getPort())
	handlers()
}

func handlers() {
	http.HandleFunc("/charge/", endpoint.EVStations)
	http.HandleFunc("/petrol/", endpoint.PetrolStation)
	http.HandleFunc("/messages/", endpoint.Messages)

	log.Println(http.ListenAndServe(getPort(), nil))
}
