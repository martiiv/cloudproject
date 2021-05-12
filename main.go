package main

import (
	"cloudproject/endpoints"
	"cloudproject/webhooks"
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
	println(endpoints.GetMessageWeight("violent rain"))

	webhooks.Init()

	// Starts uptime of program
	endpoints.Uptime = time.Now()

	log.Println("Listening on port: " + getPort())
	handlers()

	defer webhooks.Client.Close()
}

func handlers() {
	http.HandleFunc("/weather/", endpoints.CurrentWeather)
	http.HandleFunc("/poi/", endpoints.PointOfInterest)
	http.HandleFunc("/diag", endpoints.Diag)
	http.HandleFunc("/charge/", endpoints.EVStations)
	http.HandleFunc("/petrol/", endpoints.PetrolStation)
	http.HandleFunc("/messages/", endpoints.Messages)
	http.HandleFunc("/route/", endpoints.Route)
	http.HandleFunc("/hook/", webhooks.CreateWebhook)

	log.Println(http.ListenAndServe(getPort(), nil))
}
