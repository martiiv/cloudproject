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

	poi123 := strings.Split(request.URL.Path, `/`)[3]

	latitude, longitude, err := GetLocation(url.QueryEscape(address)) //Receives the latitude and longitude of the place passed in the url
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := http.Get("https://api.tomtom.com/search/2/poiSearch/" + poi123 + ".json?lat=" + latitude + "&lon=" + longitude + "&radius=1000&key=gcP26xVobGHjX2VVWGTskjelxX81WA1G")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var poi pointsOfInterest

	if err = json.Unmarshal(body, &poi); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Vil noe bli printet ut
	fmt.Fprintf(w, "PointOfInterest: %+v", poi)

}
