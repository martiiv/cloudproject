package endpoints

import (
	"cloudproject/extra"
	"fmt"
	"net/http"
	"strings"
)

func Messages(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	StartAddress := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers
	EndAddress := strings.Split(request.URL.Path, `/`)[3]   //Getting the address/name of the place we want to look for chargers

	minLatitude, minLongitude, err := extra.GetLocation(StartAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	maxLatitude, maxLongitude, err := extra.GetLocation(EndAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	latitude := maxLatitude

	bbox := minLongitude + "," + minLatitude + "," + maxLongitude + "," + latitude

	fmt.Println(bbox)

	/*response, err := http.Get("https://api.tomtom.com/traffic/services/5/incidentDetails?bbox=10.696607%2C60.792033%2C10.466012%2C61.113795" +
	"&fields=%7Bincidents%7Btype%2Cgeometry%7Btype%2Ccoordinates%7D%2Cproperties%7Bid%2CiconCategory%2CmagnitudeOfDelay%2Cevents%7Bdescription%2Ccode%7D%2CstartTime%2Cend" +
	"Time%2Cfrom%2Cto%2Clength%2Cdelay%2CroadNumbers%2Caci%7BprobabilityOfOccurrence%2CnumberOfReports%2ClastReportTime%7D%7D%7D%7D&key=gOorFpmbH5GPKh6uGqcfJN76oKFKfswA")

	*/

}
