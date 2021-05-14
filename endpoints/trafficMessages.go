package endpoints

import (
	"cloudproject/database"
	"cloudproject/structs"
	"cloudproject/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Messages(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(strings.Split(request.URL.Path, `/`)) != 4 {
		http.Error(w, "error Bad Request\nExpected input: /messages/{startLocation}/{endDestination}", http.StatusBadRequest)
		return
	}
	StartAddress := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers
	EndAddress := strings.Split(request.URL.Path, `/`)[3]   //Getting the address/name of the place we want to look for chargers

	bodyBox, err := getBBox(StartAddress, EndAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var bbox structs.AutoGenerated
	if err = json.Unmarshal(bodyBox, &bbox); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var box string
	for j := 0; j < len(bbox.Bbox); j++ {
		coordinate := strconv.FormatFloat(bbox.Bbox[j], 'f', 6, 64)
		box += coordinate + ","
	}
	box = strings.TrimRight(box, ",")

	response, err := http.Get("https://api.tomtom.com/traffic/services/5/incidentDetails?bbox=" + url.QueryEscape(box) +
		"&fields=%7Bincidents%7Btype%2Cgeometry%7Btype%2Ccoordinates%7D%2Cproperties%7Bid%2CiconCategory%2CmagnitudeOfDelay%2Cevents%7Bdescription%2Ccode%7D%2CstartTime%2Cend" +
		"Time%2Cfrom%2Cto%2Clength%2Cdelay%2CroadNumbers%2Caci%7BprobabilityOfOccurrence%2CnumberOfReports%2ClastReportTime%7D%7D%7D%7D&key=" + utils.TomtomKey)
	err = utils.TomTomErrorHandling(response.StatusCode)
	if err != nil {
		http.Error(w, err.Error(), response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var messages structs.Incidents
	if err = json.Unmarshal(body, &messages); err != nil {
		jsonError := utils.JsonUnmarshalErrorHandling(err)
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
		return
	}

	var all []structs.OutIncident
	time := time.Now().Add(-60 * time.Minute)

	for i := 0; i < len(messages.Incidents); i++ {
		if messages.Incidents[i].Properties.EndTime.Before(time) {
			startTime := messages.Incidents[i].Properties.StartTime
			endTime := messages.Incidents[i].Properties.EndTime
			FromAddress := messages.Incidents[i].Properties.From
			toAddress := messages.Incidents[i].Properties.To
			Event := messages.Incidents[i].Properties.Events[0].Description

			if strings.ContainsAny(FromAddress, "Ã¸") {
				FromAddress = strings.ReplaceAll(FromAddress, "Ã¸", "ø")
			} else if strings.ContainsAny(toAddress, "Ã¸") {
				toAddress = strings.ReplaceAll(toAddress, "Ã¸", "ø")
			}
			incidents := structs.OutIncident{From: FromAddress, To: toAddress, Start: startTime, End: endTime, Event: Event}
			all = append(all, incidents)
		}
	}

	output, err := json.Marshal(all) //Marshalling the array to JSON
	if err != nil {
		jsonError := utils.JsonUnmarshalErrorHandling(err)
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the chargers

}

func getBBox(StartAddress string, endAddress string) ([]byte, error) {
	//Todo return the statuscode
	startLat, startLong, err := database.LocationPresent(url.QueryEscape(StartAddress))
	if err != nil {
		return nil, err
	}

	EndLat, endLong, err := database.LocationPresent(url.QueryEscape(endAddress))
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.openrouteservice.org/v2/directions/driving-car?api_key=" + utils.OpenRouteServiceKey + "&start=" + startLong + "," + startLat + "&end=" + endLong + "," + EndLat)
	if err != nil {
		return nil, utils.OpenRouteError(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
