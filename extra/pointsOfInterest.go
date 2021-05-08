package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Function that will display all the points of interest from a location, within 1km
func PointOfInterest(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	address := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for points of interest
	poiPath := strings.Split(request.URL.Path, `/`)[3]

	latitude, longitude, err := GetLocation(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, "ERROR, The searched place does not exist", http.StatusBadRequest)
		return
	}

	response, err := http.Get("https://api.tomtom.com/search/2/poiSearch/" + poiPath + ".json?lat=" + latitude + "&lon=" + longitude + "&radius=1000&key=gcP26xVobGHjX2VVWGTskjelxX81WA1G")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "ERROR, The searched point of interest does not exist", http.StatusBadRequest)
		return
	}

	var poi pointsOfInterest

	if err = json.Unmarshal(body, &poi); err != nil {
		http.Error(w, "ERROR, failed to unmarshall", http.StatusInternalServerError)
		return
	}

	var total []outputPoi
	for i := 0; i < len(poi.Results); i++ {

		poiName := poi.Results[i].Poi.Name
		poiPhoneNumber := poi.Results[i].Poi.Phone
		poiAddress := poi.Results[i].Address.Freeformaddress

		jsonStruct := outputPoi{Name: poiName, PhoneNumber: poiPhoneNumber, Address: poiAddress} //Creating a JSON object
		total = append(total, jsonStruct)                                                        //Appending the json object to an array
	}

	output, err := json.Marshal(total) //Marshalling the array to JSON
	if err != nil {
		http.Error(w, "ERROR, failed to marshall", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the relevant info about poi

}
