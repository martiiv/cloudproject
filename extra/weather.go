package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var apiKey = "92721f2c7ecab4f083189daef6b7f146"

func CurrentWeather(rw http.ResponseWriter, request *http.Request /*, latitude string, longitude string*/) {
	rw.Header().Set("Content-type", "application/json")

	address := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers

	latitude, longitude, err := GetLocation(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	url := ""

	if latitude != "" && longitude != "" {
		url = "https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + apiKey
	} else {
		fmt.Fprint(rw, "Check formatting of lat and lon")
	}
	currentWeatherHandler(rw, url)
}

func currentWeatherHandler(rw http.ResponseWriter, url string) {
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

	var weather weatherData

	// Unmarshalling the body into the weatherData struct/fields
	if err := json.Unmarshal(body, &weather); err != nil {
		http.Error(rw, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var data []outputWeather

	main := weather.Weather[0].Main
	tempActual := weather.Main.Temp
	tempFeelsLike := weather.Main.FeelsLike
	tempMin := weather.Main.TempMin
	tempMax := weather.Main.TempMax
	humidity := weather.Main.Humidity
	visibility := weather.Visibility
	windSpeed := weather.Wind.Speed
	windDeg := weather.Wind.Deg
	sunrise := weather.Sys.Sunrise
	sunset := weather.Sys.Sunset

	jsonStruct := outputWeather{Main: main, Temp: tempActual, FeelsLike: tempFeelsLike, TempMin: tempMin, TempMax: tempMax,
		Humidity: humidity, Visibility: visibility, WindSpeed: windSpeed, WindDeg: windDeg, Sunrise: sunrise, Sunset: sunset}

	data = append(data, jsonStruct)

	output, err := json.Marshal(data) //Marshalling the array to JSON
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(rw, "%v", string(output)) //Outputs the chargers
}
