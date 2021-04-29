package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiKey = "a92721f2c7ecab4f083189daef6b7f146"

func CurrentWeather(rw http.ResponseWriter, request *http.Request, latitude string, longitude string) {
	rw.Header().Set("Content-type", "application/json")

	url := ""

	if latitude != "" && longitude != "" {
		url = "api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + apiKey
	} else {
		fmt.Fprint(rw, "Check formatting of lat and lon")
	}
}

func currentWeatherHandler(rw http.ResponseWriter, url string) {
	var weatherData allWeatherData

	// Uses request URL
	resp, err := http.Get(url)
	if err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Reads the data from the resp.Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshalling the body into the weatherData struct/fields
	if err := json.Unmarshal(body, &weatherData); err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		// Displays the data to the user
		mes, err := fmt.Fprintf(rw, `{
				"country": "%v",
				"continent": "%v",
				"scope": "%v",
				"confirmed": "%v",
				"recovered": "%v",
				"population_percentage": "%v"
			}`, jsonCountry, jsonContinent, scope, jsonConfirmed, jsonRecovered, math.Round(jsonPopulationPercentage*100)/100)
		if err == nil {
			fmt.Print(mes)
		}
	*/
}
