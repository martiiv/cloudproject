package endpoints

import (
	"cloudproject/database"
	structs2 "cloudproject/structs"
	"cloudproject/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Function that will display all the electric-vehicle charging stations from a location, within 1km
func EVStations(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	address := strings.Split(request.URL.Path, `/`)[2]                             //Getting the address/name of the place we want to look for chargers
	latitude, longitude, err := database.LocationPresent(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := utils.GetOptionalFilter(request.URL)

	var response *http.Response

	if len(filter) != 0 {
		if len(filter["charge"]) == 0 && len(filter["radius"]) == 0 && len(filter["power"]) == 0 {
			http.Error(w, "error, Bad Request\nNone of the filters is accepted\nAccepted filters: radius, charge, power", http.StatusBadRequest)
			return
		}
		connector := ""
		power := ""
		radius := "&radius=1000"

		if len(filter["charge"]) != 0 {
			connector = "&connectorSet=" + filter["charge"]
		}
		if len(filter["radius"]) != 0 {
			radius = "&radius=" + filter["radius"]
		}
		if len(filter["power"]) != 0 {
			power = "&minPowerKW=" + filter["power"]
		}
		response, err = http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitude + "&lon=" + longitude + radius + "&connectorSet=" + connector + power + "&categorySet=7309&key=" + extra.TomtomKey)
	} else if len(filter) == 0 {
		response, err = http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitude + "&lon=" + longitude + "&radius=1000&categorySet=7309&key=" + extra.TomtomKey)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var charge structs2.Charger
	if err = json.Unmarshal(body, &charge); err != nil {
		extra.JsonUnmarshalErrorHandling(w, err)
		return
	}

	var total []structs2.OutputCharge
	for i := 0; i < len(charge.Results); i++ {
		addresse := charge.Results[i].Address.FreeformAddress
		chargeName := charge.Results[i].Poi.Name
		phone := charge.Results[i].Poi.Phone
		var connector string
		var power float64

		var connectorStruct []structs2.Connectors

		if len(charge.Results[i].ChargingPark.Connectors) != 0 {

			for j := 0; j < len(charge.Results[i].ChargingPark.Connectors); j++ {
				connector = charge.Results[i].ChargingPark.Connectors[j].ConnectorType
				power = charge.Results[i].ChargingPark.Connectors[j].RatedPowerKW
				connectors := structs2.Connectors{ConnectorType: connector, RatedPowerKW: power}
				connectorStruct = append(connectorStruct, connectors)
			}
		}

		jsonStruct := structs2.OutputCharge{Charger: chargeName, Address: addresse, Phone: phone, Connectors: connectorStruct} //Creating a JSON object
		total = append(total, jsonStruct)                                                                                      //Appending the json object to an array
	}

	output, err := json.Marshal(total) //Marshalling the array to JSON
	if err != nil {
		extra.JsonUnmarshalErrorHandling(w, err)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the chargers

}
