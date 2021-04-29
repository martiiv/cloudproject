package extra

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetLocation(address string) (string, string, error) {

	address = strings.Replace(address, " ", "+", -1) //Replaces the spaces in location with %20, that will please the url-condition

	response, err := http.Get("https://www.mapquestapi.com/geocoding/v1/address?key=UvCctIMBPNYcpfiAkTCkVjakeCjoPpPR&inFormat=kvp&outFormat=json&location=" + address)
	if err != nil {
		return "", "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	var location geoLocation
	if err = json.Unmarshal(body, &location); err != nil {
		return "", "", err
	}

	latitude := location.Results[0].Locations[0].LatLng.Lat
	longitude := location.Results[0].Locations[0].LatLng.Lng

	longitudeS := strconv.FormatFloat(longitude, 'f', 6, 64) //Formatting the coordinates to string
	latitudeS := strconv.FormatFloat(latitude, 'f', 6, 64)

	return latitudeS, longitudeS, nil //Returning the latitude and longitude to the location
}
