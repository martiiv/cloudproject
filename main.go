package main

import (
	endpoint "cloudproject/endpoints"
	extra "cloudproject/extra"
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
	http.HandleFunc("/weather/", extra.CurrentWeather)
	http.HandleFunc("/diag", endpoint.Diag)
	http.HandleFunc("/charge/", endpoint.EVStations)
	http.HandleFunc("/petrol/", endpoint.PetrolStation)
	http.HandleFunc("/messages/", endpoint.Messages)

	log.Println(http.ListenAndServe(getPort(), nil))
}
