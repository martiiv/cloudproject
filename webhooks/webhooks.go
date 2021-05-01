package webhooks

import (
	"cloudproject/extra"
	"net/http"
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
	ID             string //ID from RouteInformation DB entry
	TrafficMessage string //Current traffic messages on route from api
	WeatherMessage string //Current weather conditions from weather api

	extra.RouteInformation //DB entry
}

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
