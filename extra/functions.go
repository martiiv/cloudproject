package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getLocation(address string) (float64, float64, error) {

	address = strings.Replace(address, " ", "%20", -1)

	response, err := http.Get("http://api.positionstack.com/v1/forward?access_key=3a2c0bbe3ee774328656aebd577398c3&query=" + address)
	if err != nil {
		return 0, 0, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, 0, err
	}

	var location geoLocation
	if err = json.Unmarshal(body, &location); err != nil {
		log.Println(err.Error())
	}

	latitude := location.Data[0].Latitude
	longitude := location.Data[0].Longitude

	return latitude, longitude, nil
}

func evStations(w http.ResponseWriter, request *http.Request) {

	address := strings.Split(request.URL.Path, `/`)[4]
	radius := strings.Split(request.URL.Path, `/`)[5]

	latitude, longitude, err := getLocation(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	longitudeS := strconv.FormatFloat(longitude, 'f', 5, 64)
	latitudeS := strconv.FormatFloat(latitude, 'f', 5, 64)

	response, err := http.Get("https://api.tomtom.com/search/2/nearbySearch/.json?lat=" + latitudeS + "&lon=" + longitudeS + "&radius=" + radius + "&categorySet=7309&fuelSet=Petrol&key=3a2c0bbe3ee774328656aebd577398c3")

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var charge EVCharge
	if err = json.Unmarshal(body, &charge); err != nil {
		log.Println(err.Error())
	}

	fmt.Println(charge.Results[0].Address.StreetName)

}
