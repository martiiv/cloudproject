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

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {
	//database.Init()

	// Creates instance of firebase
	database.Ctx = context.Background()
	sa := option.WithCredentialsFile("webhooks/trafficmessage.json")
	app, err := firebase.NewApp(database.Ctx, nil, sa)
	if err != nil {
		_ = fmt.Errorf("error initializing app: %v", err)
	}

	database.Client, err = app.Firestore(database.Ctx)
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

/*func main(){

	// Creates instance of firebase
	database.Ctx = context.Background()
	sa := option.WithCredentialsFile("webhooks/trafficmessage.json")
	app, err := firebase.NewApp(database.Ctx, nil, sa)
	if err != nil {
		_ = fmt.Errorf("error initializing app: %v", err)
	}

	database.Client, err = app.Firestore(database.Ctx)
	if err != nil {
		log.Fatalln(err)
	}

	database.Get("NKl0kZ4pxbaDVTN2G6gx")
}*/

func handlers() {
	http.HandleFunc("/weather/", endpoints.CurrentWeather)
	http.HandleFunc("/poi/", endpoints.PointOfInterest)
	http.HandleFunc("/diag", endpoints.Diag)
	http.HandleFunc("/charge/", endpoints.EVStations)
	http.HandleFunc("/petrol/", endpoints.PetrolStation)
	http.HandleFunc("/messages/", endpoints.Messages)
	http.HandleFunc("/route/", endpoints.Route)
	http.HandleFunc("/hook/", webhooks.CreateWebhook)

	log.Println(http.ListenAndServe(getPort(), nil))
}
