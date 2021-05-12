package webhooks

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"cloudproject/endpoints"
	"cloudproject/extra"
	"cloudproject/structs"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CalculateDeparture(id string) {

	information, _ := Client.Collection(Collection).Doc(id).Get(Ctx)

	var message structs.Webhook
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

	var roads structs.AutoGenerated3
	if err = json.Unmarshal(body, &roads); err != nil {
		log.Fatal(err.Error())
	}

	estimatedTravelTime := roads.Routes[0].Summary.TravelTimeInSeconds
	estimatedTravelTimeMinutes := (estimatedTravelTime + endpoints.GetMessageWeight(message.Weather)) / 60

	_, err = Client.Collection(Collection).Doc(id).Set(Ctx, map[string]interface{}{
		"estimatedTravelTime": estimatedTravelTimeMinutes,
	}, firestore.MergeAll)

}

func CallUrl(url string, content string) {

	req, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(content)))
	if err != nil {
		fmt.Errorf("%v", "Error during request creation.")
		return
	}
	if req.StatusCode != http.StatusOK {
		fmt.Println("Fuck you")
	}

}

func Invoke(id string) {
	information, _ := Client.Collection(Collection).Doc(id).Get(Ctx)

	var message structs.Webhook
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

func SendNotification(notificationId string) {

	doc, err := Client.Collection(Collection).Doc(notificationId).Get(Ctx) // Loop through all entries in collection "messages"
	if err != nil {
		_ = errors.New("The notification ID is not in our system")
		return
	}
	var firebase structs.Webhook
	var message string
	var url string
	var TimeUntilInvocation float64

	if err := doc.DataTo(&firebase); err != nil {
		return
	}
	url = firebase.Url
	message = firebase.Weather

	timeS, _ := time.Parse(time.RFC822, firebase.ArrivalTime)
	newTime := timeS.Add(time.Duration(-firebase.EstimatedTravelTime) * time.Minute)
	TimeUntilInvocation = time.Until(newTime).Minutes()
	if TimeUntilInvocation < 0 {
		return
	}
	fmt.Println(TimeUntilInvocation)

	jsonMessage := structs.NotificationResponse{
		Text: message,
	}

	jsonStart := `{"text": "`
	jsonMiddle := jsonMessage.Text
	jsonEnd := `"}`
	jsonData := []byte(jsonStart + jsonMiddle + jsonEnd)

	time.Sleep(time.Duration(TimeUntilInvocation) * time.Minute)

	//Todo Check if the firebase is deleted before invocation
	var maps map[string]interface{}
	err, maps = Get(notificationId)
	if err != nil {
		return
	}

	fmt.Println(maps)
	go CallUrl(url, string(jsonData))

}

func InvokeAll() {
	webhook, err := GetAll()
	if err != nil {
		log.Fatalf(err.Error())
	}
	for i := 0; i < len(webhook); i++ {
		go SendNotification(webhook[i].Ref.ID)
	}
}
