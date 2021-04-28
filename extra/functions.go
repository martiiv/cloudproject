package extra

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetLocation(address string) (float64, float64, error) {

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
