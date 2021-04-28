package main

import (
	endpoint "cloudproject/endpoints"
	"fmt"
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
	fmt.Println(2)
	http.HandleFunc("/charge/", endpoint.EVStations)
	fmt.Println(3)
	log.Println(http.ListenAndServe(getPort(), nil))
	fmt.Println(4)
}
