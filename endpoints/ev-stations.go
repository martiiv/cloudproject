package endpoints

import (
	extra "cloudproject/extra"
	structs "cloudproject/extra"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//Function that will display all the electric-vehicle charging stations from a location, within 1km
func EVStations(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	address := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers

	latitude, longitude, err := extra.GetLocation(address) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longitudeS := strconv.FormatFloat(longitude, 'f', 6, 64) //Formatting the coordinates to string
	latitudeS := strconv.FormatFloat(latitude, 'f', 6, 64)

	response, err := http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitudeS + "&lon=" + longitudeS + "&radius=1000&categorySet=7309&key=gOorFpmbH5GPKh6uGqcfJN76oKFKfswA")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var charge structs.AutoGenerated
	if err = json.Unmarshal(body, &charge); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var total []extra.Output
	for i := 0; i < len(charge.Results); i++ {
		addresse := charge.Results[i].Address.FreeformAddress
		chargeName := charge.Results[i].Poi.Name
		phone := charge.Results[i].Poi.Phone

		jsonStruct := extra.Output{Charger: chargeName, Address: addresse, Phone: phone} //Creating a JSON object
		total = append(total, jsonStruct)                                                //Appending the json object to an array
	}

	output, err := json.Marshal(total) //Marshalling the array to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the chargers

}
