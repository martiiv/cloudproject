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
 * Will contain database related functionality
 * Contains the following functions:
 *							Delete() 			For deleting an entry from the database
 *							Get()				For retrieving an instance from the database
 * 							GetAll()			For retrieving all instances from the database
 *							Update() 			For updating an entry in the database
 * 							GetLocation()		For getting a location from the API to be put into the database
 * 							LocationPresent()	For checking for, and retrieving a location from the database
 */

// Ctx Initializing the context to be used with firebase
var Ctx context.Context

// Client Initializing the firebase client
var Client *firestore.Client

// LocationCollection Name of the collection containing locations in firebase
var LocationCollection = "location"

// Collection Name of the collection containing webhooks in firebase
var Collection = "message"

// Delete Function for deleting an instance from the database (for instance webhooks)
func Delete(id string) error {
	_, err := Client.Collection(Collection).Doc(id).Delete(Ctx) //Deletes from the database

	if err != nil {
		return errors.New("Error occurred when trying to delete entry. Entry ID: " + id)
	}
	return nil
}

// Get Used for retrieving a specific database entry and its data
func Get(id string) (map[string]interface{}, error) {
	dbSnapShot, err := Client.Collection(Collection).Doc(id).Get(Ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred: There is no document in the db with the id: %v", id)
	}

	entry := dbSnapShot.Data()
	return entry, nil
}

// GetAll Retrieves all entries in a database
// Source: https://stackoverflow.com/a/61429531
// 		- Decided to use this because it is a quick and easy way to retrieve all entries from the database
// Returns an object containing document snapshots of the entries in the database
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

// Update Updates information of an entry in the database
func Update(id string, data interface{}) error {
	_, err := Client.Collection(Collection).Doc(id).Set(Ctx, data)
	if err != nil {
		return errors.New("Error while updating information for entry: " + id + " in the database: " + err.Error())
	}
	return nil
}

// GetLocation Gets GeoCode from the API for the different locations the user inputs
func GetLocation(address string) (string, string, error) {

	address = strings.Replace(address, " ", "+", -1) //Replaces the spaces in location with +, which will please the url-condition

	// Asks the API for the location data
	response, err := http.Get("https://www.mapquestapi.com/geocoding/v1/address?key=" + utils.MapQuestKey + "&inFormat=kvp&outFormat=json&location=" + address)
	if response.StatusCode == http.StatusBadRequest {
		return "", "", errors.New("Syntax Error, Bad request, Status code: " + strconv.Itoa(response.StatusCode) + "\nPlease ensure you have entered an existing location")
	} else if response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusForbidden {
		return "", "", errors.New("Internal Error, Status code: " + strconv.Itoa(response.StatusCode) + "\nPlease try again later")
	} else if err != nil {
		return "", "", errors.New("Internal Error, Status code: " + strconv.Itoa(response.StatusCode) + "\n" + err.Error())
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

// LocationPresent Tries to get the location the user asks for from the database, if the location is not present in
// the database, ask GetLocation to retrieve the data from the API, and then stores it into the database
func LocationPresent(address string) (string, string, error) {
	// To remove broken syntax for some UTF8 characters
	addressUnescaped, errQuery := url.QueryUnescape(address)
	addressUnescaped = strings.ToLower(addressUnescaped)
	if errQuery != nil {
		log.Println(errQuery.Error())
		return "", "", errQuery
	}

	// Tries to retrieve the given document from the database
	loc, errRetrieve := Client.Collection(LocationCollection).Doc(addressUnescaped).Get(Ctx)
	if errRetrieve != nil {
		log.Println("Address: " + addressUnescaped + " is not present in the location database. It will be added.")
	}

	locLat, locLon, err := "", "", errors.New("")

	// If we were able to retrieve the location data from the database, validate the data and bind it to variables to be returned
	if errRetrieve == nil {
		var location structs.LocationLonLat
		if err = loc.DataTo(&location); err != nil {
			log.Println(err.Error())
		}
		locLat = location.Latitude
		locLon = location.Longitude
	} else { // Not able to retrieve the location data from the database
		// Call the API to retrieve location data
		locLat, locLon, err = GetLocation(address)
		if err != nil {
			return "", "", err
		}

		// Add the new location instance to the database to be easily access next time
		_, errSetLoc := Client.Collection(LocationCollection).Doc(addressUnescaped).Set(Ctx, map[string]interface{}{
			"Latitude":  locLat,
			"Longitude": locLon,
		})
		if errSetLoc != nil {
			err = errSetLoc
			log.Println(errSetLoc.Error())
			return "", "", errSetLoc
		} else {
			log.Println("Address " + addressUnescaped + " was successfully added to the database.")
		}
	}
	return locLat, locLon, err
}
