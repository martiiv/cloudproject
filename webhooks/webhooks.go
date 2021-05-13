package webhooks

import (
	"cloudproject/database"
	"cloudproject/endpoints"
	"cloudproject/structs"
	"cloudproject/utils"
	"encoding/json"
	"fmt"
	_ "fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"sync"
	"time"
	_ "time"
)

/**
 * Function Check
 * Will check for updates in weather conditions and traffic incidents
 */
func Check(w http.ResponseWriter) {
	iter := database.Client.Collection(database.Collection).Documents(database.Ctx) // Loop through all entries in collection "messages"
	var hook structs.Webhook

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err := doc.DataTo(&hook); err != nil {
			return
		}

		weatherMessage := hook.Weather

		latitude, longitude, err := database.LocationPresent(url2.QueryEscape(hook.DepartureLocation)) //Receives the latitude and longitude of the place passed in the url
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		url := ""
		if latitude != "" && longitude != "" {
			// Defines the url to the openweathermap API with relevant latitude and longitude and apiKey
			url = "https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + utils.OpenweathermapKey
		} else {
			fmt.Fprint(w, "Check formatting of lat and lon")
		}

		newMessage := endpoints.CurrentWeatherHandler(w, url).Main.Message
		if !(newMessage == weatherMessage) {
			hook.Weather = newMessage
			database.Update(doc.Ref.ID, hook)
			fmt.Fprintf(w, "WeatherMessage update new registered weather for:"+hook.DepartureLocation+" is:"+hook.Weather)
		}

	}

	time.Sleep(time.Minute * 30)
	Check(w)
}

func CreateWebhook(w http.ResponseWriter, r *http.Request) {
	AddWebhook(w, r)

}

func AddWebhook(w http.ResponseWriter, r *http.Request) (structs.Webhook, string) {

	wg := new(sync.WaitGroup)

	if r.Method != http.MethodPost {
		http.Error(w, "Expected POST method", http.StatusMethodNotAllowed)
	}

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(input) == 0 {
		http.Error(w, "Your message appears to be empty", http.StatusBadRequest)
	}

	var notification structs.Webhook
	if err = json.Unmarshal(input, &notification); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	message, ok := webhookFormat(notification)
	if !ok {
		http.Error(w, message, http.StatusNoContent)
	}

	id, _, err := database.Client.Collection(database.Collection).Add(database.Ctx,
		map[string]interface{}{
			"url":                notification.Url,
			"ArrivalDestination": notification.ArrivalDestination,
			"ArrivalTime":        notification.ArrivalTime,
			"Weather":            notification.Weather,
			"DepartureLocation":  notification.DepartureLocation,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return structs.Webhook{}, ""
	} else {
		http.Error(w, "Registered with ID: "+id.ID, http.StatusCreated)
		go Check(w)
		wg.Wait()
		CalculateDeparture(id.ID)
		go SendNotification(id.ID)
	}

	return notification, id.ID
}

func webhookFormat(web structs.Webhook) (string, bool) {

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
	iter := database.Client.Collection(database.Collection).Documents(database.Ctx) // Loop through all entries in collection "messages"

	var firebase structs.Webhook

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
			log.Fatalf(err.Error())
		}

		if arrival.Before(time.Now().AddDate(0, 0, -1)) {
			err := database.Delete(doc.Ref.ID)
			if err != nil {
				//Todo Error handling
			}
			fmt.Println("Webhook deleted")
		}
	}
	time.Sleep(time.Hour * 24)
	DeleteExpiredWebhooks()
}
