package webhooks

import (
	"cloudproject/endpoints"
	"cloudproject/extra"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/**
 * Class webhooks.go
 * Will contain all webhooks functionality for the application (may get separated into more files)
 * Will contain the following funcitons:
 *										Handler
 *
 * @author Martin Iversen
 * @date 01.05.2021
 * @version 0.2
 */

// Struct WebHook will be used for storing information
type WebHook struct {
	ID             	string //ID from RouteInformation DB entry
	URL            	string //Webhook URL to be invoked
	TrafficMessage 	string //Current traffic messages on route from api
	WeatherMessage 	string //Current weather conditions from weather api

	route			extra.RouteInformation
}
var webHookInit []WebHook
var weatherApi = "92721f2c7ecab4f083189daef6b7f146"


/**
 * Function Handler
 * Will handle all the requests sent to the webhook endpoint
 * MethodPost:
 * MethodGet:
 * MethodPut:
 * MethodDelete:
 */
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		newWebhook := WebHook{}
		err:= json.NewDecoder(r.Body).Decode(&newWebhook)
		if err != nil {
			http.Error(w, "Unable to decode POST Request"+ err.Error(), http.StatusBadRequest)
		}
		webHookInit = append(webHookInit,newWebhook)
		newWebhook = CreateWebhook(w,r webHookInit)


	case http.MethodGet:

	case http.MethodPut:

	case http.MethodDelete:

	}
}

/**
 * Function Response
 * Will format the Json Response for the user
 */
func Response(webhook WebHook) {

}

/**
 * Function Check
 * Will check for updates in weather conditions and traffic incidents
 */
func Check() {

}

func CreateWebhook(w http.ResponseWriter, r*http.Request, route extra.RouteInformation, hook WebHook) WebHook{


	startPoint := route.StartDestination
	latitude, longitude, err := extra.GetLocation(url.QueryEscape(startPoint)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, "Error! Couldnt get latitude and longitude from api"+err.Error(), http.StatusBadRequest)
	}


	//Define the current trafficMessage for the route
	weatherUrl := "https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + weatherApi
	hook.WeatherMessage = weatherUrl //TODO Change this to currentweather
	// TODO Get current weather string fra Tormod weather := extra.CurrentWeather(w, r, weatherUrl)



}
