package webhooks

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"cloudproject/extra"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CalculateDeparture(id string) {

	information, _ := Client.Collection(Collection).Doc(id).Get(Ctx)

	var message extra.Webhook
	if err := information.DataTo(&message); err != nil {
		log.Println(err.Error())
	}

	location := message.DepartureLocation

	fmt.Println(location)

	startLat, startLong, err := extra.GetLocation(location)
	if err != nil {
		//Todo error handling
	}

	endLat, endLong, err := extra.GetLocation(message.ArrivalDestination)
	if err != nil {
		//Todo error handling
	}

	coordinates := startLat + "%2C" + startLong + "%3A" + endLat + "%2C" + endLong

	resp, err := http.Get("https://api.tomtom.com/routing/1/calculateRoute/" + coordinates + "/json?instructionsType=coded&traffic=false&avoid=unpavedRoads&travelMode=car&key=" + extra.TomtomKey)
	if err != nil {
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	var roads extra.AutoGenerated3
	if err = json.Unmarshal(body, &roads); err != nil {
		log.Fatal(err.Error())
	}

	estimatedTravelTime := roads.Routes[0].Summary.TravelTimeInSeconds
	estimatedTravelTimeMinutes := estimatedTravelTime / 60

	_, err = Client.Collection(Collection).Doc(id).Set(Ctx, map[string]interface{}{
		"estimatedTravelTime": estimatedTravelTimeMinutes,
	}, firestore.MergeAll)

}

func CallUrl(url string, content string) {

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(content)))
	if err != nil {
		fmt.Errorf("%v", "Error during request creation.")
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error in HTTP request: " + err.Error())
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Something is wrong with invocation response: " + err.Error())
	}
}

func Invoke(id string) {
	information, _ := Client.Collection(Collection).Doc(id).Get(Ctx)

	var message extra.Webhook
	if err := information.DataTo(&message); err != nil {
		log.Println(err.Error())
	}

	timeS, _ := time.Parse(time.RFC822, message.ArrivalTime)
	newTime := timeS.Add(time.Duration(-message.EstimatedTravelTime) * time.Minute)
	TimeUntilInvocation := time.Until(newTime).Seconds()
	fmt.Println(int(TimeUntilInvocation))

	fmt.Println(timeS)
	//timeTo := time.Duration(int(TimeUntilInvocation) * 1000000000)

	time.Sleep(5 * time.Second)
}
