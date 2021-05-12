package utils

import (
	"cloudproject/structs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/**
Function to GeoCode the different locations the user inputs
*/
func GetLocation(address string) (string, string, error) {

	address = strings.Replace(address, " ", "+", -1) //Replaces the spaces in location with %20, that will please the url-condition

	response, err := http.Get("https://www.mapquestapi.com/geocoding/v1/address?key=" + MapQuestKey + "&inFormat=kvp&outFormat=json&location=" + address)
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

func JsonUnmarshalErrorHandling(w http.ResponseWriter, err error) {
	errorString := "Unable to continue your request\n" +
		"Internal error " + err.Error()
	http.Error(w, errorString, http.StatusInternalServerError)
}

func TomTomErrorHandling(w http.ResponseWriter, status int) error {
	if status == http.StatusBadRequest {
		return errors.New("error, Bad Request\nNo valid location requested.")
	} else if status == http.StatusForbidden {
		return errors.New("error\nThe service is no longer provided.")
	} else if status == http.StatusInternalServerError {
		return errors.New("error Internal Server Error\nThe service is for the moment down. Please try again later")

	} else if status == http.StatusOK {
		return nil

	}
	return errors.New("error\n An unexpected error has occurred")
}

func openRouteError(w http.ResponseWriter, status int) {
	//Todo make this error function
	//https://openrouteservice.org/dev/#/api-docs/v2/directions/{profile}/get
}

func GetOptionalFilter(url *url.URL) map[string]string {
	var optionals = map[string]string{}
	optional := strings.Split(url.RawQuery, "?")
	if len(optional) != 0 && optional[0] != "" {

		for i := 0; i <= len(optional)-1; i++ {
			nameOfFilter := strings.Split(optional[i], "=")
			valueName := nameOfFilter[1]
			mapName := nameOfFilter[0]

			optionals[mapName] = valueName

		}
		return optionals
	}
	return nil
}
