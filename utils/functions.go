package utils

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// JsonUnmarshalErrorHandling Universal json unmarshalling error handler
func JsonUnmarshalErrorHandling(w http.ResponseWriter, err error) {
	errorString := "Unable to continue your request\n" +
		"Internal error: " + err.Error()
	http.Error(w, errorString, http.StatusInternalServerError)
}

// TomTomErrorHandling Universal TomTom error handler
func TomTomErrorHandling(status int) error {
	if status == http.StatusBadRequest {
		return errors.New("Error, Bad Request: " + strconv.Itoa(status) + "\nNo valid location requested.")
	} else if status == http.StatusForbidden {
		return errors.New("Error, Status Forbidden: " + strconv.Itoa(status) + "\nThe service is no longer provided.")
	} else if status == http.StatusInternalServerError {
		return errors.New("Error, Internal Server Error: " + strconv.Itoa(status) + "\nThe service is for the moment down. Please try again later.")
	} else if status == http.StatusOK {
		return nil
	}
	return errors.New("Error: " + strconv.Itoa(status) + "\n An unexpected error has occurred")
}

// OpenRouteError Universal OpenRoute error handler
func OpenRouteError(w http.ResponseWriter, status int) {
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
