package endpoints

import (
	extra "cloudproject/extra"
	structs "cloudproject/extra"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func EVStations(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	address := strings.Split(request.URL.Path, `/`)[2]
	fmt.Println(address)

	latitude, longitude, err := extra.GetLocation(address)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	longitudeS := strconv.FormatFloat(longitude, 'f', 5, 64)
	latitudeS := strconv.FormatFloat(latitude, 'f', 5, 64)

	response, err := http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitudeS + "&lon=" + longitudeS + "&radius=1000&categorySet=7309&key=gOorFpmbH5GPKh6uGqcfJN76oKFKfswA")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var charge structs.AutoGenerated
	if err = json.Unmarshal(body, &charge); err != nil {
		log.Println(err.Error())
	}

	addresse := charge.Results[0].Address.FreeformAddress
	chargeName := charge.Results[0].Poi.Name
	phone := charge.Results[0].Poi.Phone
	availability := charge.Results[0].DataSources.ChargingAvailability.ID

	charger := "Charger"
	addressJson := "Address"
	PhoneJson := "Phone"
	json := make(map[string]string)

	json[charger] = chargeName
	json[addressJson] = addresse
	json[PhoneJson] = phone

	fmt.Println(addresse)

	fmt.Fprintf(w, `{
   	"Charger": "%v",
   "Address": "%v",
	"Phone": "%v", 
	"AvaliabilityID: "%v"
}`,
		chargeName, addresse, phone, availability)

}
