package database

import (
	"cloud.google.com/go/firestore"
	"cloudproject/structs"
	"cloudproject/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/**
 * Class database.go
 * Will contain all database related functionality
 * Contains the following functions:
 *									Init() 		For initializing the database connection
 *									Add() 		For adding an entry to the database
 *									Delete() 	For deleting an entry from the database
 *									Update() 	For updating an entry in the database
 * @author Martin Iversen
 * @date 01.05.2021
 * @version 0.1
 */

//Initializing DB
var Ctx context.Context
var Client *firestore.Client

var LocationCollection = "location"
var Collection = "message"

//const Collection = "RouteInformation" //Defining the name of the collection in FireStore

/*
 * Function for initializing the database, will be used when starting the app
 */ /*
func Init() error {
	// Firebase initialisation
	Ctx = context.Background()

	// Authenticate with key file from firebase
	opt := option.WithCredentialsFile("webhooks/cloudprojecttwo-firebase-adminsdk-uke12-fc63f46582.json")
	app, err := firebase.NewApp(Ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing DataBase: %v", err)
	}

	Client, err = app.Firestore(Ctx)
	if err != nil {
		return fmt.Errorf("error occurred initializing Client: %v", err)
	}

	go webhooks.InvokeAll()
	go webhooks.DeleteExpiredWebhooks()

	return nil
}*/

/*
 * Function for adding RouteInformation to the database
 * Returns the ID an object is given when the database creates
 */
func Add(webhook structs.Webhook) (string, error) {
	newEntry, _, err := Client.Collection(Collection).Add(Ctx, webhook) //Adds RouteInformation
	if err != nil {
		return "", errors.New("Error occurred when adding RouteInformation to database: " + err.Error())
	}
	return newEntry.ID, nil //Returns the id of the entry in the database collection
}

/*
 * Function for deleting a webhook from the database
 */
func Delete(id string) error {
	_, err := Client.Collection(Collection).Doc(id).Delete(Ctx) //Deletes from the database

	if err != nil {
		return errors.New("Error occurred when trying to delete entry. Entry ID: " + id)
	}
	return nil
}

/**
 * Function Get
 * Used for selecting a specific DB entry
 */
func Get(id string) (error, map[string]interface{}) {
	dbSnapShot, err := Client.Collection(Collection).Doc(id).Get(Ctx)
	if err != nil {
		return fmt.Errorf("Error occurred There is no document in the db with the id: %v!", id), nil
	}

	entry := dbSnapShot.Data()
	return nil, entry
}

/*
 * Function for getting all entries in a database
 * Source: https://stackoverflow.com/a/61429531
 * Returns an object containing document snapshots of the entries in the database
 */
func GetAll() ([]*firestore.DocumentSnapshot, error) {
	var docs []*firestore.DocumentSnapshot               //Defining object to be returned
	iter := Client.Collection(Collection).Documents(Ctx) //Gets all entries in the database
	for {                                                //Iterates through the database
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil //Returns a list of entries
}

/*
 * Function for updating information on an entry in the database
 */
func Update(id string, data interface{}) error {
	_, err := Client.Collection(Collection).Doc(id).Set(Ctx, data)
	if err != nil {
		return errors.New("Error while updating Route Information entry in the database: " + err.Error())
	}
	return nil
}

/**
Function to GeoCode the different locations the user inputs
*/
func GetLocation(address string) (string, string, error) {

	address = strings.Replace(address, " ", "+", -1) //Replaces the spaces in location with %20, that will please the url-condition

	response, err := http.Get("https://www.mapquestapi.com/geocoding/v1/address?key=" + utils.MapQuestKey + "&inFormat=kvp&outFormat=json&location=" + address)
	if response.StatusCode == http.StatusBadRequest {
		return "", "", errors.New("Syntax Error, Bad request\nPlease ensure you have entered an existing location")
	} else if response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusForbidden {
		return "", "", errors.New("Internal Error\nPlease try again later")
	} else if err != nil {
		return "", "", errors.New("Internal Error\n" + err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", errors.New("error, no content\n" + err.Error()) //Return an error
	}

	var location structs.GeoLocation
	if err = json.Unmarshal(body, &location); err != nil {
		return "", "", errors.New("internal error\n" + err.Error())
	}

	latitude := location.Results[0].Locations[0].LatLng.Lat
	longitude := location.Results[0].Locations[0].LatLng.Lng

	longitudeS := strconv.FormatFloat(longitude, 'f', 6, 64) //Formatting the coordinates to string
	latitudeS := strconv.FormatFloat(latitude, 'f', 6, 64)

	return latitudeS, longitudeS, nil //Returning the latitude and longitude to the location
}

func LocationPresent(address string) (string, string, error) {
	loc, err := Client.Collection(LocationCollection).Doc(address).Get(Ctx)
	if err != nil {
		err.Error()
	}

	var location structs.LocationLonLat
	if err := loc.DataTo(&location); err != nil {
		log.Println(err.Error())
	}

	if location.Latitude != -1 && location.Longitude != -1 {
		fmt.Println(location.Latitude, location.Longitude)
	}

	a, s, err := GetLocation(address)
	if err != nil {
		return "", "", err
	}

	addressUnescaped, err := url.QueryUnescape(address)

	newEntry, _ := Client.Collection(LocationCollection).Doc(addressUnescaped).Set(Ctx, map[string]interface{}{
		"Latitude":  a,
		"Longitude": s,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	fmt.Println(newEntry)

	fmt.Println(a, s)
	return a, s, err
}
