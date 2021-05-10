package webhooks

import (
	"cloudproject/extra"
	"encoding/json"
	"fmt"
	_ "fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	_ "time"
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
	URL            string //Webhook URL to be invoked
	TrafficMessage string //Current traffic messages on route from api
	WeatherMessage string //Current weather conditions from weather api

	route extra.RouteInformation
}

var webHookInit []WebHook
var Collection = "message"

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
		err := json.NewDecoder(r.Body).Decode(&newWebhook)
		if err != nil {
			http.Error(w, "Unable to decode POST Request"+err.Error(), http.StatusBadRequest)
		}
		webHookInit = append(webHookInit, newWebhook)
		//newWebhook = CreateWebhook(w,r webHookInit)

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

func CreateWebhook(w http.ResponseWriter, r *http.Request, route extra.RouteInformation, hook WebHook) {

	startPoint := route.StartDestination
	latitude, longitude, err := extra.GetLocation(url.QueryEscape(startPoint)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, "Error! Couldnt get latitude and longitude from api"+err.Error(), http.StatusBadRequest)
	}

	//Define the current trafficMessage for the route
	weatherUrl := "https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + extra.OpenweathermapKey
	hook.WeatherMessage = weatherUrl //TODO Change this to currentweather
	// TODO Get current weather string fra Tormod weather := extra.CurrentWeather(w, r, weatherUrl)

}

func AddWebhook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Expected POST method", http.StatusMethodNotAllowed)
		return
	}

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if len(input) == 0 {
		http.Error(w, "Your message appears to be empty", http.StatusBadRequest)
		return
	}

	var notification extra.Webhook
	if err = json.Unmarshal(input, &notification); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	message, ok := webhookFormat(notification)
	if !ok {
		http.Error(w, message, http.StatusNoContent)
		return
	}

	id, _, err := Client.Collection(Collection).Add(Ctx,
		map[string]interface{}{
			"ArrivalDestination": notification.ArrivalDestination,
			"ArrivalTime":        notification.ArrivalTime,
			"DepartureLocation":  notification.DepartureLocation,
			"Repeat":             notification.Repeat,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		http.Error(w, "Registered with ID: "+id.ID, http.StatusCreated)
		CalculateDeparture(id.ID)
	}

}

func webhookFormat(web extra.Webhook) (string, bool) {

	if web.DepartureLocation == "" {
		return "Departure location cannot be empty", false
	} else if web.ArrivalDestination == "" {
		return "Arrival destination cannot be empty", false
	} else if web.ArrivalTime == "" {
		return "Arrival time cannot be empty", false
	}

	/*time, err := time.Parse(time.RFC822, web.ArrivalTime )
	if err != nil {
		return "Time format is not valid. Supported format is DD:MM:YY HH:mm", false
	}*/

	return "", true
}

func DeleteExpiredWebhooks() {
	iter := Client.Collection(Collection).Documents(Ctx) // Loop through all entries in collection "messages"

	var firebase extra.Webhook

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err := doc.DataTo(&firebase); err != nil {
			return
		}

		arrival, err := time.Parse(time.RFC822, firebase.ArrivalTime)

		if err != nil {
			//Todo Error handling
		}

		if arrival.After(time.Now().AddDate(0, 0, -1)) && firebase.Repeat == "" {
			err := Delete(doc.Ref.ID)
			if err != nil {
				//Todo Error handling
			}
			fmt.Println("Webhook deleted")

		}

	}

	time.Sleep(time.Hour * 24)
	DeleteExpiredWebhooks()

}
