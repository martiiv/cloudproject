package database

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// APIs
// https://apilayer.com
// https://developer.tomtom.com/
// https://openrouteservice.org/

// Diag shows diagnostics interface
func Diag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var positionstackStatusCode int
	var tomtomStatusCode int
	var openRouteServiceStatusCode int
	var startTime = time.Now()

	// Does a request to the posistionstack API.
	respPositionStack, err := http.Get("https://api.positionstack.com")
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the mMedia API, %v", err)
		positionstackStatusCode = 500
	} else {
		positionstackStatusCode = respPositionStack.StatusCode
		defer respPositionStack.Body.Close()
	}

	// Does a request to the TomTom API.
	respTomTom, err := http.Get("https://api.tomtom.com/search/2/search/pizza.json?key=gOorFpmbH5GPKh6uGqcfJN76oKFKfswA&lat=37.8085&lon=-122.4239")
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the TomTom API, %v", err)
		tomtomStatusCode = 500
	} else {
		tomtomStatusCode = respTomTom.StatusCode
		defer respTomTom.Body.Close()
	}

	// Does a request to the TomTom API.
	respOpenRouteService, err := http.Get("https://openroute......") // Must be fixed
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the OpenRouteService API, %v", err)
		openRouteServiceStatusCode = 500
	} else {
		openRouteServiceStatusCode = respOpenRouteService.StatusCode
		defer respOpenRouteService.Body.Close()
	}

	fmt.Fprintf(w, `{"positionstack": "%v", "tomtom": "%v", "openrouteservice": "%v", "version": "v1", "uptime": %v}`,
		positionstackStatusCode, tomtomStatusCode, openRouteServiceStatusCode, int(time.Since(startTime)/time.Second))
}
