package main

import (
	endpoint "cloudproject/endpoints"
	"log"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {

	// Starts uptime of program
	endpoint.Uptime = time.Now()

	log.Println("Listening on port: " + getPort())
	handlers()
}

func handlers() {
	http.HandleFunc("/diag", endpoint.Diag)
	http.HandleFunc("/charge/", endpoint.EVStations)
	log.Println(http.ListenAndServe(getPort(), nil))
}
