package endpoints

import (
	"cloudproject/extra"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var maneuvers = map[string]string{
	"ARRIVE":               "You have arrived.",
	"ARRIVE_LEFT":          "You have arrived. Your destination is on the left.",
	"ARRIVE_RIGHT":         "You have arrived. Your destination is on the right.",
	"DEPART":               "Leave.",
	"STRAIGHT":             "Keep straight on.",
	"KEEP_RIGHT":           "Keep right.",
	"BEAR_RIGHT":           "Bear right.",
	"TURN_RIGHT":           "Turn right.",
	"SHARP_RIGHT":          "Turn sharp right.",
	"KEEP_LEFT":            "Keep left.",
	"BEAR_LEFT":            "Bear left.",
	"TURN_LEFT":            "Turn left.",
	"SHARP_LEFT":           "Turn sharp left.",
	"MAKE_UTURN":           "Make a U-turn.",
	"ENTER_MOTORWAY":       "Take the motorway.",
	"ENTER_FREEWAY":        "Take the freeway.",
	"ENTER_HIGHWAY":        "Take the highway.",
	"TAKE_EXIT":            "Take the exit.",
	"MOTORWAY_EXIT_LEFT":   "Take the left exit.",
	"MOTORWAY_EXIT_RIGHT":  "Take the right exit.",
	"TAKE_FERRY":           "Take the ferry.",
	"ROUNDABOUT_CROSS":     "Cross the roundabout.",
	"ROUNDABOUT_RIGHT":     "At the roundabout take the exit on the right.",
	"ROUNDABOUT_LEFT":      "At the roundabout take the exit on the left.",
	"ROUNDABOUT_BACK":      "Go around the roundabout.",
	"TRY_MAKE_UTURN":       "Try to make a U-turn.",
	"FOLLOW":               "Follow.",
	"SWITCH_PARALLEL_ROAD": "Switch to the parallel road.",
	"SWITCH_MAIN_ROAD":     "Switch to the main road.",
	"ENTRANCE_RAMP":        "Take the ramp.",
	"WAYPOINT_LEFT":        "You have reached the waypoint. It is on the left.",
	"WAYPOINT_RIGHT":       "You have reached the waypoint. It is on the right.",
	"WAYPOINT_REACHED":     "You have reached the waypoint."}

func Route(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	StartAddress := strings.Split(request.URL.Path, `/`)[2] //Getting the address/name of the place we want to look for chargers
	EndAddress := strings.Split(request.URL.Path, `/`)[3]   //Getting the address/name of the place we want to look for chargers

	startLat, startLong, err := extra.GetLocation(StartAddress)
	if err != nil {
	}

	EndLat, endLong, err := extra.GetLocation(EndAddress)
	if err != nil {
	}

	coordinates := startLat + "%2C" + startLong + "%3A" + EndLat + "%2C" + endLong

	fmt.Println(coordinates)

	resp, err := http.Get("https://api.tomtom.com/routing/1/calculateRoute/" + coordinates + "/json?instructionsType=coded&traffic=false&avoid=unpavedRoads&travelMode=car&key=" + extra.TomtomKey)
	if err != nil {
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	var roads extra.
	if err = json.Unmarshal(body, &roads); err != nil {
		log.Fatal(err.Error())
	}

	var maneuver string
	var junctionType string
	var RoadNumber string
	var Street string

	var total []extra.Route

	drivingDuration := roads.Routes[0].Summary.TravelTimeInSeconds

	drivingDurationMinutes := drivingDuration / 69
	fmt.Println(drivingDurationMinutes)

	drivingLength := roads.Routes[0].Summary.LengthInMeters / 1000
	estimatedTime := roads.Routes[0].Summary.ArrivalTime
	estimatedTimeString := estimatedTime.Format("2006-01-02 15:04:05")

	for i := 0; i < len(roads.Routes[0].Guidance.Instructions); i++ {
		maneuver = roads.Routes[0].Guidance.Instructions[i].Maneuver
		maneuver = maneuvers[maneuver]
		junctionType = roads.Routes[0].Guidance.Instructions[i].JunctionType
		if roads.Routes[0].Guidance.Instructions[i].RoadNumbers != nil {
			RoadNumber = roads.Routes[0].Guidance.Instructions[i].RoadNumbers[0]
		}

		Street = roads.Routes[0].Guidance.Instructions[i].Street

		route := extra.Route{Street: Street, RoadNumber: RoadNumber, Maneuver: maneuver, JunctionType: junctionType}
		total = append(total, route)
	}

	information := extra.RoadInformation{EstimatedArrival: estimatedTimeString, LengthKM: drivingLength, Route: total}

	output, err := json.Marshal(information) //Marshalling the array to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(output)) //Outputs the chargers

}
