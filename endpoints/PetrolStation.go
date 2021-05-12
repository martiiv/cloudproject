package endpoints

import (
	"cloudproject/structs"
	"cloudproject/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Function that will display all the electric-vehicle charging stations from a location, within 1km
func PetrolStation(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	address := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers

	latitude, longitude, err := utils.GetLocation(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitude + "&lon=" + longitude + "&radius=1000&categorySet=7311&key=" + utils.TomtomKey)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var petrol structs.Petrol
	if err = json.Unmarshal(body, &petrol); err != nil {
		utils.JsonUnmarshalErrorHandling(w, err)
		return
	}

	var total []structs.OutputPetrol
	for i := 0; i < len(petrol.Results); i++ {

		stationName := petrol.Results[i].Poi.Name
		var stationBrand string
		if len(petrol.Results[i].Poi.Brands) != 0 {
			stationBrand = petrol.Results[i].Poi.Brands[0].Name
		}
		address := petrol.Results[i].Address.FreeformAddress

		jsonStruct := structs.OutputPetrol{StationName: stationName, StationBrand: stationBrand, Address: address} //Creating a JSON object
		total = append(total, jsonStruct)                                                                          //Appending the json object to an array
	}

	output, err := json.Marshal(total) //Marshalling the array to JSON
	if err != nil {
		utils.JsonUnmarshalErrorHandling(w, err)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the chargers

}
