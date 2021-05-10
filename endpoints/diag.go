package endpoints

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Uptime of the program
var Uptime time.Time

// APIs
// https://openrouteservice.org/
// https://developer.tomtom.com/
// https://openweathermap.org/api
// https://positionstack.com/
// https://developer.mapquest.com

// Diag shows diagnostics interface
func Diag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var positionstackStatusCode int
	var tomtomStatusCode int
	var openRouteServiceStatusCode int
	var openWeatherMapStatusCode int
	var mapQuestStatusCode int

	// Does a request to the posistionstack API.
	respPositionStack, err := http.Get("https://api.positionstack.com")
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the positionstack API, %v", err)
		positionstackStatusCode = http.StatusInternalServerError
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
		tomtomStatusCode = http.StatusInternalServerError
	} else {
		tomtomStatusCode = respTomTom.StatusCode
		defer respTomTom.Body.Close()
	}

	// Does a request to the OpenRouteService API.
	respOpenRouteService, err := http.Get("https://openrouteservice.org")
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the OpenRouteService API, %v", err)
		openRouteServiceStatusCode = http.StatusInternalServerError
	} else {
		openRouteServiceStatusCode = respOpenRouteService.StatusCode
		defer respOpenRouteService.Body.Close()
	}

	// Does a request to the OpenWeatherMap API.
	respOpenWeatherMap, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Gjovik&appid=92721f2c7ecab4f083189daef6b7f146")
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the OpenWeatherMap API, %v", err)
		openWeatherMapStatusCode = http.StatusInternalServerError
	} else {
		openWeatherMapStatusCode = respOpenWeatherMap.StatusCode
		defer respOpenWeatherMap.Body.Close()
	}

	// Does a request to the MapQuest API.
	respMapQuest, err := http.Get("http://open.mapquestapi.com/geocoding/v1/address?key=UvCctIMBPNYcpfiAkTCkVjakeCjoPpPR&location=1600+Pennsylvania+Ave+NW,Washington,DC,20500")
	// If any errors occur, log it and set the status code to 500,
	// otherwise set the status code to the received status code
	if err != nil {
		log.Printf("Something went wrong with the MapQuest API, %v", err)
		mapQuestStatusCode = http.StatusInternalServerError
	} else {
		mapQuestStatusCode = respMapQuest.StatusCode
		defer respMapQuest.Body.Close()
	}

	fmt.Fprintf(w, `{"positionstack": "%v", "tomtom": "%v", "openrouteservice": "%v", "openweathermap": "%v", "mapquest": "%v", "version": "v1", "uptime": %v}`,
		positionstackStatusCode, tomtomStatusCode, openRouteServiceStatusCode, openWeatherMapStatusCode, mapQuestStatusCode, int(time.Since(Uptime)/time.Second))
}
