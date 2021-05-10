package endpoints

import (
	extra "cloudproject/extra"
	structs "cloudproject/extra"
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
	address := strings.Split(request.URL.Path, `/`)[2]                      //Getting the address/name of the place we want to look for chargers
	latitude, longitude, err := extra.GetLocation(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(request.URL.RawQuery) != 0 {
		fmt.Println(strings.Split(request.URL.RawQuery, `=`)[1])

	}

	response, err := http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitude + "&lon=" + longitude + "&radius=1000&categorySet=7309&key=" + extra.TomtomKey)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var charge structs.Charger
	if err = json.Unmarshal(body, &charge); err != nil {
		extra.JsonUnmarshalErrorHandling(w, err)
		return
	}

	var total []extra.OutputCharge
	for i := 0; i < len(charge.Results); i++ {
		addresse := charge.Results[i].Address.FreeformAddress
		chargeName := charge.Results[i].Poi.Name
		phone := charge.Results[i].Poi.Phone
		var connector string
		var power float64
		if len(charge.Results[i].ChargingPark.Connectors) != 0 {
			connector = charge.Results[i].ChargingPark.Connectors[0].ConnectorType
			power = charge.Results[i].ChargingPark.Connectors[0].RatedPowerKW
		}

		jsonStruct := extra.OutputCharge{Charger: chargeName, Address: addresse, Phone: phone, Connectors: connector, PowerKW: power} //Creating a JSON object
		total = append(total, jsonStruct)                                                                                             //Appending the json object to an array
	}

	output, err := json.Marshal(total) //Marshalling the array to JSON
	if err != nil {
		extra.JsonUnmarshalErrorHandling(w, err)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the chargers

}
