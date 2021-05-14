package main

import (
	"cloudproject/database"
	"cloudproject/endpoints"
	"cloudproject/webhooks"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"time"
)

//getPort sets the port to 8080
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

//main Function to start application, initializes database and webhooks
func main() {
	// Creates instance of firebase
	database.Ctx = context.Background()
	sa := option.WithCredentialsFile("webhooks/trafficmessage.json") //Initializes database
	app, err := firebase.NewApp(database.Ctx, nil, sa)
	if err != nil {
		log.Println("error occured when initializing database" + err.Error())
		_ = fmt.Errorf("error initializing app: %v", err)
	}

	database.Client, err = app.Firestore(database.Ctx) //Connects to the database
	if err != nil {
		log.Fatalln(err)
	}

	// Starts uptime of program
	endpoints.Uptime = time.Now()
	//Webhook handling
	go webhooks.InvokeAll()
	go webhooks.DeleteExpiredWebhooks()

	log.Println("Listening on port: " + getPort())
	handlers()

	defer database.Client.Close()
}

//handlers Function for redirecting endpoints
func handlers() {
	http.HandleFunc("/rtc/v1/weather/", endpoints.CurrentWeather)
	http.HandleFunc("/rtc/v1/poi/", endpoints.PointOfInterest)
	http.HandleFunc("/rtc/v1/diag", endpoints.Diag)
	http.HandleFunc("/rtc/v1/charge/", endpoints.EVStations)
	http.HandleFunc("/rtc/v1/petrol/", endpoints.PetrolStation)
	http.HandleFunc("/rtc/v1/messages/", endpoints.Messages)
	http.HandleFunc("/rtc/v1/route/", endpoints.Route)
	http.HandleFunc("/rtc/v1/hook/", webhooks.AddWebhook)

	log.Println(http.ListenAndServe(getPort(), nil))
}
